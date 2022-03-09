package repository

import (
	"errors"
	"time"

	"github.com/sataapon/btc/internal/amt"
	"github.com/sataapon/btc/internal/config"
	"github.com/sataapon/btc/internal/sql"
	ttime "github.com/sataapon/btc/internal/time"
	"github.com/sataapon/btc/internal/wallet/entity"
)

type Wallet struct {
	db sql.DB
}

func New() *Wallet {
	return new(config.GetDB())
}

func new(db sql.DB) *Wallet {
	return &Wallet{
		db: db,
	}
}

func (w *Wallet) Insert(record *entity.Record) error {
	sqlScript := `
INSERT INTO txn (
	amount,
	acc_amount,
	date_time
) VALUES 
	($1, $2, $3);
`
	_, err := w.db.Exec(sqlScript, record.Amount.String(), record.AccAmount.String(), record.Datetime.String())
	if err != nil {
		return err
	}
	return nil
}

func (w *Wallet) GetLastRecord() (*entity.Record, bool, error) {
	sqlScript := `
SELECT
	acc_amount,
	date_time
FROM
	txn
ORDER BY
	id DESC
LIMIT 1;
`
	var (
		accAmountStr string
		datetimeStr  string
	)
	err := w.db.QueryRow(sqlScript).Scan(&accAmountStr, &datetimeStr)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, false, nil
		}
		return nil, false, err
	}
	accAmount, err := amt.New(accAmountStr)
	if err != nil {
		return nil, false, err
	}
	datetime, err := ttime.New(datetimeStr)
	if err != nil {
		return nil, false, err
	}
	return &entity.Record{
		AccAmount: accAmount,
		Datetime:  datetime,
	}, true, nil
}

func (w *Wallet) GetRecordsEveryHour(startDatetime, endDatetime ttime.Time) ([]*entity.Record, error) {
	sqlScript := `
SELECT
	id,
	acc_amount,
	date_time
FROM
	txn
WHERE
	date_time >= $1 AND date_time <= $2
`
	records, err := generateRecordsEveryHour(startDatetime, endDatetime)
	if err != nil {
		return nil, err
	}
	rows, err := w.db.Query(sqlScript, startDatetime.String(), endDatetime.String())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var (
		found = false
		idx   = 0
	)
	if rows.Next() {
		var (
			id           int
			accAmountStr string
			datetimeStr  string
		)
		err = rows.Scan(&id, &accAmountStr, &datetimeStr)
		if err != nil {
			return nil, err
		}
		accAmount, err := amt.New(accAmountStr)
		if err != nil {
			return nil, err
		}
		datetime, err := ttime.New(datetimeStr)
		if err != nil {
			return nil, err
		}

		for ; idx < len(records) && records[idx].Datetime.Before(datetime); idx++ {
		}
		if idx != 0 && id != 1 {
			record, err := w.getByID(id - 1)
			if err != nil {
				return nil, err
			}
			for i := 0; i < idx; i++ {
				records[i].AccAmount = record.AccAmount
			}
		}

		for ; idx < len(records) && !records[idx].Datetime.After(datetime); idx++ {
			records[idx].AccAmount = accAmount
		}

		found = true
	}

	for rows.Next() {
		var (
			id           int
			accAmountStr string
			datetimeStr  string
		)
		err = rows.Scan(&id, &accAmountStr, &datetimeStr)
		if err != nil {
			return nil, err
		}
		accAmount, err := amt.New(accAmountStr)
		if err != nil {
			return nil, err
		}
		datetime, err := ttime.New(datetimeStr)
		if err != nil {
			return nil, err
		}
		for ; idx < len(records) && !records[idx].Datetime.After(datetime); idx++ {
			records[idx].AccAmount = accAmount
		}
	}
	if found {
		for ; idx < len(records); idx++ {
			records[idx].AccAmount = records[idx-1].AccAmount
		}
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return records, nil
}

func (w *Wallet) getByID(id int) (*entity.Record, error) {
	sqlScript := `
SELECT
	acc_amount,
	date_time
FROM
	txn
WHERE
	id = $1
`
	var (
		accAmountStr string
		datetimeStr  string
	)
	err := w.db.QueryRow(sqlScript, id).Scan(&accAmountStr, &datetimeStr)
	if err != nil {
		return nil, err
	}
	accAmount, err := amt.New(accAmountStr)
	if err != nil {
		return nil, err
	}
	datetime, err := ttime.New(datetimeStr)
	if err != nil {
		return nil, err
	}
	return &entity.Record{
		AccAmount: accAmount,
		Datetime:  datetime,
	}, nil
}

func generateRecordsEveryHour(start ttime.Time, end ttime.Time) ([]*entity.Record, error) {
	tStart, err := time.Parse(ttime.Layout, start.String())
	if err != nil {
		return nil, err
	}
	sEnd, err := time.Parse(ttime.Layout, end.String())
	if err != nil {
		return nil, err
	}
	records := make([]*entity.Record, 0)
	tCurr := time.Date(tStart.Year(), tStart.Month(), tStart.Day(), tStart.Hour(), 0, 0, 0, tStart.Location())
	for !tCurr.After(sEnd) {
		if !tCurr.Before(tStart) {
			amount, err := amt.New("0")
			if err != nil {
				return nil, err
			}
			curr, err := ttime.New(tCurr.Format(ttime.Layout))
			if err != nil {
				return nil, err
			}
			records = append(records, &entity.Record{
				Datetime:  curr,
				AccAmount: amount,
			})
		}
		tCurr = tCurr.Add(time.Hour)
	}
	return records, nil
}
