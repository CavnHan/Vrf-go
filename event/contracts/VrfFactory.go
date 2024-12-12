package  contracts

import (
	"context"
	"fmt"
	"math/big"
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

type VrfFactory struct {
	VrfFactoryAbi    *abi.ABI
	VrfFactoryFilter *bindings.VRFFactoryFilterer
}

func NewVrfFactory() (*VrfFactory, error) {
	VrfFactoryAbi, err := bindings.VRFFactoryMetaData.GetAbi()
	if err != nil {
		log.Error("get  vrf factory abi fail", "err", err)
		return nil, err
	}
	VrfFactoryFilterer, err := bindings.NewVRFFactoryFilterer(common.Address{}, nil)
	if err != nil {
		log.Error("new  vrf factory filter fail", "err", err)
		return nil, err
	}
	return &VrfFactory{
		VrfFactoryAbi:    VrfFactoryAbi,
		VrfFactoryFilter: VrfFactoryFilterer,
	}, nil
}

func (vff *VrfFactory) ProcessVrfFactoryEvent(db *database.DB, VrfFactoryAddres string, startHeight, endHeight *big.Int) error {
	contactFilter := event.ContractEvent{ContractAddress: common.HexToAddress(VrfFactoryAddres)}
	contractEventList, err := db.ContractEvent.ContractEventsWithFilter(contactFilter, startHeight, endHeight)
	if err != nil {
		log.Error("query contacts event fail", "err", err)
		return err
	}
	var proxyCreatedList []worker.PoxyCreated
	for _, contractEvent := range contractEventList {
		if contractEvent.EventSignature.String() == vff.VrfFactoryAbi.Events["ProxyCreated"].ID.String() {
			proxyCreated, err := vff.VrfFactoryFilter.ParseProxyCreated(*contractEvent.RLPLog)
			if err != nil {
				log.Error("proxy created fail", "err", err)
				return err
			}
			log.Info("proxy created event", "MintProxyAddress", proxyCreated.ProxyAddress)
			pc := worker.PoxyCreated{
				GUID:         uuid.New(),
				ProxyAddress: proxyCreated.ProxyAddress,
				Timestamp:    uint64(time.Now().Unix()),
			}
			proxyCreatedList = append(proxyCreatedList, pc)
		}
	}

	retryStrategy := &retry.ExponentialStrategy{Min: 1000, Max: 20_000, MaxJitter: 250}
	if _, err := retry.Do[interface{}](context.Background(), 10, retryStrategy, func() (interface{}, error) {
		if err := db.Transaction(func(tx *database.DB) error {
			if err := tx.PoxyCreated.StorePoxyCreated(proxyCreatedList); err != nil {
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
