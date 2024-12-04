// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/repository/base.go
//
// Generated by this command:
//
//	mockgen -source=./internal/repository/base.go -destination=test/mock/./repository/base.go
//

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
	gorm "gorm.io/gorm"
)

// MockBaseRepository is a mock of BaseRepository interface.
type MockBaseRepository struct {
	ctrl     *gomock.Controller
	recorder *MockBaseRepositoryMockRecorder
	isgomock struct{}
}

// MockBaseRepositoryMockRecorder is the mock recorder for MockBaseRepository.
type MockBaseRepositoryMockRecorder struct {
	mock *MockBaseRepository
}

// NewMockBaseRepository creates a new mock instance.
func NewMockBaseRepository(ctrl *gomock.Controller) *MockBaseRepository {
	mock := &MockBaseRepository{ctrl: ctrl}
	mock.recorder = &MockBaseRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBaseRepository) EXPECT() *MockBaseRepositoryMockRecorder {
	return m.recorder
}

// BeginTransaction mocks base method.
func (m *MockBaseRepository) BeginTransaction() *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BeginTransaction")
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// BeginTransaction indicates an expected call of BeginTransaction.
func (mr *MockBaseRepositoryMockRecorder) BeginTransaction() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeginTransaction", reflect.TypeOf((*MockBaseRepository)(nil).BeginTransaction))
}

// Commit mocks base method.
func (m *MockBaseRepository) Commit(tx *gorm.DB) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Commit", tx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Commit indicates an expected call of Commit.
func (mr *MockBaseRepositoryMockRecorder) Commit(tx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockBaseRepository)(nil).Commit), tx)
}

// Rollback mocks base method.
func (m *MockBaseRepository) Rollback(tx *gorm.DB) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Rollback", tx)
}

// Rollback indicates an expected call of Rollback.
func (mr *MockBaseRepositoryMockRecorder) Rollback(tx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rollback", reflect.TypeOf((*MockBaseRepository)(nil).Rollback), tx)
}

// SingleTransaction mocks base method.
func (m *MockBaseRepository) SingleTransaction() *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SingleTransaction")
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// SingleTransaction indicates an expected call of SingleTransaction.
func (mr *MockBaseRepositoryMockRecorder) SingleTransaction() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SingleTransaction", reflect.TypeOf((*MockBaseRepository)(nil).SingleTransaction))
}

// WithTransaction mocks base method.
func (m *MockBaseRepository) WithTransaction(fn func(*gorm.DB) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithTransaction", fn)
	ret0, _ := ret[0].(error)
	return ret0
}

// WithTransaction indicates an expected call of WithTransaction.
func (mr *MockBaseRepositoryMockRecorder) WithTransaction(fn any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithTransaction", reflect.TypeOf((*MockBaseRepository)(nil).WithTransaction), fn)
}
