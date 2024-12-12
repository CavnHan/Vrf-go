package contracts

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"

	"github.com/CavnHan/Vrf-go/bindings"
	"github.com/CavnHan/Vrf-go/database"
	"github.com/CavnHan/Vrf-go/database/event"
	"github.com/CavnHan/Vrf-go/database/worker"
	"github.com/CavnHan/Vrf-go/synchronizer/retry"
)

type Vrf struct {
	VrfAbi    *abi.ABI
	VrfFilter *bindings.VRFFilterer
}

func NewVrf() (*Vrf, error) {
	//get abi from bindings
	VrfAbi, err := bindings.VRFMetaData.GetAbi()
	if err != nil {
		log.Error("Failed to get VRF abi")
		return nil, err
	}

	VrfFilter, err := bindings.NewVRFFilterer(common.Address{}, nil)
	if err != nil {
		log.Error("Failed to create VRF filter")
		return nil, err
	}
	return &Vrf{
		VrfAbi:    VrfAbi,
		VrfFilter: VrfFilter,
	}, nil

}

func (vf *Vrf) ProcessVrfEvent(db *database.DB, VrfAddress string, startHeight, endHeight *big.Int) error {
	contractFilter := event.ContractEvent{ContractAddress: common.HexToAddress(VrfAddress)}
	contractEventList,err := db.ContractEvent.ContractEventsWithFilter(contractFilter,startHeight,endHeight)
	if err != nil {
		log.Error("Failed to get contract event list")
		return err
	}
	var RequestSendList []worker.RequestSend
	var FillRandomWordList []worker.FillRandomWords
	for _,contractEvent := range contractEventList {
		//check if the event is RequestSent
		if contractEvent.EventSignature.String() == vf.VrfAbi.Events["RequestSent"].ID.String(){
			//parse the event
			requestSendEvent,err := vf.VrfFilter.ParseRequestSend(*contractEvent.RLPLog)
			if err != nil {
				log.Error("Failed to parse RequestSend event")
				return err
			}
			log.Info("Request sent event", "RequestId", requestSendEvent.RequestId, "NumWords", requestSendEvent.Numwords, "Current", requestSendEvent.Current)
			rs := worker.RequestSend{
				GUID: uuid.New(),
				RequestId: requestSendEvent.RequestId,
				VrfAddress: requestSendEvent.Current,
				NumWords: requestSendEvent.Numwords,
				Status: 0,
				Timestamp: uint64(time.Now().Unix()),
			}
			RequestSendList = append(RequestSendList,rs)
		}
		if contractEvent.EventSignature.String() == vf.VrfAbi.Events["FillRandomWords"].ID.String() {
			fillRandomWords, err := vf.VrfFilter.ParseFillRandomWords(*contractEvent.RLPLog)
			if err != nil {
				log.Error("parse fill random fail", "err", err)
				return err
			}
			log.Info("Fill random words event", "RequestId", fillRandomWords.RequestId, "RandomWords", fillRandomWords.RandomWords)
			var sb strings.Builder
			//遍历randomwords并进行拼接
			for _, rword := range fillRandomWords.RandomWords {
				sb.WriteString(rword.String())
				sb.WriteString(",")
			}
			randomWords := sb.String()
			if len(randomWords) > 0 {
				randomWords = randomWords[:len(randomWords)-1]
			}
			frw := worker.FillRandomWords{
				GUID:        uuid.New(),
				RequestId:   fillRandomWords.RequestId,
				RandomWords: randomWords,
				Timestamp:   uint64(time.Now().Unix()),
			}
			FillRandomWordList = append(FillRandomWordList, frw)
		}
		
	}
	retryStrategy := &retry.ExponentialStrategy{Min: 1000, Max: 20_000, MaxJitter: 250}
	if _, err := retry.Do[interface{}](context.Background(), 10, retryStrategy, func() (interface{}, error) {
		if err := db.Transaction(func(tx *database.DB) error {
			if err := tx.FillRandomWords.StoreFillRandomWords(FillRandomWordList); err != nil {
				return err
			}
			if err := tx.RequestSend.StoreRequestSend(RequestSendList); err != nil {
				return err
			}
			return nil
		}); err != nil {
			log.Info("unable to persist batch", err)
			return nil, fmt.Errorf("unable to persist batch: %w", err)
		}
		return nil, nil
	}); err != nil {
		return err
	}

	return nil
}
