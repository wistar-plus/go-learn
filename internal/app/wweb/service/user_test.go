package service

import (
	"go-learn/internal/app/wweb/model"
	"go-learn/internal/app/wweb/repository"
	"go-learn/internal/pkg/e"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUser_Create_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := repository.NewMockIUserRepository(ctrl)

	repo.EXPECT().FindByEmail("wistar").Return(nil, nil)
	repo.EXPECT().Create(&model.User{Email: "wistar"}).Return(nil)

	s := NewUserService(repo)
	assert.Nil(t, s.Create(&model.User{Email: "wistar"}))
}

func TestUser_Create_UserExist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := repository.NewMockIUserRepository(ctrl)

	repo.EXPECT().FindByEmail("wistar").Return(&model.User{Email: "wistar"}, nil)
	s := NewUserService(repo)
	assert.Equal(t, e.ERROR_EXIST, s.Create(&model.User{Email: "wistar"}))
}

// ......
