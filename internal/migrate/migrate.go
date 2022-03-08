package migrate

import (
	"github.com/sataapon/btc/internal/sql"

	"github.com/sataapon/btc/internal/config"
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

func (w *Wallet) CreateTable() error {
	sqlScript := `
DROP TABLE IF EXISTS txn;

DROP INDEX IF EXISTS idx_txn_date_time;

CREATE TABLE txn (
	id SERIAL PRIMARY KEY,
	amount TEXT NOT NULL,
	acc_amount TEXT NOT NULL,
	date_time TIMESTAMPTZ UNIQUE NOT NULL
);

CREATE INDEX idx_txn_date_time 
ON txn(date_time);
`
	_, err := w.db.Exec(sqlScript)
	if err != nil {
		return err
	}
	return nil
}
