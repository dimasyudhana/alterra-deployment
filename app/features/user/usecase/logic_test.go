package usecase_test

import (
	"errors"
	"testing"

	user "github.com/dimasyudhana/latihan-deployment.git/app/features/user"
	mock "github.com/dimasyudhana/latihan-deployment.git/app/features/user/mocks"
	usecase "github.com/dimasyudhana/latihan-deployment.git/app/features/user/usecase"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestRegister_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock.NewMockRepository(ctrl)

	// Prepare the mock expectation.
	mockRepo.EXPECT().
		Register(gomock.Any()).
		Return(user.Core{}, nil)

	uc := usecase.New(mockRepo)
	err := uc.Register(user.Core{
		Phone:    "08123456789",
		Username: "Made",
		Password: "Password@20",
		Email:    "made@example.com",
	})

	assert.Nil(t, err)
}

func TestRegister_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock.NewMockRepository(ctrl)

	// Prepare the mock expectation.
	mockRepo.EXPECT().
		Register(gomock.Any()).
		Return(user.Core{}, errors.New("register failed"))

	uc := usecase.New(mockRepo)
	err := uc.Register(user.Core{
		Phone:    "08123456789",
		Username: "Made",
		Password: "Password@20",
		Email:    "made@example.com",
	})

	assert.NotNil(t, err)
	assert.EqualError(t, err, "terjadi kesalahan pada server")
}

func TestLogin_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock.NewMockRepository(ctrl)

	// Prepare the mock expectation.
	mockRepo.EXPECT().
		Login("08123456789", "Password@20").
		Return(user.Core{
			Phone:    "08123456789",
			Username: "Made",
			Email:    "made@example.com",
		}, nil)

	uc := usecase.New(mockRepo)
	userData, err := uc.Login("08123456789", "Password@20")

	assert.Nil(t, err)
	assert.Equal(t, "Made", userData.Username)
}

func TestLogin_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock.NewMockRepository(ctrl)

	// Prepare the mock expectation.
	mockRepo.EXPECT().
		Login("08123456789", "Password@20").
		Return(user.Core{}, errors.New("login failed"))

	uc := usecase.New(mockRepo)
	userData, err := uc.Login("08123456789", "Password@20")

	assert.NotNil(t, err)
	assert.EqualError(t, err, "terdapat permasalahan pada server")
	assert.Equal(t, user.Core{}, userData)
}

func TestFindByPhone_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock.NewMockRepository(ctrl)

	mockUsers := []*user.Core{
		{
			Username: "Made",
			Phone:    "08123456789",
			Email:    "made@example.com",
			Password: "Password@20",
		},
	}

	mockRepo.EXPECT().
		FindByPhone("08123456789").
		Return(mockUsers, nil)

	uc := usecase.New(mockRepo)

	users, err := uc.FindByPhone("08123456789")

	assert.NoError(t, err)
	assert.Equal(t, mockUsers, users)
}

func TestFindByPhone_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock.NewMockRepository(ctrl)

	mockRepo.EXPECT().
		FindByPhone("08123456789").
		Return(nil, errors.New("failed to find user by phone"))

	uc := usecase.New(mockRepo)

	users, err := uc.FindByPhone("08123456789")

	assert.Nil(t, users)
	assert.Error(t, err)
}

func TestUpdateByPhone_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock.NewMockRepository(ctrl)

	mockUser := []*user.Core{{
		Username: "newusername",
		Phone:    "08123456789",
		Email:    "newemail",
		Password: "Password@20",
	}}

	mockRepo.EXPECT().FindByPhone(mockUser[0].Phone).Return(mockUser, nil)

	mockRepo.EXPECT().UpdateByPhone(mockUser[0].Phone, gomock.AssignableToTypeOf(user.Core{})).
		DoAndReturn(func(phone string, updatedUser user.Core) error {
			assert.Equal(t, mockUser[0].Username, updatedUser.Username)
			assert.Equal(t, mockUser[0].Email, updatedUser.Email)
			return nil
		})

	uc := usecase.New(mockRepo)

	_, err := uc.UpdateByPhone(mockUser[0].Phone, "newusername", "newemail")

	assert.NoError(t, err)
}

func TestUpdateByPhone_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock.NewMockRepository(ctrl)

	mockUser := []*user.Core{{
		Username: "Made",
		Phone:    "08123456789",
		Email:    "made@example.com",
		Password: "Password@20",
	}}

	mockRepo.EXPECT().FindByPhone(mockUser[0].Phone).Return(nil, errors.New("database error"))

	uc := usecase.New(mockRepo)

	_, err := uc.UpdateByPhone(mockUser[0].Phone, "newusername", "newemail")

	assert.Error(t, err)
}
