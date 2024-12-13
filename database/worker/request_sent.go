package worker

import (
	"errors"
	"fmt"
	"math/big"

	"gorm.io/gorm"

	"github.com/google/uuid"

	"github.com/ethereum/go-ethereum/common"
)

type RequestSend struct {
	GUID       uuid.UUID      `gorm:"primaryKey" json:"guid"`
	RequestId  *big.Int       `json:"request_id" gorm:"serializer:u256"`
	VrfAddress common.Address `json:"vrf_address" gorm:"serializer:bytes"`
	NumWords   *big.Int       `json:"num_words" gorm:"serializer:u256"`
	Status     uint8          `json:"status"` // 0:扫到合约事件,1:已经上传随机数
	Timestamp  uint64
}

type RequestSendView interface {
	QueryUnHandleRequestSendList() ([]RequestSend, error)
}

type RequestSendDB interface {
	RequestSendView

	StoreRequestSend([]RequestSend) error
	MarkRequestSendFinish(RequestSend) error
}

type requestSendDB struct {
	gorm *gorm.DB
}

func NewRequestSendDB(db *gorm.DB) RequestSendDB {
	return &requestSendDB{gorm: db}
}

func (db requestSendDB) StoreRequestSend(RequestSendList []RequestSend) error {
	result := db.gorm.CreateInBatches(&RequestSendList, len(RequestSendList))
	return result.Error
}

func (db requestSendDB) QueryUnHandleRequestSendList() ([]RequestSend, error) {
	var requestSendList []RequestSend
	err := db.gorm.Table("request_send").Where("status = ?", 0).Find(&requestSendList).Error
	if err != nil {
		return nil, fmt.Errorf("query unhandle request sent list failed: %w", err)
	}
	return requestSendList, nil
}
func (db requestSendDB) MarkRequestSendFinish(requestSend RequestSend) error {
	var requestSendSingle = RequestSend{}
	result := db.gorm.Table("request_send").Where(&RequestSend{GUID: requestSend.GUID}).Take(&requestSendSingle)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil
		}
		return result.Error
	}
	requestSendSingle.Status = 1
	err := db.gorm.Table("request_send").Save(&requestSendSingle).Error
	if err != nil {
		return err
	}
	return nil
}
