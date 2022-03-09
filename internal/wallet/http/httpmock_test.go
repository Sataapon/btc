// Code generated by MockGen. DO NOT EDIT.
// Source: ./http.go

// Package http is a generated GoMock package.
package http

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	amt "github.com/sataapon/btc/internal/amt"
	time "github.com/sataapon/btc/internal/time"
	entity "github.com/sataapon/btc/internal/wallet/entity"
)

// MockwalletUsecase is a mock of walletUsecase interface.
type MockwalletUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockwalletUsecaseMockRecorder
}

// MockwalletUsecaseMockRecorder is the mock recorder for MockwalletUsecase.
type MockwalletUsecaseMockRecorder struct {
	mock *MockwalletUsecase
}

// NewMockwalletUsecase creates a new mock instance.
func NewMockwalletUsecase(ctrl *gomock.Controller) *MockwalletUsecase {
	mock := &MockwalletUsecase{ctrl: ctrl}
	mock.recorder = &MockwalletUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockwalletUsecase) EXPECT() *MockwalletUsecaseMockRecorder {
	return m.recorder
}

// GetHistory mocks base method.
func (m *MockwalletUsecase) GetHistory(startDatetime, stopDatetime time.Time) ([]*entity.Record, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHistory", startDatetime, stopDatetime)
	ret0, _ := ret[0].([]*entity.Record)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHistory indicates an expected call of GetHistory.
func (mr *MockwalletUsecaseMockRecorder) GetHistory(startDatetime, stopDatetime interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHistory", reflect.TypeOf((*MockwalletUsecase)(nil).GetHistory), startDatetime, stopDatetime)
}

// SaveRecord mocks base method.
func (m *MockwalletUsecase) SaveRecord(datetime time.Time, amount amt.Amount) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveRecord", datetime, amount)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveRecord indicates an expected call of SaveRecord.
func (mr *MockwalletUsecaseMockRecorder) SaveRecord(datetime, amount interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveRecord", reflect.TypeOf((*MockwalletUsecase)(nil).SaveRecord), datetime, amount)
}
