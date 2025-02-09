// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/openshift/assisted-service/restapi (interfaces: ManifestsAPI)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	middleware "github.com/go-openapi/runtime/middleware"
	gomock "github.com/golang/mock/gomock"
	manifests "github.com/openshift/assisted-service/restapi/operations/manifests"
)

// MockManifestsAPI is a mock of ManifestsAPI interface.
type MockManifestsAPI struct {
	ctrl     *gomock.Controller
	recorder *MockManifestsAPIMockRecorder
}

// MockManifestsAPIMockRecorder is the mock recorder for MockManifestsAPI.
type MockManifestsAPIMockRecorder struct {
	mock *MockManifestsAPI
}

// NewMockManifestsAPI creates a new mock instance.
func NewMockManifestsAPI(ctrl *gomock.Controller) *MockManifestsAPI {
	mock := &MockManifestsAPI{ctrl: ctrl}
	mock.recorder = &MockManifestsAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockManifestsAPI) EXPECT() *MockManifestsAPIMockRecorder {
	return m.recorder
}

// CreateClusterManifest mocks base method.
func (m *MockManifestsAPI) CreateClusterManifest(arg0 context.Context, arg1 manifests.CreateClusterManifestParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateClusterManifest", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// CreateClusterManifest indicates an expected call of CreateClusterManifest.
func (mr *MockManifestsAPIMockRecorder) CreateClusterManifest(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateClusterManifest", reflect.TypeOf((*MockManifestsAPI)(nil).CreateClusterManifest), arg0, arg1)
}

// DeleteClusterManifest mocks base method.
func (m *MockManifestsAPI) DeleteClusterManifest(arg0 context.Context, arg1 manifests.DeleteClusterManifestParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteClusterManifest", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// DeleteClusterManifest indicates an expected call of DeleteClusterManifest.
func (mr *MockManifestsAPIMockRecorder) DeleteClusterManifest(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteClusterManifest", reflect.TypeOf((*MockManifestsAPI)(nil).DeleteClusterManifest), arg0, arg1)
}

// DownloadClusterManifest mocks base method.
func (m *MockManifestsAPI) DownloadClusterManifest(arg0 context.Context, arg1 manifests.DownloadClusterManifestParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DownloadClusterManifest", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// DownloadClusterManifest indicates an expected call of DownloadClusterManifest.
func (mr *MockManifestsAPIMockRecorder) DownloadClusterManifest(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DownloadClusterManifest", reflect.TypeOf((*MockManifestsAPI)(nil).DownloadClusterManifest), arg0, arg1)
}

// ListClusterManifests mocks base method.
func (m *MockManifestsAPI) ListClusterManifests(arg0 context.Context, arg1 manifests.ListClusterManifestsParams) middleware.Responder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListClusterManifests", arg0, arg1)
	ret0, _ := ret[0].(middleware.Responder)
	return ret0
}

// ListClusterManifests indicates an expected call of ListClusterManifests.
func (mr *MockManifestsAPIMockRecorder) ListClusterManifests(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListClusterManifests", reflect.TypeOf((*MockManifestsAPI)(nil).ListClusterManifests), arg0, arg1)
}
