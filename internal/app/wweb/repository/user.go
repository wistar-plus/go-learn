package repository

import (
	"go-learn/internal/app/wweb/model"
)

// IUserRepository UserRepository interface
type IUserRepository interface {
	FindByPage(page, size int) ([]*model.User, error)
	FindByEmail(email string) (*model.User, error)
	Create(*model.User) error
	Update(*model.User) error
	Delete(id uint32) error
}

//mockgen -destination user_mock.go -source user.go -package repository
