// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/microservices/profile/repository.go

// Package repository is a generated GoMock package.
package repository

import (
	proto "myapp/internal/microservices/profile/proto"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockStorage is a mock of Storage interface.
type MockStorage struct {
	ctrl     *gomock.Controller
	recorder *MockStorageMockRecorder
}

// MockStorageMockRecorder is the mock recorder for MockStorage.
type MockStorageMockRecorder struct {
	mock *MockStorage
}

// NewMockStorage creates a new mock instance.
func NewMockStorage(ctrl *gomock.Controller) *MockStorage {
	mock := &MockStorage{ctrl: ctrl}
	mock.recorder = &MockStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorage) EXPECT() *MockStorageMockRecorder {
	return m.recorder
}

// AddLike mocks base method.
func (m *MockStorage) AddLike(data *proto.LikeData) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddLike", data)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddLike indicates an expected call of AddLike.
func (mr *MockStorageMockRecorder) AddLike(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddLike", reflect.TypeOf((*MockStorage)(nil).AddLike), data)
}

// DeleteFile mocks base method.
func (m *MockStorage) DeleteFile(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFile", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteFile indicates an expected call of DeleteFile.
func (mr *MockStorageMockRecorder) DeleteFile(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFile", reflect.TypeOf((*MockStorage)(nil).DeleteFile), arg0)
}

// EditAvatar mocks base method.
func (m *MockStorage) EditAvatar(data *proto.EditAvatarData) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EditAvatar", data)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EditAvatar indicates an expected call of EditAvatar.
func (mr *MockStorageMockRecorder) EditAvatar(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditAvatar", reflect.TypeOf((*MockStorage)(nil).EditAvatar), data)
}

// EditProfile mocks base method.
func (m *MockStorage) EditProfile(data *proto.EditProfileData) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EditProfile", data)
	ret0, _ := ret[0].(error)
	return ret0
}

// EditProfile indicates an expected call of EditProfile.
func (mr *MockStorageMockRecorder) EditProfile(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditProfile", reflect.TypeOf((*MockStorage)(nil).EditProfile), data)
}

// GetAvatar mocks base method.
func (m *MockStorage) GetAvatar(userID int64) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAvatar", userID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAvatar indicates an expected call of GetAvatar.
func (mr *MockStorageMockRecorder) GetAvatar(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAvatar", reflect.TypeOf((*MockStorage)(nil).GetAvatar), userID)
}

// GetFavorites mocks base method.
func (m *MockStorage) GetFavorites(userID int64) (*proto.Favorites, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFavorites", userID)
	ret0, _ := ret[0].(*proto.Favorites)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFavorites indicates an expected call of GetFavorites.
func (mr *MockStorageMockRecorder) GetFavorites(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFavorites", reflect.TypeOf((*MockStorage)(nil).GetFavorites), userID)
}

// GetRating mocks base method.
func (m *MockStorage) GetRating(data *proto.MovieRating) (*proto.Rating, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRating", data)
	ret0, _ := ret[0].(*proto.Rating)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRating indicates an expected call of GetRating.
func (mr *MockStorageMockRecorder) GetRating(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRating", reflect.TypeOf((*MockStorage)(nil).GetRating), data)
}

// GetUserProfile mocks base method.
func (m *MockStorage) GetUserProfile(userID int64) (*proto.ProfileData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserProfile", userID)
	ret0, _ := ret[0].(*proto.ProfileData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserProfile indicates an expected call of GetUserProfile.
func (mr *MockStorageMockRecorder) GetUserProfile(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserProfile", reflect.TypeOf((*MockStorage)(nil).GetUserProfile), userID)
}

// RemoveLike mocks base method.
func (m *MockStorage) RemoveLike(data *proto.LikeData) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveLike", data)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveLike indicates an expected call of RemoveLike.
func (mr *MockStorageMockRecorder) RemoveLike(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveLike", reflect.TypeOf((*MockStorage)(nil).RemoveLike), data)
}

// UploadAvatar mocks base method.
func (m *MockStorage) UploadAvatar(data *proto.UploadInputFile) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadAvatar", data)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UploadAvatar indicates an expected call of UploadAvatar.
func (mr *MockStorageMockRecorder) UploadAvatar(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadAvatar", reflect.TypeOf((*MockStorage)(nil).UploadAvatar), data)
}
