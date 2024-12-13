package contracts

import (
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

func (vff *VrfFactory) ProcessVrfFactoryEvent(db *database.DB, VrfFactoryAddres string, startHeight, endHeight *big.Int) ([]worker.PoxyCreated, error) {
	var proxyCreatedList []worker.PoxyCreated
	contactFilter := event.ContractEvent{ContractAddress: common.HexToAddress(VrfFactoryAddres)}
	contractEventList, err := db.ContractEvent.ContractEventsWithFilter(contactFilter, startHeight, endHeight)
	if err != nil {
		log.Error("query contacts event fail", "err", err)
		return proxyCreatedList, err
	}
	for _, contractEvent := range contractEventList {
		if contractEvent.EventSignature.String() == vff.VrfFactoryAbi.Events["ProxyCreated"].ID.String() {
			proxyCreated, err := vff.VrfFactoryFilter.ParseProxyCreated(*contractEvent.RLPLog)
			if err != nil {
				log.Error("proxy created fail", "err", err)
				return proxyCreatedList, err
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
	return proxyCreatedList, nil
}
