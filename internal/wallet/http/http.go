package http

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/sataapon/btc/internal/amt"
	"github.com/sataapon/btc/internal/meta"
	"github.com/sataapon/btc/internal/time"
	"github.com/sataapon/btc/internal/wallet/entity"
	"github.com/sataapon/btc/internal/wallet/usecase"
)

func NewServer(addr string) *http.Server {
	var (
		httpServer = newServer(usecase.New())
		router     = mux.NewRouter()
	)

	router.HandleFunc("/", httpServer.SaveRecord).Methods("POST")
	router.HandleFunc("/", httpServer.GetHistory).Methods("GET")

	return &http.Server{
		Addr:    addr,
		Handler: router,
	}
}

type walletUsecase interface {
	SaveRecord(datetime time.Time, amount amt.Amount) error
	GetHistory(startDatetime, stopDatetime time.Time) ([]*entity.Record, error)
}

type httpServer struct {
	walletUsecase walletUsecase
}

func newServer(walletUsecase walletUsecase) *httpServer {
	return &httpServer{
		walletUsecase: walletUsecase,
	}
}

func (s *httpServer) SaveRecord(w http.ResponseWriter, r *http.Request) {
	req := &InsertRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		if metaErr, ok := meta.IsError(err); ok {
			http.Error(w, metaErr.Error(), metaErr.HTTPStatus())
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = s.walletUsecase.SaveRecord(req.Record.Datetime, req.Record.Amount)
	if err != nil {
		if metaErr, ok := meta.IsError(err); ok {
			http.Error(w, metaErr.Error(), metaErr.HTTPStatus())
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res := &InsertResponse{Success: true}
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		if metaErr, ok := meta.IsError(err); ok {
			http.Error(w, metaErr.Error(), metaErr.HTTPStatus())
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *httpServer) GetHistory(w http.ResponseWriter, r *http.Request) {
	req := &GetHistoryRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		if metaErr, ok := meta.IsError(err); ok {
			http.Error(w, metaErr.Error(), metaErr.HTTPStatus())
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	records, err := s.walletUsecase.GetHistory(req.StartDatetime, req.StopDatetime)
	if err != nil {
		if metaErr, ok := meta.IsError(err); ok {
			http.Error(w, metaErr.Error(), metaErr.HTTPStatus())
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res := &GetHistoryResponse{Records: convertRecords(records)}
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		if metaErr, ok := meta.IsError(err); ok {
			http.Error(w, metaErr.Error(), metaErr.HTTPStatus())
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
