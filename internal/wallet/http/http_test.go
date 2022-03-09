package http

import (
	"testing"

	"github.com/golang/mock/gomock"
)

func TestHTTPServer_SaveRecord(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockWalletUsecase := NewMockwalletUsecase(ctrl)
	httpServer := newServer(mockWalletUsecase)
	// TODO Add unit-test
	_ = httpServer
}

func TestHTTPServer_GetHistory(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockWalletUsecase := NewMockwalletUsecase(ctrl)
	httpServer := newServer(mockWalletUsecase)
	// TODO Add unit-test
	_ = httpServer
}
