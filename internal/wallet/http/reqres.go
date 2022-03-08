package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/sataapon/btc/internal/amt"
	"github.com/sataapon/btc/internal/meta"
	"github.com/sataapon/btc/internal/time"
)

type Record struct {
	Datetime time.Time
	Amount   amt.Amount
}

func (r *Record) UnmarshalJSON(b []byte) error {
	st := &struct {
		Datetime string  `json:"datetime"`
		Amount   float64 `json:"amount"`
	}{}
	err := json.Unmarshal(b, st)
	if err != nil {
		return meta.NewError(http.StatusBadRequest, err)
	}
	datetime, err := time.New(st.Datetime)
	if err != nil {
		return meta.NewError(http.StatusBadRequest, err)
	}
	r.Datetime = datetime
	amount, err := amt.New(strconv.FormatFloat(st.Amount, 'f', -1, 64))
	if err != nil {
		return meta.NewError(http.StatusBadRequest, err)
	}
	r.Amount = amount
	return nil
}

func (r *Record) MarshalJSON() ([]byte, error) {
	st := &struct {
		Datetime string  `json:"datetime"`
		Amount   float64 `json:"amount"`
	}{}
	st.Datetime = r.Datetime.String()
	amount, err := r.Amount.ToFloat()
	if err != nil {
		return nil, meta.NewError(http.StatusInternalServerError, err)
	}
	st.Amount = amount
	bytes, err := json.Marshal(st)
	if err != nil {
		return nil, meta.NewError(http.StatusInternalServerError, err)
	}
	return bytes, nil
}

type InsertRequest struct {
	Record *Record
}

func (iq *InsertRequest) UnmarshalJSON(b []byte) error {
	record := &Record{}
	err := json.Unmarshal(b, record)
	if err != nil {
		return meta.NewError(http.StatusBadRequest, err)
	}
	iq.Record = record
	return nil
}

type InsertResponse struct {
	Success bool `json:"success"`
}

type GetHistoryRequest struct {
	StartDatetime time.Time
	StopDatetime  time.Time
}

func (g *GetHistoryRequest) UnmarshalJSON(b []byte) error {
	st := &struct {
		StartDatetime string `json:"startDatetime"`
		StopDatetime  string `json:"EndDatetime"`
	}{}
	err := json.Unmarshal(b, st)
	if err != nil {
		return meta.NewError(http.StatusBadRequest, err)
	}
	startDatetime, err := time.New(st.StartDatetime)
	if err != nil {
		return meta.NewError(http.StatusBadRequest, err)
	}
	g.StartDatetime = startDatetime
	stopDatetime, err := time.New(st.StopDatetime)
	if err != nil {
		return meta.NewError(http.StatusBadRequest, err)
	}
	g.StopDatetime = stopDatetime
	return nil
}

type GetHistoryResponse struct {
	Records []*Record
}

func (gp *GetHistoryResponse) MarshalJSON() ([]byte, error) {
	bytes, err := json.Marshal(gp.Records)
	if err != nil {
		return nil, meta.NewError(http.StatusInternalServerError, err)
	}
	return bytes, err
}
