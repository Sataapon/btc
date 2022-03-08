package http

import (
	"github.com/sataapon/btc/internal/wallet/entity"
)

func convertRecords(rs []*entity.Record) []*Record {
	var records = make([]*Record, len(rs))
	for idx, r := range rs {
		records[idx] = &Record{
			Datetime: r.Datetime,
			Amount:   r.AccAmount,
		}
	}
	return records
}
