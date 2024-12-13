package driver

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"math/big"
	"strings"
	"sync"
	"time"

	"github.com/CavnHan/Vrf-go/bindings"
	"github.com/CavnHan/Vrf-go/txmgr"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
)

var (
	errMaxPriorityFeePerGasNotFound = errors.New(
		"Method eth_maxPriorityFeePerGas not found",
	)
	FallbackGasTipCap = big.NewInt(1500000000)
)

type DriverEingineConfig struct {
	ChainClient        *ethclient.Client
	ChainId            *big.Int
	VRFContractAddress common.Address
	CallerAddress      common.Address
	PrivateKey         *ecdsa.PrivateKey //calleraddress 和 privatekey是一对
	//nonce too low报错计数器
	SafeAbortNonceTooLowCount uint64
	NumConfirmations          uint64
}

type DriverEingine struct {
	Ctx context.Context
	Cfg *DriverEingineConfig
	//contract methods:callet transfer filter
	VRFContractMethod *bindings.VRF
	//crate a new contract instance ,include :address abi rpcclient
	RawVrfContract *bind.BoundContract
	VRfContractABI *abi.ABI
	TxMgr          txmgr.TxManager
	cancel         func()
	wg             sync.WaitGroup
}

func NewDriverEingine(ctx context.Context, cfg *DriverEingineConfig) (*DriverEingine, error) {
	_, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	VRFContract, err := bindings.NewVRF(cfg.VRFContractAddress, cfg.ChainClient)
	if err != nil {
		log.Error("NewVRFContract failed", "err", err)
		return nil, err
	}
	parsed, err := abi.JSON(strings.NewReader(bindings.VRFMetaData.ABI))
	if err != nil {
		log.Error("abi.JSON failed", "err", err)
		return nil, err
	}
	VRFContractAbi, err := bindings.VRFMetaData.GetAbi()
	if err != nil {
		log.Error("GetAbi failed", "err", err)
		return nil, err
	}

	rawVRFContract := bind.NewBoundContract(cfg.VRFContractAddress, parsed, cfg.ChainClient, cfg.ChainClient, cfg.ChainClient)
	txManagerConfig := txmgr.Config{
		//超时时间
		ResubmissionTimeout: 5 * time.Second,
		//多久查询一次recipt
		ReceiptQueryInterval: time.Second,
		//确认位
		NumConfirmations: cfg.NumConfirmations,
		//nonce too low报错计数器
		SafeAbortNonceTooLowCount: cfg.SafeAbortNonceTooLowCount,
	}

	txManager := txmgr.NewSimpleTxManager(txManagerConfig, cfg.ChainClient)

	return &DriverEingine{
		Ctx:               ctx,
		Cfg:               cfg,
		VRFContractMethod: VRFContract,
		RawVrfContract:    rawVRFContract,
		VRfContractABI:    VRFContractAbi,
		TxMgr:             txManager,
		cancel:            cancel,
	}, nil
}

func (de *DriverEingine) UpdateGasPrice(ctx context.Context, tx *types.Transaction) (*types.Transaction, error) {
	var opts *bind.TransactOpts
	var err error
	opts, err = bind.NewKeyedTransactorWithChainID(de.Cfg.PrivateKey, de.Cfg.ChainId)
	if err != nil {
		log.Error("NewKeyedTransactorWithChainID failed", "err", err)
		return nil, err
	}

	opts.Context = ctx
	opts.Nonce = new(big.Int).SetUint64(tx.Nonce())
	opts.NoSend = true

	finalTx, err := de.RawVrfContract.RawTransact(opts, tx.Data())
	switch {
	case err == nil:
		return finalTx, nil
	case de.isMaxPriorityFeePerGasNotFoundError(err):
		log.Info("Don't support priority fee")
		opts.GasTipCap = FallbackGasTipCap
		return de.RawVrfContract.RawTransact(opts, tx.Data())
	default:
		return nil, err
	}
}

func (de *DriverEingine) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	return de.Cfg.ChainClient.SendTransaction(ctx, tx)
}

func (de *DriverEingine) isMaxPriorityFeePerGasNotFoundError(err error) bool {
	return strings.Contains(err.Error(), errMaxPriorityFeePerGasNotFound.Error())
}

func (de *DriverEingine) fulfillRandomWords(ctx context.Context, requestId *big.Int, randomList []*big.Int) (*types.Transaction, error) {
	nonce, err := de.Cfg.ChainClient.NonceAt(ctx, de.Cfg.CallerAddress, nil)
	if err != nil {
		log.Error("NonceAt failed", "err", err)
		return nil, err
	}
	opts, err := bind.NewKeyedTransactorWithChainID(de.Cfg.PrivateKey, de.Cfg.ChainId)
	if err != nil {
		log.Error("NewKeyedTransactorWithChainID failed", "err", err)
		return nil, err
	}
	opts.Context = ctx
	opts.Nonce = new(big.Int).SetUint64(nonce)
	opts.NoSend = true

	tx, err := de.VRFContractMethod.FulfillRandomWords(opts, requestId, randomList)
	switch {
	case err == nil:
		return tx, nil

	case de.isMaxPriorityFeePerGasNotFoundError(err):
		log.Info("Don't support priority fee")
		opts.GasTipCap = FallbackGasTipCap
		return de.VRFContractMethod.FulfillRandomWords(opts, requestId, randomList)

	default:
		return nil, err
	}

}

func (de *DriverEingine) FulfillRandomWords(requestId *big.Int, randomList []*big.Int) (*types.Receipt, error) {
	//构建交易
	tx, err := de.fulfillRandomWords(de.Ctx, requestId, randomList)
	if err != nil {
		log.Error("requestRandomWords failed", "err", err)
		return nil, err
	}
	//更新交易的gasprice
	updateGasprice := func(ctx context.Context) (*types.Transaction, error) {
		return de.UpdateGasPrice(ctx, tx)
	}
	receipt, err := de.TxMgr.Send(de.Ctx, updateGasprice, de.SendTransaction)
	if err != nil {
		log.Error("Send tx fail", "err", err)
		return nil, err
	}
	return receipt, nil

}
