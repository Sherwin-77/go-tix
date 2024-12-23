// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/repository/event_approval.go
//
// Generated by this command:
//
//	mockgen -source=./internal/repository/event_approval.go -destination=test/mock/./repository/event_approval.go
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

// MockEventApprovalRepository is a mock of EventApprovalRepository interface.
type MockEventApprovalRepository struct {
	ctrl     *gomock.Controller
	recorder *MockEventApprovalRepositoryMockRecorder
	isgomock struct{}
}

// MockEventApprovalRepositoryMockRecorder is the mock recorder for MockEventApprovalRepository.
type MockEventApprovalRepositoryMockRecorder struct {
	mock *MockEventApprovalRepository
}

// NewMockEventApprovalRepository creates a new mock instance.
func NewMockEventApprovalRepository(ctrl *gomock.Controller) *MockEventApprovalRepository {
	mock := &MockEventApprovalRepository{ctrl: ctrl}
	mock.recorder = &MockEventApprovalRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEventApprovalRepository) EXPECT() *MockEventApprovalRepositoryMockRecorder {
	return m.recorder
}

// BeginTransaction mocks base method.
func (m *MockEventApprovalRepository) BeginTransaction() *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BeginTransaction")
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// BeginTransaction indicates an expected call of BeginTransaction.
func (mr *MockEventApprovalRepositoryMockRecorder) BeginTransaction() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeginTransaction", reflect.TypeOf((*MockEventApprovalRepository)(nil).BeginTransaction))
}

// Commit mocks base method.
func (m *MockEventApprovalRepository) Commit(tx *gorm.DB) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Commit", tx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Commit indicates an expected call of Commit.
func (mr *MockEventApprovalRepositoryMockRecorder) Commit(tx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockEventApprovalRepository)(nil).Commit), tx)
}

// CreateEventApproval mocks base method.
func (m *MockEventApprovalRepository) CreateEventApproval(ctx context.Context, tx *gorm.DB, eventApproval *entity.EventApproval) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEventApproval", ctx, tx, eventApproval)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateEventApproval indicates an expected call of CreateEventApproval.
func (mr *MockEventApprovalRepositoryMockRecorder) CreateEventApproval(ctx, tx, eventApproval any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEventApproval", reflect.TypeOf((*MockEventApprovalRepository)(nil).CreateEventApproval), ctx, tx, eventApproval)
}

// DeleteEventApproval mocks base method.
func (m *MockEventApprovalRepository) DeleteEventApproval(ctx context.Context, tx *gorm.DB, eventApproval *entity.EventApproval) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteEventApproval", ctx, tx, eventApproval)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteEventApproval indicates an expected call of DeleteEventApproval.
func (mr *MockEventApprovalRepositoryMockRecorder) DeleteEventApproval(ctx, tx, eventApproval any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteEventApproval", reflect.TypeOf((*MockEventApprovalRepository)(nil).DeleteEventApproval), ctx, tx, eventApproval)
}

// GetEventApprovalByID mocks base method.
func (m *MockEventApprovalRepository) GetEventApprovalByID(ctx context.Context, tx *gorm.DB, id string) (*entity.EventApproval, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEventApprovalByID", ctx, tx, id)
	ret0, _ := ret[0].(*entity.EventApproval)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEventApprovalByID indicates an expected call of GetEventApprovalByID.
func (mr *MockEventApprovalRepositoryMockRecorder) GetEventApprovalByID(ctx, tx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEventApprovalByID", reflect.TypeOf((*MockEventApprovalRepository)(nil).GetEventApprovalByID), ctx, tx, id)
}

// GetEventApprovals mocks base method.
func (m *MockEventApprovalRepository) GetEventApprovals(ctx context.Context, tx *gorm.DB) ([]entity.EventApproval, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEventApprovals", ctx, tx)
	ret0, _ := ret[0].([]entity.EventApproval)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEventApprovals indicates an expected call of GetEventApprovals.
func (mr *MockEventApprovalRepositoryMockRecorder) GetEventApprovals(ctx, tx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEventApprovals", reflect.TypeOf((*MockEventApprovalRepository)(nil).GetEventApprovals), ctx, tx)
}

// GetUserEventApprovals mocks base method.
func (m *MockEventApprovalRepository) GetUserEventApprovals(ctx context.Context, tx *gorm.DB, userID string) ([]entity.EventApproval, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserEventApprovals", ctx, tx, userID)
	ret0, _ := ret[0].([]entity.EventApproval)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserEventApprovals indicates an expected call of GetUserEventApprovals.
func (mr *MockEventApprovalRepositoryMockRecorder) GetUserEventApprovals(ctx, tx, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserEventApprovals", reflect.TypeOf((*MockEventApprovalRepository)(nil).GetUserEventApprovals), ctx, tx, userID)
}

// Rollback mocks base method.
func (m *MockEventApprovalRepository) Rollback(tx *gorm.DB) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Rollback", tx)
}

// Rollback indicates an expected call of Rollback.
func (mr *MockEventApprovalRepositoryMockRecorder) Rollback(tx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rollback", reflect.TypeOf((*MockEventApprovalRepository)(nil).Rollback), tx)
}

// SingleTransaction mocks base method.
func (m *MockEventApprovalRepository) SingleTransaction() *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SingleTransaction")
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// SingleTransaction indicates an expected call of SingleTransaction.
func (mr *MockEventApprovalRepositoryMockRecorder) SingleTransaction() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SingleTransaction", reflect.TypeOf((*MockEventApprovalRepository)(nil).SingleTransaction))
}

// UpdateEventApproval mocks base method.
func (m *MockEventApprovalRepository) UpdateEventApproval(ctx context.Context, tx *gorm.DB, eventApproval *entity.EventApproval) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateEventApproval", ctx, tx, eventApproval)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateEventApproval indicates an expected call of UpdateEventApproval.
func (mr *MockEventApprovalRepositoryMockRecorder) UpdateEventApproval(ctx, tx, eventApproval any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateEventApproval", reflect.TypeOf((*MockEventApprovalRepository)(nil).UpdateEventApproval), ctx, tx, eventApproval)
}

// WithPreloads mocks base method.
func (m *MockEventApprovalRepository) WithPreloads(tx *gorm.DB, preloads map[string][]any) *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithPreloads", tx, preloads)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// WithPreloads indicates an expected call of WithPreloads.
func (mr *MockEventApprovalRepositoryMockRecorder) WithPreloads(tx, preloads any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithPreloads", reflect.TypeOf((*MockEventApprovalRepository)(nil).WithPreloads), tx, preloads)
}

// WithTransaction mocks base method.
func (m *MockEventApprovalRepository) WithTransaction(fn func(*gorm.DB) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithTransaction", fn)
	ret0, _ := ret[0].(error)
	return ret0
}

// WithTransaction indicates an expected call of WithTransaction.
func (mr *MockEventApprovalRepositoryMockRecorder) WithTransaction(fn any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithTransaction", reflect.TypeOf((*MockEventApprovalRepository)(nil).WithTransaction), fn)
}
