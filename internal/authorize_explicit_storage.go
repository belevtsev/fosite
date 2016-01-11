// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/ory-am/fosite/handler/authorize (interfaces: AuthorizeExplicitStorage)

package internal

import (
	gomock "github.com/golang/mock/gomock"
	fosite "github.com/ory-am/fosite"
	authorize "github.com/ory-am/fosite/handler/authorize"
	token "github.com/ory-am/fosite/handler/token"
)

// Mock of AuthorizeExplicitStorage interface
type MockAuthorizeExplicitStorage struct {
	ctrl     *gomock.Controller
	recorder *_MockAuthorizeExplicitStorageRecorder
}

// Recorder for MockAuthorizeExplicitStorage (not exported)
type _MockAuthorizeExplicitStorageRecorder struct {
	mock *MockAuthorizeExplicitStorage
}

func NewMockAuthorizeExplicitStorage(ctrl *gomock.Controller) *MockAuthorizeExplicitStorage {
	mock := &MockAuthorizeExplicitStorage{ctrl: ctrl}
	mock.recorder = &_MockAuthorizeExplicitStorageRecorder{mock}
	return mock
}

func (_m *MockAuthorizeExplicitStorage) EXPECT() *_MockAuthorizeExplicitStorageRecorder {
	return _m.recorder
}

func (_m *MockAuthorizeExplicitStorage) CreateAccessTokenSession(_param0 string, _param1 fosite.AccessRequester, _param2 *token.TokenSession) error {
	ret := _m.ctrl.Call(_m, "CreateAccessTokenSession", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockAuthorizeExplicitStorageRecorder) CreateAccessTokenSession(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateAccessTokenSession", arg0, arg1, arg2)
}

func (_m *MockAuthorizeExplicitStorage) CreateAuthorizeCodeSession(_param0 string, _param1 fosite.AuthorizeRequester, _param2 *authorize.AuthorizeSession) error {
	ret := _m.ctrl.Call(_m, "CreateAuthorizeCodeSession", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockAuthorizeExplicitStorageRecorder) CreateAuthorizeCodeSession(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateAuthorizeCodeSession", arg0, arg1, arg2)
}

func (_m *MockAuthorizeExplicitStorage) CreateRefreshTokenSession(_param0 string, _param1 fosite.AccessRequester, _param2 *token.TokenSession) error {
	ret := _m.ctrl.Call(_m, "CreateRefreshTokenSession", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockAuthorizeExplicitStorageRecorder) CreateRefreshTokenSession(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateRefreshTokenSession", arg0, arg1, arg2)
}

func (_m *MockAuthorizeExplicitStorage) DeleteAccessTokenSession(_param0 string) error {
	ret := _m.ctrl.Call(_m, "DeleteAccessTokenSession", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockAuthorizeExplicitStorageRecorder) DeleteAccessTokenSession(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteAccessTokenSession", arg0)
}

func (_m *MockAuthorizeExplicitStorage) DeleteAuthorizeCodeSession(_param0 string) error {
	ret := _m.ctrl.Call(_m, "DeleteAuthorizeCodeSession", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockAuthorizeExplicitStorageRecorder) DeleteAuthorizeCodeSession(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteAuthorizeCodeSession", arg0)
}

func (_m *MockAuthorizeExplicitStorage) DeleteRefreshTokenSession(_param0 string) error {
	ret := _m.ctrl.Call(_m, "DeleteRefreshTokenSession", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockAuthorizeExplicitStorageRecorder) DeleteRefreshTokenSession(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteRefreshTokenSession", arg0)
}

func (_m *MockAuthorizeExplicitStorage) GetAccessTokenSession(_param0 string, _param1 *token.TokenSession) (fosite.AccessRequester, error) {
	ret := _m.ctrl.Call(_m, "GetAccessTokenSession", _param0, _param1)
	ret0, _ := ret[0].(fosite.AccessRequester)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAuthorizeExplicitStorageRecorder) GetAccessTokenSession(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetAccessTokenSession", arg0, arg1)
}

func (_m *MockAuthorizeExplicitStorage) GetAuthorizeCodeSession(_param0 string, _param1 *authorize.AuthorizeSession) (fosite.AuthorizeRequester, error) {
	ret := _m.ctrl.Call(_m, "GetAuthorizeCodeSession", _param0, _param1)
	ret0, _ := ret[0].(fosite.AuthorizeRequester)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAuthorizeExplicitStorageRecorder) GetAuthorizeCodeSession(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetAuthorizeCodeSession", arg0, arg1)
}

func (_m *MockAuthorizeExplicitStorage) GetRefreshTokenSession(_param0 string, _param1 *token.TokenSession) (fosite.AccessRequester, error) {
	ret := _m.ctrl.Call(_m, "GetRefreshTokenSession", _param0, _param1)
	ret0, _ := ret[0].(fosite.AccessRequester)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAuthorizeExplicitStorageRecorder) GetRefreshTokenSession(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetRefreshTokenSession", arg0, arg1)
}
