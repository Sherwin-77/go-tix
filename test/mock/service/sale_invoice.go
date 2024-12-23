// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/service/sale_invoice.go
//
// Generated by this command:
//
//	mockgen -source=./internal/service/sale_invoice.go -destination=test/mock/./service/sale_invoice.go
//

// Package mock_service is a generated GoMock package.
package mock_service

import (
	context "context"
	url "net/url"
	reflect "reflect"

	domain "github.com/sherwin-77/go-tix/internal/domain"
	entity "github.com/sherwin-77/go-tix/internal/entity"
	dto "github.com/sherwin-77/go-tix/internal/http/dto"
	response "github.com/sherwin-77/go-tix/pkg/response"
	gomock "go.uber.org/mock/gomock"
)

// MockSaleInvoiceService is a mock of SaleInvoiceService interface.
type MockSaleInvoiceService struct {
	ctrl     *gomock.Controller
	recorder *MockSaleInvoiceServiceMockRecorder
	isgomock struct{}
}

// MockSaleInvoiceServiceMockRecorder is the mock recorder for MockSaleInvoiceService.
type MockSaleInvoiceServiceMockRecorder struct {
	mock *MockSaleInvoiceService
}

// NewMockSaleInvoiceService creates a new mock instance.
func NewMockSaleInvoiceService(ctrl *gomock.Controller) *MockSaleInvoiceService {
	mock := &MockSaleInvoiceService{ctrl: ctrl}
	mock.recorder = &MockSaleInvoiceServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSaleInvoiceService) EXPECT() *MockSaleInvoiceServiceMockRecorder {
	return m.recorder
}

// Bill mocks base method.
func (m *MockSaleInvoiceService) Bill(ctx context.Context, request dto.CheckoutRequest) (*domain.InvoicePricing, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Bill", ctx, request)
	ret0, _ := ret[0].(*domain.InvoicePricing)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Bill indicates an expected call of Bill.
func (mr *MockSaleInvoiceServiceMockRecorder) Bill(ctx, request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Bill", reflect.TypeOf((*MockSaleInvoiceService)(nil).Bill), ctx, request)
}

// Checkout mocks base method.
func (m *MockSaleInvoiceService) Checkout(ctx context.Context, request dto.CheckoutRequest, userID string) (*domain.CheckoutData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Checkout", ctx, request, userID)
	ret0, _ := ret[0].(*domain.CheckoutData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Checkout indicates an expected call of Checkout.
func (mr *MockSaleInvoiceServiceMockRecorder) Checkout(ctx, request, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Checkout", reflect.TypeOf((*MockSaleInvoiceService)(nil).Checkout), ctx, request, userID)
}

// GetSaleInvoiceByID mocks base method.
func (m *MockSaleInvoiceService) GetSaleInvoiceByID(ctx context.Context, id string) (*entity.SaleInvoice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSaleInvoiceByID", ctx, id)
	ret0, _ := ret[0].(*entity.SaleInvoice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSaleInvoiceByID indicates an expected call of GetSaleInvoiceByID.
func (mr *MockSaleInvoiceServiceMockRecorder) GetSaleInvoiceByID(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSaleInvoiceByID", reflect.TypeOf((*MockSaleInvoiceService)(nil).GetSaleInvoiceByID), ctx, id)
}

// GetSaleInvoices mocks base method.
func (m *MockSaleInvoiceService) GetSaleInvoices(ctx context.Context, queryParams url.Values) ([]entity.SaleInvoice, *response.Meta, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSaleInvoices", ctx, queryParams)
	ret0, _ := ret[0].([]entity.SaleInvoice)
	ret1, _ := ret[1].(*response.Meta)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetSaleInvoices indicates an expected call of GetSaleInvoices.
func (mr *MockSaleInvoiceServiceMockRecorder) GetSaleInvoices(ctx, queryParams any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSaleInvoices", reflect.TypeOf((*MockSaleInvoiceService)(nil).GetSaleInvoices), ctx, queryParams)
}

// GetUserSaleInvoiceByID mocks base method.
func (m *MockSaleInvoiceService) GetUserSaleInvoiceByID(ctx context.Context, id, userID string) (*entity.SaleInvoice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserSaleInvoiceByID", ctx, id, userID)
	ret0, _ := ret[0].(*entity.SaleInvoice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserSaleInvoiceByID indicates an expected call of GetUserSaleInvoiceByID.
func (mr *MockSaleInvoiceServiceMockRecorder) GetUserSaleInvoiceByID(ctx, id, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserSaleInvoiceByID", reflect.TypeOf((*MockSaleInvoiceService)(nil).GetUserSaleInvoiceByID), ctx, id, userID)
}

// GetUserSaleInvoices mocks base method.
func (m *MockSaleInvoiceService) GetUserSaleInvoices(ctx context.Context, queryParams url.Values, userID string) ([]entity.SaleInvoice, *response.Meta, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserSaleInvoices", ctx, queryParams, userID)
	ret0, _ := ret[0].([]entity.SaleInvoice)
	ret1, _ := ret[1].(*response.Meta)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetUserSaleInvoices indicates an expected call of GetUserSaleInvoices.
func (mr *MockSaleInvoiceServiceMockRecorder) GetUserSaleInvoices(ctx, queryParams, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserSaleInvoices", reflect.TypeOf((*MockSaleInvoiceService)(nil).GetUserSaleInvoices), ctx, queryParams, userID)
}
