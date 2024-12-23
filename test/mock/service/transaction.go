// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/service/transaction.go
//
// Generated by this command:
//
//	mockgen -source=./internal/service/transaction.go -destination=test/mock/./service/transaction.go
//

// Package mock_service is a generated GoMock package.
package mock_service

import (
	context "context"
	reflect "reflect"

	entity "github.com/sherwin-77/go-tix/internal/entity"
	gomock "go.uber.org/mock/gomock"
	gorm "gorm.io/gorm"
)

// MockTransactionService is a mock of TransactionService interface.
type MockTransactionService struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionServiceMockRecorder
	isgomock struct{}
}

// MockTransactionServiceMockRecorder is the mock recorder for MockTransactionService.
type MockTransactionServiceMockRecorder struct {
	mock *MockTransactionService
}

// NewMockTransactionService creates a new mock instance.
func NewMockTransactionService(ctrl *gomock.Controller) *MockTransactionService {
	mock := &MockTransactionService{ctrl: ctrl}
	mock.recorder = &MockTransactionServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransactionService) EXPECT() *MockTransactionServiceMockRecorder {
	return m.recorder
}

// CreateTransaction mocks base method.
func (m *MockTransactionService) CreateTransaction(ctx context.Context, tx *gorm.DB, saleInvoice *entity.SaleInvoice) (*entity.SnapPayment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTransaction", ctx, tx, saleInvoice)
	ret0, _ := ret[0].(*entity.SnapPayment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTransaction indicates an expected call of CreateTransaction.
func (mr *MockTransactionServiceMockRecorder) CreateTransaction(ctx, tx, saleInvoice any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTransaction", reflect.TypeOf((*MockTransactionService)(nil).CreateTransaction), ctx, tx, saleInvoice)
}