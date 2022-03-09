package usecase

import (
	"errors"
	"net/http"
	"sync"

	"github.com/sataapon/btc/internal/amt"
	"github.com/sataapon/btc/internal/meta"
	"github.com/sataapon/btc/internal/time"
	"github.com/sataapon/btc/internal/wallet/entity"
	"github.com/sataapon/btc/internal/wallet/repository"
)

type walletRepository interface {
	Insert(record *entity.Record) error
	GetLastRecord() (*entity.Record, bool, error)
	GetRecordsEveryHour(startDatetime, endDatetime time.Time) ([]*entity.Record, error)
}

type Wallet struct {
	walletRepository walletRepository
	mu               sync.Mutex
}

func New() *Wallet {
	return new(repository.New())
}

func new(walletRepository walletRepository) *Wallet {
	return &Wallet{
		walletRepository: walletRepository,
	}
}

func (w *Wallet) SaveRecord(datetime time.Time, amount amt.Amount) error {
	if amount.IsZero() || amount.IsNegative() {
		return meta.NewError(http.StatusPreconditionFailed, errors.New("amount must be positive number"))
	}
	w.mu.Lock()
	defer w.mu.Unlock()
	lastRecord, found, err := w.walletRepository.GetLastRecord()
	if err != nil {
		return meta.NewError(http.StatusInternalServerError, err)
	}
	accAmount := amount
	if found {
		if !datetime.After(lastRecord.Datetime) {
			return meta.NewError(http.StatusPreconditionFailed, errors.New("date time must greater than last date time"))
		}
		accAmount = lastRecord.AccAmount.Plus(amount)
	}
	err = w.walletRepository.Insert(&entity.Record{
		Amount:    amount,
		AccAmount: accAmount,
		Datetime:  datetime,
	})
	if err != nil {
		return meta.NewError(http.StatusInternalServerError, err)
	}
	return nil
}

func (w *Wallet) GetHistory(startDatetime, endDatetime time.Time) ([]*entity.Record, error) {
	if endDatetime.Before(startDatetime) {
		return nil, meta.NewError(http.StatusPreconditionFailed, errors.New("end date time must less than or equal to start date time"))
	}
	records, err := w.walletRepository.GetRecordsEveryHour(startDatetime, endDatetime)
	if err != nil {
		return nil, meta.NewError(http.StatusInternalServerError, err)
	}
	return records, nil
}
