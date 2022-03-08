package entity

import (
	"github.com/sataapon/btc/internal/amt"
	"github.com/sataapon/btc/internal/time"
)

type Record struct {
	Datetime  time.Time
	Amount    amt.Amount
	AccAmount amt.Amount
}
