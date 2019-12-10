// Code generated by MockGen. DO NOT EDIT.
// Source: worker/appm/store/store.go

// Package store is a generated GoMock package.
package store

import (
	gomock "github.com/golang/mock/gomock"
	model "github.com/goodrain/rainbond/db/model"
	v1 "github.com/goodrain/rainbond/worker/appm/types/v1"
	v10 "k8s.io/api/core/v1"
	v11 "k8s.io/client-go/listers/core/v1"
	reflect "reflect"
)

// MockStorer is a mock of Storer interface
type MockStorer struct {
	ctrl     *gomock.Controller
	recorder *MockStorerMockRecorder
}

// MockStorerMockRecorder is the mock recorder for MockStorer
type MockStorerMockRecorder struct {
	mock *MockStorer
}

// NewMockStorer creates a new mock instance
func NewMockStorer(ctrl *gomock.Controller) *MockStorer {
	mock := &MockStorer{ctrl: ctrl}
	mock.recorder = &MockStorerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStorer) EXPECT() *MockStorerMockRecorder {
	return m.recorder
}

// Start mocks base method
func (m *MockStorer) Start() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start")
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start
func (mr *MockStorerMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockStorer)(nil).Start))
}

// Ready mocks base method
func (m *MockStorer) Ready() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ready")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Ready indicates an expected call of Ready
func (mr *MockStorerMockRecorder) Ready() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ready", reflect.TypeOf((*MockStorer)(nil).Ready))
}

// RegistAppService mocks base method
func (m *MockStorer) RegistAppService(arg0 *v1.AppService) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegistAppService", arg0)
}

// RegistAppService indicates an expected call of RegistAppService
func (mr *MockStorerMockRecorder) RegistAppService(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegistAppService", reflect.TypeOf((*MockStorer)(nil).RegistAppService), arg0)
}

// GetAppService mocks base method
func (m *MockStorer) GetAppService(serviceID string) *v1.AppService {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAppService", serviceID)
	ret0, _ := ret[0].(*v1.AppService)
	return ret0
}

// GetAppService indicates an expected call of GetAppService
func (mr *MockStorerMockRecorder) GetAppService(serviceID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAppService", reflect.TypeOf((*MockStorer)(nil).GetAppService), serviceID)
}

// UpdateGetAppService mocks base method
func (m *MockStorer) UpdateGetAppService(serviceID string) *v1.AppService {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateGetAppService", serviceID)
	ret0, _ := ret[0].(*v1.AppService)
	return ret0
}

// UpdateGetAppService indicates an expected call of UpdateGetAppService
func (mr *MockStorerMockRecorder) UpdateGetAppService(serviceID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateGetAppService", reflect.TypeOf((*MockStorer)(nil).UpdateGetAppService), serviceID)
}

// GetAllAppServices mocks base method
func (m *MockStorer) GetAllAppServices() []*v1.AppService {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllAppServices")
	ret0, _ := ret[0].([]*v1.AppService)
	return ret0
}

// GetAllAppServices indicates an expected call of GetAllAppServices
func (mr *MockStorerMockRecorder) GetAllAppServices() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllAppServices", reflect.TypeOf((*MockStorer)(nil).GetAllAppServices))
}

// GetAppServiceStatus mocks base method
func (m *MockStorer) GetAppServiceStatus(serviceID string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAppServiceStatus", serviceID)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetAppServiceStatus indicates an expected call of GetAppServiceStatus
func (mr *MockStorerMockRecorder) GetAppServiceStatus(serviceID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAppServiceStatus", reflect.TypeOf((*MockStorer)(nil).GetAppServiceStatus), serviceID)
}

// GetAppServicesStatus mocks base method
func (m *MockStorer) GetAppServicesStatus(serviceIDs []string) map[string]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAppServicesStatus", serviceIDs)
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// GetAppServicesStatus indicates an expected call of GetAppServicesStatus
func (mr *MockStorerMockRecorder) GetAppServicesStatus(serviceIDs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAppServicesStatus", reflect.TypeOf((*MockStorer)(nil).GetAppServicesStatus), serviceIDs)
}

// GetTenantResource mocks base method
func (m *MockStorer) GetTenantResource(tenantID string) *v1.TenantResource {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTenantResource", tenantID)
	ret0, _ := ret[0].(*v1.TenantResource)
	return ret0
}

// GetTenantResource indicates an expected call of GetTenantResource
func (mr *MockStorerMockRecorder) GetTenantResource(tenantID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTenantResource", reflect.TypeOf((*MockStorer)(nil).GetTenantResource), tenantID)
}

// GetTenantRunningApp mocks base method
func (m *MockStorer) GetTenantRunningApp(tenantID string) []*v1.AppService {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTenantRunningApp", tenantID)
	ret0, _ := ret[0].([]*v1.AppService)
	return ret0
}

// GetTenantRunningApp indicates an expected call of GetTenantRunningApp
func (mr *MockStorerMockRecorder) GetTenantRunningApp(tenantID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTenantRunningApp", reflect.TypeOf((*MockStorer)(nil).GetTenantRunningApp), tenantID)
}

// GetNeedBillingStatus mocks base method
func (m *MockStorer) GetNeedBillingStatus(serviceIDs []string) map[string]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNeedBillingStatus", serviceIDs)
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// GetNeedBillingStatus indicates an expected call of GetNeedBillingStatus
func (mr *MockStorerMockRecorder) GetNeedBillingStatus(serviceIDs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNeedBillingStatus", reflect.TypeOf((*MockStorer)(nil).GetNeedBillingStatus), serviceIDs)
}

// OnDelete mocks base method
func (m *MockStorer) OnDelete(obj interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnDelete", obj)
}

// OnDelete indicates an expected call of OnDelete
func (mr *MockStorerMockRecorder) OnDelete(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnDelete", reflect.TypeOf((*MockStorer)(nil).OnDelete), obj)
}

// GetPodLister mocks base method
func (m *MockStorer) GetPodLister() v11.PodLister {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPodLister")
	ret0, _ := ret[0].(v11.PodLister)
	return ret0
}

// GetPodLister indicates an expected call of GetPodLister
func (mr *MockStorerMockRecorder) GetPodLister() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPodLister", reflect.TypeOf((*MockStorer)(nil).GetPodLister))
}

// RegistPodUpdateListener mocks base method
func (m *MockStorer) RegistPodUpdateListener(arg0 string, arg1 chan<- *v10.Pod) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegistPodUpdateListener", arg0, arg1)
}

// RegistPodUpdateListener indicates an expected call of RegistPodUpdateListener
func (mr *MockStorerMockRecorder) RegistPodUpdateListener(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegistPodUpdateListener", reflect.TypeOf((*MockStorer)(nil).RegistPodUpdateListener), arg0, arg1)
}

// UnRegistPodUpdateListener mocks base method
func (m *MockStorer) UnRegistPodUpdateListener(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UnRegistPodUpdateListener", arg0)
}

// UnRegistPodUpdateListener indicates an expected call of UnRegistPodUpdateListener
func (mr *MockStorerMockRecorder) UnRegistPodUpdateListener(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnRegistPodUpdateListener", reflect.TypeOf((*MockStorer)(nil).UnRegistPodUpdateListener), arg0)
}

// InitOneThirdPartService mocks base method
func (m *MockStorer) InitOneThirdPartService(service *model.TenantServices) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InitOneThirdPartService", service)
	ret0, _ := ret[0].(error)
	return ret0
}

// InitOneThirdPartService indicates an expected call of InitOneThirdPartService
func (mr *MockStorerMockRecorder) InitOneThirdPartService(service interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitOneThirdPartService", reflect.TypeOf((*MockStorer)(nil).InitOneThirdPartService), service)
}
