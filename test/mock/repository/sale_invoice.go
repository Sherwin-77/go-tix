// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/repository/sale_invoice.go
//
// Generated by this command:
//
//	mockgen -source=./internal/repository/sale_invoice.go -destination=test/mock/./repository/sale_invoice.go
//

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	context "context"
	reflect "reflect"

	entity "github.com/sherwin-77/go-tix/internal/entity"
	gomock "go.uber.org/mock/gomock"
	gorm "gorm.io/gorm"
)

// MockSaleInvoiceRepository is a mock of SaleInvoiceRepository interface.
type MockSaleInvoiceRepository struct {
	ctrl     *gomock.Controller
	recorder *MockSaleInvoiceRepositoryMockRecorder
	isgomock struct{}
}

// MockSaleInvoiceRepositoryMockRecorder is the mock recorder for MockSaleInvoiceRepository.
type MockSaleInvoiceRepositoryMockRecorder struct {
	mock *MockSaleInvoiceRepository
}

// NewMockSaleInvoiceRepository creates a new mock instance.
func NewMockSaleInvoiceRepository(ctrl *gomock.Controller) *MockSaleInvoiceRepository {
	mock := &MockSaleInvoiceRepository{ctrl: ctrl}
	mock.recorder = &MockSaleInvoiceRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSaleInvoiceRepository) EXPECT() *MockSaleInvoiceRepositoryMockRecorder {
	return m.recorder
}

// BeginTransaction mocks base method.
func (m *MockSaleInvoiceRepository) BeginTransaction() *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BeginTransaction")
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// BeginTransaction indicates an expected call of BeginTransaction.
func (mr *MockSaleInvoiceRepositoryMockRecorder) BeginTransaction() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeginTransaction", reflect.TypeOf((*MockSaleInvoiceRepository)(nil).BeginTransaction))
}

// Commit mocks base method.
func (m *MockSaleInvoiceRepository) Commit(tx *gorm.DB) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Commit", tx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Commit indicates an expected call of Commit.
func (mr *MockSaleInvoiceRepositoryMockRecorder) Commit(tx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockSaleInvoiceRepository)(nil).Commit), tx)
}

// CreateSaleInvoice mocks base method.
func (m *MockSaleInvoiceRepository) CreateSaleInvoice(ctx context.Context, tx *gorm.DB, saleInvoice *entity.SaleInvoice) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSaleInvoice", ctx, tx, saleInvoice)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateSaleInvoice indicates an expected call of CreateSaleInvoice.
func (mr *MockSaleInvoiceRepositoryMockRecorder) CreateSaleInvoice(ctx, tx, saleInvoice any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSaleInvoice", reflect.TypeOf((*MockSaleInvoiceRepository)(nil).CreateSaleInvoice), ctx, tx, saleInvoice)
}

// DeleteSaleInvoice mocks base method.
func (m *MockSaleInvoiceRepository) DeleteSaleInvoice(ctx context.Context, tx *gorm.DB, saleInvoice *entity.SaleInvoice) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSaleInvoice", ctx, tx, saleInvoice)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSaleInvoice indicates an expected call of DeleteSaleInvoice.
func (mr *MockSaleInvoiceRepositoryMockRecorder) DeleteSaleInvoice(ctx, tx, saleInvoice any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSaleInvoice", reflect.TypeOf((*MockSaleInvoiceRepository)(nil).DeleteSaleInvoice), ctx, tx, saleInvoice)
}

// GetSaleInvoiceByID mocks base method.
func (m *MockSaleInvoiceRepository) GetSaleInvoiceByID(ctx context.Context, tx *gorm.DB, id string) (*entity.SaleInvoice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSaleInvoiceByID", ctx, tx, id)
	ret0, _ := ret[0].(*entity.SaleInvoice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSaleInvoiceByID indicates an expected call of GetSaleInvoiceByID.
func (mr *MockSaleInvoiceRepositoryMockRecorder) GetSaleInvoiceByID(ctx, tx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSaleInvoiceByID", reflect.TypeOf((*MockSaleInvoiceRepository)(nil).GetSaleInvoiceByID), ctx, tx, id)
}

// GetSaleInvoices mocks base method.
func (m *MockSaleInvoiceRepository) GetSaleInvoices(ctx context.Context, tx *gorm.DB) ([]entity.SaleInvoice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSaleInvoices", ctx, tx)
	ret0, _ := ret[0].([]entity.SaleInvoice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSaleInvoices indicates an expected call of GetSaleInvoices.
func (mr *MockSaleInvoiceRepositoryMockRecorder) GetSaleInvoices(ctx, tx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSaleInvoices", reflect.TypeOf((*MockSaleInvoiceRepository)(nil).GetSaleInvoices), ctx, tx)
}

// GetUserSaleInvoices mocks base method.
func (m *MockSaleInvoiceRepository) GetUserSaleInvoices(ctx context.Context, tx *gorm.DB, userID string) ([]entity.SaleInvoice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserSaleInvoices", ctx, tx, userID)
	ret0, _ := ret[0].([]entity.SaleInvoice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserSaleInvoices indicates an expected call of GetUserSaleInvoices.
func (mr *MockSaleInvoiceRepositoryMockRecorder) GetUserSaleInvoices(ctx, tx, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserSaleInvoices", reflect.TypeOf((*MockSaleInvoiceRepository)(nil).GetUserSaleInvoices), ctx, tx, userID)
}

// Rollback mocks base method.
func (m *MockSaleInvoiceRepository) Rollback(tx *gorm.DB) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Rollback", tx)
}

// Rollback indicates an expected call of Rollback.
func (mr *MockSaleInvoiceRepositoryMockRecorder) Rollback(tx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rollback", reflect.TypeOf((*MockSaleInvoiceRepository)(nil).Rollback), tx)
}

// SaleInvoiceNumberExists mocks base method.
func (m *MockSaleInvoiceRepository) SaleInvoiceNumberExists(ctx context.Context, tx *gorm.DB, number string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaleInvoiceNumberExists", ctx, tx, number)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SaleInvoiceNumberExists indicates an expected call of SaleInvoiceNumberExists.
func (mr *MockSaleInvoiceRepositoryMockRecorder) SaleInvoiceNumberExists(ctx, tx, number any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaleInvoiceNumberExists", reflect.TypeOf((*MockSaleInvoiceRepository)(nil).SaleInvoiceNumberExists), ctx, tx, number)
}

// SingleTransaction mocks base method.
func (m *MockSaleInvoiceRepository) SingleTransaction() *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SingleTransaction")
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// SingleTransaction indicates an expected call of SingleTransaction.
func (mr *MockSaleInvoiceRepositoryMockRecorder) SingleTransaction() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SingleTransaction", reflect.TypeOf((*MockSaleInvoiceRepository)(nil).SingleTransaction))
}

// UpdateSaleInvoice mocks base method.
func (m *MockSaleInvoiceRepository) UpdateSaleInvoice(ctx context.Context, tx *gorm.DB, saleInvoice *entity.SaleInvoice) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSaleInvoice", ctx, tx, saleInvoice)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateSaleInvoice indicates an expected call of UpdateSaleInvoice.
func (mr *MockSaleInvoiceRepositoryMockRecorder) UpdateSaleInvoice(ctx, tx, saleInvoice any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSaleInvoice", reflect.TypeOf((*MockSaleInvoiceRepository)(nil).UpdateSaleInvoice), ctx, tx, saleInvoice)
}

// WithPreloads mocks base method.
func (m *MockSaleInvoiceRepository) WithPreloads(tx *gorm.DB, preloads map[string][]any) *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithPreloads", tx, preloads)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// WithPreloads indicates an expected call of WithPreloads.
func (mr *MockSaleInvoiceRepositoryMockRecorder) WithPreloads(tx, preloads any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithPreloads", reflect.TypeOf((*MockSaleInvoiceRepository)(nil).WithPreloads), tx, preloads)
}

// WithTransaction mocks base method.
func (m *MockSaleInvoiceRepository) WithTransaction(fn func(*gorm.DB) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithTransaction", fn)
	ret0, _ := ret[0].(error)
	return ret0
}

// WithTransaction indicates an expected call of WithTransaction.
func (mr *MockSaleInvoiceRepositoryMockRecorder) WithTransaction(fn any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithTransaction", reflect.TypeOf((*MockSaleInvoiceRepository)(nil).WithTransaction), fn)
}
