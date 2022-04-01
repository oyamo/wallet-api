// Code generated by MockGen. DO NOT EDIT.
// Source: redis_repo.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/oyamo/wallet-api/internal/models"
)

// MockRedisRepository is a mock of RedisRepository interface.
type MockRedisRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRedisRepositoryMockRecorder
}

// MockRedisRepositoryMockRecorder is the mock recorder for MockRedisRepository.
type MockRedisRepositoryMockRecorder struct {
	mock *MockRedisRepository
}

// NewMockRedisRepository creates a new mock instance.
func NewMockRedisRepository(ctrl *gomock.Controller) *MockRedisRepository {
	mock := &MockRedisRepository{ctrl: ctrl}
	mock.recorder = &MockRedisRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRedisRepository) EXPECT() *MockRedisRepositoryMockRecorder {
	return m.recorder
}

// DeleteTxCtx mocks base method.
func (m *MockRedisRepository) DeleteTxCtx(ctx context.Context, key string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTxCtx", ctx, key)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTxCtx indicates an expected call of DeleteTxCtx.
func (mr *MockRedisRepositoryMockRecorder) DeleteTxCtx(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTxCtx", reflect.TypeOf((*MockRedisRepository)(nil).DeleteTxCtx), ctx, key)
}

// GetByIDCtx mocks base method.
func (m *MockRedisRepository) GetByIDCtx(ctx context.Context, key string) (*models.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByIDCtx", ctx, key)
	ret0, _ := ret[0].(*models.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByIDCtx indicates an expected call of GetByIDCtx.
func (mr *MockRedisRepositoryMockRecorder) GetByIDCtx(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByIDCtx", reflect.TypeOf((*MockRedisRepository)(nil).GetByIDCtx), ctx, key)
}

// SetTxCtx mocks base method.
func (m *MockRedisRepository) SetTxCtx(ctx context.Context, key string, seconds int, user *models.Transaction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetTxCtx", ctx, key, seconds, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetTxCtx indicates an expected call of SetTxCtx.
func (mr *MockRedisRepositoryMockRecorder) SetTxCtx(ctx, key, seconds, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTxCtx", reflect.TypeOf((*MockRedisRepository)(nil).SetTxCtx), ctx, key, seconds, user)
}