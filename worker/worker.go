package worker

import (
	"context"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/log"

	"github.com/CavnHan/Vrf-go/common/tasks"
	"github.com/CavnHan/Vrf-go/database"
	"github.com/CavnHan/Vrf-go/driver"
)

type WorkerConfig struct {
	LoopInterval time.Duration
}

type worker struct {
	workerConfig  *WorkerConfig
	db            *database.DB
	deg           *driver.DriverEingine
	resourceCtx   context.Context
	resourceCacel context.CancelFunc
	tasks         tasks.Group
}

func NewWorker(db *database.DB, deg *driver.DriverEingine, workerConfig *WorkerConfig, shutdown context.CancelCauseFunc) (*worker, error) {
	resCtx, resCancel := context.WithCancel(context.Background())
	return &worker{
		db:            db,
		deg:           deg,
		resourceCtx:   resCtx,
		resourceCacel: resCancel,
		tasks: tasks.Group{
			HandleCrit: func(err error) {
				shutdown(fmt.Errorf("critical error in birdge processor : %w", err))
			},
		},
	}, nil
}

func (w *worker) Start() error {
	log.Info("starting worker processor...")
	tickerEventWorker := time.NewTicker(w.workerConfig.LoopInterval)
	w.tasks.Go(func() error {
		for range tickerEventWorker.C {
			log.Info("start handler random for vrf")
			err := w.ProcessCallerVrf()
			if err != nil {
				log.Error("process caller vrf fail", "err", err)
				return err
			}
		}
		return nil
	})
	return nil
}

func (w *worker) Close() error {
	w.resourceCacel()
	return w.tasks.Wait()
}

func (w *worker) ProcessCallerVrf() error {
	//获取requestsend合约事件
	requestUnSendList, err := w.db.RequestSend.QueryUnHandleRequestSendList()
	if err != nil {
		log.Error("query unhandle request send list fail", "err", err)
		return err
	}
	for _, requestSend := range requestUnSendList {
		//组装随机数列表交易发到vrf合约
		txRecipit, err := w.deg.FulfillRandomWords(requestSend.RequestId, nil)
		if err != nil {
			log.Error("Fulfill random words fail", "err", err)
			return err
		}
		//判断交易状态
		if txRecipit.Status == 1 {
			err := w.db.RequestSend.MarkRequestSendFinish(requestSend)
			if err != nil {
				log.Error("mark request send event list fail", "err", err)
				return err
			}
		}

	}
	return nil
}
