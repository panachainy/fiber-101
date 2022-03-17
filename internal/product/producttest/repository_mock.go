
// Code generated by MockGen. DO NOT EDIT.
// Source: internal/product/repository.go

// Package mock_product is a generated GoMock package.
package producttest

import (
        product "fiber-101/internal/product"
        reflect "reflect"

        gomock "github.com/golang/mock/gomock"
)

// MockProductRepository is a mock of ProductRepository interface.
type MockProductRepository struct {
        ctrl     *gomock.Controller
        recorder *MockProductRepositoryMockRecorder
}

// MockProductRepositoryMockRecorder is the mock recorder for MockProductRepository.
type MockProductRepositoryMockRecorder struct {
        mock *MockProductRepository
}

// NewMockProductRepository creates a new mock instance.
func NewMockProductRepository(ctrl *gomock.Controller) *MockProductRepository {
        mock := &MockProductRepository{ctrl: ctrl}
        mock.recorder = &MockProductRepositoryMockRecorder{mock}
        return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductRepository) EXPECT() *MockProductRepositoryMockRecorder {
        return m.recorder
}

// AddNewRegister mocks base method.
func (m *MockProductRepository) AddNewRegister(u product.Product) error {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "AddNewRegister", u)
        ret0, _ := ret[0].(error)
        return ret0
}

// AddNewRegister indicates an expected call of AddNewRegister.
func (mr *MockProductRepositoryMockRecorder) AddNewRegister(u interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddNewRegister", reflect.TypeOf((*MockProductRepository)(nil).AddNewRegister), u)
}

// GetAll mocks base method.
func (m *MockProductRepository) GetAll() []product.Product {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetAll")
        ret0, _ := ret[0].([]product.Product)
        return ret0
}

// GetAll indicates an expected call of GetAll.
func (mr *MockProductRepositoryMockRecorder) GetAll() *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockProductRepository)(nil).GetAll))
}

// RemoveById mocks base method.
func (m *MockProductRepository) RemoveById(uid string) error {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "RemoveById", uid)
        ret0, _ := ret[0].(error)
        return ret0
}

// RemoveById indicates an expected call of RemoveById.
func (mr *MockProductRepositoryMockRecorder) RemoveById(uid interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveById", reflect.TypeOf((*MockProductRepository)(nil).RemoveById), uid)
}