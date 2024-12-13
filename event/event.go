package event

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"math/big"
	"time"

	"github.com/google/uuid"

	"github.com/ethereum/go-ethereum/log"

	"github.com/CavnHan/Vrf-go/common/bigint"
	"github.com/CavnHan/Vrf-go/common/tasks"
	"github.com/CavnHan/Vrf-go/database"
	"github.com/CavnHan/Vrf-go/database/common"
	"github.com/CavnHan/Vrf-go/database/worker"
	"github.com/CavnHan/Vrf-go/event/contracts"
	"github.com/CavnHan/Vrf-go/synchronizer/retry"
)

var blockLimit = 10_000

type EventsHandlerConfig struct {
	VrfAddress        string
	VrfFactoryAddress string
	LoopInterval      time.Duration
	StartHeigth       *big.Int
	Epoch             uint64
}

type EventsHandler struct {
	Vrf                  *contracts.Vrf
	VrfFactory           *contracts.VrfFactory
	db                   *database.DB
	eventsHanddlerConfig *EventsHandlerConfig
	latestBlcokHeader    *common.BlockHeader
	resourceCtx          context.Context
	resourceCancel       context.CancelFunc
	tasks                tasks.Group
}

func NewEventsHandler(db *database.DB, eventsHandlerConfig *EventsHandlerConfig, shutdown context.CancelCauseFunc) (*EventsHandler, error) {
	dappLinkVrf, err := contracts.NewVrf()
	if err != nil {
		log.Error("new dapplink vrf fail", "err", err)
		return nil, err
	}
	dappLinkVrfFactory, err := contracts.NewVrfFactory()
	if err != nil {
		log.Error("new dapplink vrf factory fail", "err", err)
		return nil, err
	}

	ltBlockHeader, err := db.EventBlocks.LatestEventBlockHeader()
	if err != nil {
		log.Error("fetch latest block header fail", "err", err)
		return nil, err
	}

	resCtx, resCancel := context.WithCancel(context.Background())

	return &EventsHandler{
		Vrf:                  dappLinkVrf,
		VrfFactory:           dappLinkVrfFactory,
		db:                   db,
		eventsHanddlerConfig: eventsHandlerConfig,
		latestBlcokHeader:    ltBlockHeader,
		resourceCtx:          resCtx,
		resourceCancel:       resCancel,
		tasks: tasks.Group{HandleCrit: func(err error) {
			shutdown(fmt.Errorf("critical error in bridge processor: %w", err))
		}},
	}, nil
}

func (eh *EventsHandler) Start() error {
	log.Info("starting event processor...")
	tickerEventWorker := time.NewTicker(eh.eventsHanddlerConfig.LoopInterval)
	eh.tasks.Go(func() error {
		for range tickerEventWorker.C {
			log.Info("start parse event logs")
			err := eh.processEvent()
			if err != nil {
				log.Info("process event error", "err", err)
				return err
			}
		}
		return nil
	})
	return nil
}

func (eh *EventsHandler) Close() error {
	eh.resourceCancel()
	return eh.tasks.Wait()
}

func (eh *EventsHandler) processEvent() error {
	lastBlockNumber := eh.eventsHanddlerConfig.StartHeigth
	if eh.latestBlcokHeader != nil {
		lastBlockNumber = eh.latestBlcokHeader.Number
	}
	log.Info("process event latest block number", "lastBlockNumber", lastBlockNumber)
	latestHeaderScope := func(db *gorm.DB) *gorm.DB {
		newQuery := db.Session(&gorm.Session{NewDB: true})
		headers := newQuery.Model(common.BlockHeader{}).Where("number > ?", lastBlockNumber)
		return db.Where("number = (?)", newQuery.Table("(?) as block_numbers", headers.Order("number ASC").Limit(blockLimit)).Select("MAX(number)"))
	}

	if latestHeaderScope == nil {
		return nil
	}

	latestBlockHeader, err := eh.db.Blocks.BlockHeaderWithScope(latestHeaderScope)
	if err != nil {
		log.Error("get latest block header with scope fail", "err", err)
		return err
	} else if latestBlockHeader == nil {
		log.Debug("no new block for process event")
		return nil
	}

	//get eventblock
	fromHeight, toHeight := new(big.Int).Add(lastBlockNumber, bigint.One), latestBlockHeader.Number
	eventBLocks := make([]worker.EventBlocks, 0, toHeight.Uint64()-fromHeight.Uint64())
	for index := fromHeight.Uint64(); index < toHeight.Uint64(); index++ {
		blockHeader, err := eh.db.Blocks.BlockHeaderByNumber(big.NewInt(int64(index)))
		if err != nil {
			return err
		}
		evBlock := worker.EventBlocks{
			GUID:       uuid.New(),
			Hash:       blockHeader.Hash,
			ParentHash: blockHeader.ParentHash,
			Number:     blockHeader.Number,
			Timestamp:  blockHeader.Timestamp,
		}
		eventBLocks = append(eventBLocks, evBlock)
	}
	//get event with address filter
	requestSendList, fillRandomWordList, err := eh.Vrf.ProcessVrfEvent(eh.db, eh.eventsHanddlerConfig.VrfAddress, fromHeight, toHeight)
	if err != nil {
		log.Error("process dapplink vrf event fail", "err", err)
		return err
	}

	proxyCreatedList, err := eh.VrfFactory.ProcessVrfFactoryEvent(eh.db, eh.eventsHanddlerConfig.VrfAddress, fromHeight, toHeight)
	if err != nil {
		return err
	}


	retryStrategy := &retry.ExponentialStrategy{Min: 1000, Max: 20_000, MaxJitter: 250}
	if _, err := retry.Do[interface{}](eh.resourceCtx, 10, retryStrategy, func() (interface{}, error) {
		if err := eh.db.Transaction(func(tx *database.DB) error {
			if len(requestSendList) > 0 {
				err := eh.db.RequestSend.StoreRequestSend(requestSendList)
				if err != nil {
					log.Error("store request send fail", "err", err)
					return err
				}
			}

			if len(fillRandomWordList) > 0 {
				err := eh.db.FillRandomWords.StoreFillRandomWords(fillRandomWordList)
				if err != nil {
					log.Error("store fill random words fail", "err", err)
					return err
				}
			}

			if len(proxyCreatedList) > 0 {
				err := eh.db.PoxyCreated.StorePoxyCreated(proxyCreatedList)
				if err != nil {
					log.Error("store proxy created fail", "err", err)
					return err
				}
			}

			if len(eventBLocks) > 0 {
				err := eh.db.EventBlocks.StoreEventBlocks(eventBLocks)
				if err != nil {
					log.Error("store event blocks fail", "err", err)
					return err
				}
			}
			return nil
		}); err != nil {
			log.Debug("unable to persist batch", err)
			return nil, fmt.Errorf("unable to persist batch: %w", err)
		}
		return nil, nil
	}); err != nil {
		return err
	}

	return nil
}
