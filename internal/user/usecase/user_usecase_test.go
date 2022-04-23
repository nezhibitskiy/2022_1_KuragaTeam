package usecase

import (
	"errors"
	"fmt"
	"io"
	"myapp/internal/models"
	"myapp/internal/utils/constants"
	"myapp/mock"
	"strconv"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUserUseCase_SignUp(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	mockStorage := mock.NewMockStorage(ctl)
	mockRedis := mock.NewMockRedisStore(ctl)
	mockImages := mock.NewMockImageStorage(ctl)

	tests := []struct {
		name            string
		mock            func()
		input           *models.CreateUserDTO
		expectedSession string
		expectedMsg     string
		expectedErr     error
	}{
		{
			name: "Successfully",
			mock: func() {
				userModel := &models.User{
					Name:     "Danya",
					Email:    "danya@mail.ru",
					Password: "danya123321",
				}
				gomock.InOrder(
					mockStorage.EXPECT().IsUserUnique(userModel).Return(true, nil),
					mockStorage.EXPECT().CreateUser(userModel).Return(int64(1), nil),
					mockRedis.EXPECT().StoreSession(int64(1)).Return("session", nil),
				)
			},
			input: &models.CreateUserDTO{
				Name:     "Danya",
				Email:    "danya@mail.ru",
				Password: "danya123321",
			},
			expectedSession: "session",
			expectedMsg:     "",
			expectedErr:     nil,
		},
		{
			name: "Wrong validation",
			mock: func() {
			},
			input: &models.CreateUserDTO{
				Name:     "Danya",
				Email:    "danya@mail.ru",
				Password: "danya",
			},
			expectedSession: "",
			expectedMsg:     "",
			expectedErr:     constants.ErrNum,
		},
		{
			name: "User is not unique",
			mock: func() {
				userModel := &models.User{
					Name:     "Danya",
					Email:    "danya@mail.ru",
					Password: "danya123321",
				}
				gomock.InOrder(
					mockStorage.EXPECT().IsUserUnique(userModel).Return(false, nil),
				)
			},
			input: &models.CreateUserDTO{
				Name:     "Danya",
				Email:    "danya@mail.ru",
				Password: "danya123321",
			},
			expectedSession: "",
			expectedMsg:     "ERROR: Email is not unique",
			expectedErr:     nil,
		},
		{
			name: "Error occurred in IsUserUnique",
			mock: func() {
				userModel := &models.User{
					Name:     "Danya",
					Email:    "danya@mail.ru",
					Password: "danya123321",
				}
				gomock.InOrder(
					mockStorage.EXPECT().IsUserUnique(userModel).Return(false, errors.New("error")),
				)
			},
			input: &models.CreateUserDTO{
				Name:     "Danya",
				Email:    "danya@mail.ru",
				Password: "danya123321",
			},
			expectedSession: "",
			expectedMsg:     "",
			expectedErr:     errors.New("error"),
		},
		{
			name: "Error occurred in CreateUser",
			mock: func() {
				userModel := &models.User{
					Name:     "Danya",
					Email:    "danya@mail.ru",
					Password: "danya123321",
				}
				gomock.InOrder(
					mockStorage.EXPECT().IsUserUnique(userModel).Return(true, nil),
					mockStorage.EXPECT().CreateUser(userModel).Return(int64(1), errors.New("error")),
				)
			},
			input: &models.CreateUserDTO{
				Name:     "Danya",
				Email:    "danya@mail.ru",
				Password: "danya123321",
			},
			expectedSession: "",
			expectedMsg:     "",
			expectedErr:     errors.New("error"),
		},
		{
			name: "Error occurred in StoreSession",
			mock: func() {
				userModel := &models.User{
					Name:     "Danya",
					Email:    "danya@mail.ru",
					Password: "danya123321",
				}
				gomock.InOrder(
					mockStorage.EXPECT().IsUserUnique(userModel).Return(true, nil),
					mockStorage.EXPECT().CreateUser(userModel).Return(int64(1), nil),
					mockRedis.EXPECT().StoreSession(int64(1)).Return("", errors.New("error")),
				)
			},
			input: &models.CreateUserDTO{
				Name:     "Danya",
				Email:    "danya@mail.ru",
				Password: "danya123321",
			},
			expectedSession: "",
			expectedMsg:     "",
			expectedErr:     errors.New("error"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			th := test
			th.mock()

			service := NewService(mockStorage, mockRedis, mockImages)

			session, msg, err := service.SignUp(th.input)

			if th.expectedErr != nil {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, th.expectedSession, session)
				assert.Equal(t, th.expectedMsg, msg)
			}
		})
	}
}

func TestUserUseCase_LogIn(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	mockStorage := mock.NewMockStorage(ctl)
	mockRedis := mock.NewMockRedisStore(ctl)
	mockImages := mock.NewMockImageStorage(ctl)

	tests := []struct {
		name            string
		mock            func()
		input           *models.LogInUserDTO
		expectedSession string
		expectedErr     error
	}{
		{
			name: "Successfully, user exists",
			mock: func() {
				userModel := &models.User{
					Email:    "danya@mail.ru",
					Password: "danya123321",
				}
				gomock.InOrder(
					mockStorage.EXPECT().IsUserExists(userModel).Return(int64(1), true, nil),
					mockRedis.EXPECT().StoreSession(int64(1)).Return("session", nil),
				)
			},
			input: &models.LogInUserDTO{
				Email:    "danya@mail.ru",
				Password: "danya123321",
			},
			expectedSession: "session",
			expectedErr:     nil,
		},
		{
			name: "Successfully, user not exists",
			mock: func() {
				userModel := &models.User{
					Email:    "danya@mail.ru",
					Password: "danya123321",
				}
				gomock.InOrder(
					mockStorage.EXPECT().IsUserExists(userModel).Return(int64(0), false, nil),
				)
			},
			input: &models.LogInUserDTO{
				Email:    "danya@mail.ru",
				Password: "danya123321",
			},
			expectedSession: "",
			expectedErr:     nil,
		},
		{
			name: "Error occurred in IsUserExists",
			mock: func() {
				userModel := &models.User{
					Email:    "danya@mail.ru",
					Password: "danya123321",
				}
				gomock.InOrder(
					mockStorage.EXPECT().IsUserExists(userModel).Return(int64(0), false, errors.New("eroor")),
				)
			},
			input: &models.LogInUserDTO{
				Email:    "danya@mail.ru",
				Password: "danya123321",
			},
			expectedSession: "",
			expectedErr:     errors.New("error"),
		},
		{
			name: "Error occurred in StoreSession",
			mock: func() {
				userModel := &models.User{
					Email:    "danya@mail.ru",
					Password: "danya123321",
				}
				gomock.InOrder(
					mockStorage.EXPECT().IsUserExists(userModel).Return(int64(1), true, nil),
					mockRedis.EXPECT().StoreSession(int64(1)).Return("", errors.New("error")),
				)
			},
			input: &models.LogInUserDTO{
				Email:    "danya@mail.ru",
				Password: "danya123321",
			},
			expectedSession: "",
			expectedErr:     errors.New("error"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			th := test
			th.mock()

			service := NewService(mockStorage, mockRedis, mockImages)

			session, err := service.LogIn(th.input)

			if th.expectedErr != nil {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, th.expectedSession, session)
			}
		})
	}
}

func TestUserUseCase_LogOut(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	mockStorage := mock.NewMockStorage(ctl)
	mockRedis := mock.NewMockRedisStore(ctl)
	mockImages := mock.NewMockImageStorage(ctl)

	tests := []struct {
		name        string
		mock        func()
		input       string
		expectedErr error
	}{
		{
			name: "Successfully",
			mock: func() {
				gomock.InOrder(
					mockRedis.EXPECT().DeleteSession("session").Return(nil),
				)
			},
			input:       "session",
			expectedErr: nil,
		},
		{
			name: "Error occurred in DeleteSession",
			mock: func() {
				gomock.InOrder(
					mockRedis.EXPECT().DeleteSession("session").Return(errors.New("error")),
				)
			},
			input:       "session",
			expectedErr: errors.New("error"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			th := test
			th.mock()

			service := NewService(mockStorage, mockRedis, mockImages)

			err := service.LogOut(th.input)

			if th.expectedErr != nil {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestUserUseCase_CheckAuthorization(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	mockStorage := mock.NewMockStorage(ctl)
	mockRedis := mock.NewMockRedisStore(ctl)
	mockImages := mock.NewMockImageStorage(ctl)

	tests := []struct {
		name        string
		mock        func()
		input       string
		expectedId  int64
		expectedErr error
	}{
		{
			name: "Successfully",
			mock: func() {
				gomock.InOrder(
					mockRedis.EXPECT().GetUserId("session").Return(int64(1), nil),
				)
			},
			input:       "session",
			expectedId:  int64(1),
			expectedErr: nil,
		},
		{
			name: "Error occurred in GetUserId",
			mock: func() {
				gomock.InOrder(
					mockRedis.EXPECT().GetUserId("session").Return(int64(-1), errors.New("error")),
				)
			},
			input:       "session",
			expectedId:  int64(-1),
			expectedErr: errors.New("error"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			th := test
			th.mock()

			service := NewService(mockStorage, mockRedis, mockImages)

			id, err := service.CheckAuthorization(th.input)

			if th.expectedErr != nil {
				assert.NotNil(t, err)
				assert.Equal(t, th.expectedId, id)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, th.expectedId, id)
			}
		})
	}
}

func TestUserUseCase_GetUserProfile(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	mockStorage := mock.NewMockStorage(ctl)
	mockRedis := mock.NewMockRedisStore(ctl)
	mockImages := mock.NewMockImageStorage(ctl)

	tests := []struct {
		name        string
		mock        func()
		input       int64
		expected    *models.ProfileUserDTO
		expectedErr error
	}{
		{
			name: "Successfully",
			mock: func() {
				userModel := &models.User{
					Name:   "Danya",
					Email:  "danya@mail.ru",
					Avatar: "avatar.webp",
				}
				gomock.InOrder(
					mockStorage.EXPECT().GetUserProfile(int64(1)).Return(userModel, nil),
				)
			},
			input: int64(1),
			expected: &models.ProfileUserDTO{
				Name:   "Danya",
				Email:  "danya@mail.ru",
				Avatar: "avatar.webp",
			},
			expectedErr: nil,
		},
		{
			name: "Error occurred in GetUserProfile",
			mock: func() {
				gomock.InOrder(
					mockStorage.EXPECT().GetUserProfile(int64(1)).Return(nil, errors.New("error")),
				)
			},
			input:       int64(1),
			expected:    nil,
			expectedErr: errors.New("error"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			th := test
			th.mock()

			service := NewService(mockStorage, mockRedis, mockImages)

			profile, err := service.GetUserProfile(th.input)

			if th.expectedErr != nil {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, *th.expected, *profile)
			}
		})
	}
}

func TestUserUseCase_EditProfile(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	mockStorage := mock.NewMockStorage(ctl)
	mockRedis := mock.NewMockRedisStore(ctl)
	mockImages := mock.NewMockImageStorage(ctl)

	tests := []struct {
		name        string
		mock        func()
		input       *models.EditProfileDTO
		expectedErr error
	}{
		{
			name: "Successfully",
			mock: func() {
				userModel := &models.User{
					ID:       int64(1),
					Name:     "Danya",
					Password: "danya123321",
				}
				gomock.InOrder(
					mockStorage.EXPECT().EditProfile(userModel).Return(nil),
				)
			},
			input: &models.EditProfileDTO{
				ID:       int64(1),
				Name:     "Danya",
				Password: "danya123321",
			},
			expectedErr: nil,
		},
		{
			name: "Error occurred in EditProfile",
			mock: func() {
				userModel := &models.User{
					ID:       int64(1),
					Name:     "Danya",
					Password: "danya123321",
				}
				gomock.InOrder(
					mockStorage.EXPECT().EditProfile(userModel).Return(errors.New("error")),
				)
			},
			input: &models.EditProfileDTO{
				ID:       int64(1),
				Name:     "Danya",
				Password: "danya123321",
			},
			expectedErr: errors.New("error"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			th := test
			th.mock()

			service := NewService(mockStorage, mockRedis, mockImages)

			err := service.EditProfile(th.input)

			if th.expectedErr != nil {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestUserUseCase_EditAvatar(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	mockStorage := mock.NewMockStorage(ctl)
	mockRedis := mock.NewMockRedisStore(ctl)
	mockImages := mock.NewMockImageStorage(ctl)

	tests := []struct {
		name        string
		mock        func()
		input       *models.EditAvatarDTO
		expectedErr error
	}{
		{
			name: "Successfully",
			mock: func() {
				userModel := &models.User{
					ID:     int64(1),
					Avatar: "avatar.webp",
				}
				gomock.InOrder(
					mockStorage.EXPECT().EditAvatar(userModel).Return("old_avatar", nil),
					mockImages.EXPECT().DeleteFile("old_avatar").Return(nil),
				)
			},
			input: &models.EditAvatarDTO{
				ID:     int64(1),
				Avatar: "avatar.webp",
			},
			expectedErr: nil,
		},
		{
			name: "Error occurred in EditAvatar",
			mock: func() {
				userModel := &models.User{
					ID:     int64(1),
					Avatar: "avatar.webp",
				}
				gomock.InOrder(
					mockStorage.EXPECT().EditAvatar(userModel).Return("", errors.New("error")),
				)
			},
			input: &models.EditAvatarDTO{
				ID:     int64(1),
				Avatar: "avatar.webp",
			},
			expectedErr: errors.New("error"),
		},
		{
			name: "Error occurred in DeleteFile",
			mock: func() {
				userModel := &models.User{
					ID:     int64(1),
					Avatar: "avatar.webp",
				}
				gomock.InOrder(
					mockStorage.EXPECT().EditAvatar(userModel).Return("old_avatar", nil),
					mockImages.EXPECT().DeleteFile("old_avatar").Return(errors.New("error")),
				)
			},
			input: &models.EditAvatarDTO{
				ID:     int64(1),
				Avatar: "avatar.webp",
			},
			expectedErr: errors.New("error"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			th := test
			th.mock()

			service := NewService(mockStorage, mockRedis, mockImages)

			err := service.EditAvatar(th.input)

			if th.expectedErr != nil {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestUserUseCase_UploadAvatar(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	mockStorage := mock.NewMockStorage(ctl)
	mockRedis := mock.NewMockRedisStore(ctl)
	mockImages := mock.NewMockImageStorage(ctl)

	tests := []struct {
		name             string
		mock             func()
		inputFile        io.Reader
		inputSize        int64
		inputContentType string
		inputUserID      int64
		expectedName     string
		expectedErr      error
	}{
		{
			name: "Successfully",
			mock: func() {
				userModel := models.UploadInput{
					UserID:      int64(1),
					File:        *new(io.Reader),
					Size:        int64(5),
					ContentType: "type",
				}
				gomock.InOrder(
					mockImages.EXPECT().UploadFile(userModel).Return(
						fmt.Sprintf("%s_%s.%s", strconv.Itoa(int(int64(1))),
							fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
								time.Now().Year(), time.Now().Month(), time.Now().Day(),
								time.Now().Hour(), time.Now().Minute(), time.Now().Second()), "webp"), nil),
				)
			},
			inputFile:        *new(io.Reader),
			inputSize:        int64(5),
			inputContentType: "type",
			inputUserID:      int64(1),
			expectedName: fmt.Sprintf("%s_%s.%s", strconv.Itoa(int(int64(1))),
				fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
					time.Now().Year(), time.Now().Month(), time.Now().Day(),
					time.Now().Hour(), time.Now().Minute(), time.Now().Second()), "webp"),
			expectedErr: nil,
		},
		{
			name: "Error occurred in UploadFile",
			mock: func() {
				userModel := models.UploadInput{
					UserID:      int64(1),
					File:        *new(io.Reader),
					Size:        int64(5),
					ContentType: "type",
				}
				gomock.InOrder(
					mockImages.EXPECT().UploadFile(userModel).Return("", errors.New("error")),
				)
			},
			inputFile:        *new(io.Reader),
			inputSize:        int64(5),
			inputContentType: "type",
			inputUserID:      int64(1),
			expectedName:     "",
			expectedErr:      errors.New("error"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			th := test
			th.mock()

			service := NewService(mockStorage, mockRedis, mockImages)

			path, err := service.UploadAvatar(th.inputFile, th.inputSize, th.inputContentType, th.inputUserID)

			if th.expectedErr != nil {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, th.expectedName, path)
			}
		})
	}
}

func TestUserUseCase_GetAvatar(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	mockStorage := mock.NewMockStorage(ctl)
	mockRedis := mock.NewMockRedisStore(ctl)
	mockImages := mock.NewMockImageStorage(ctl)

	tests := []struct {
		name        string
		mock        func()
		input       int64
		expected    string
		expectedErr error
	}{
		{
			name: "Successfully",
			mock: func() {
				gomock.InOrder(
					mockStorage.EXPECT().GetAvatar(int64(1)).Return(constants.DefaultImage, nil),
				)
			},
			input:       int64(1),
			expected:    constants.DefaultImage,
			expectedErr: nil,
		},
		{
			name: "Error occurred in GetAvatar",
			mock: func() {
				gomock.InOrder(
					mockStorage.EXPECT().GetAvatar(int64(1)).Return("", errors.New("error")),
				)
			},
			input:       int64(1),
			expectedErr: errors.New("error"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			th := test
			th.mock()

			service := NewService(mockStorage, mockRedis, mockImages)

			avatar, err := service.GetAvatar(th.input)

			if th.expectedErr != nil {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, th.expected, avatar)
			}
		})
	}
}
