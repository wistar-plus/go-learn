package service

import (
	"go-learn/internal/app/wweb/model"
	"go-learn/internal/app/wweb/repository"
	"go-learn/internal/pkg/e"
)

// UserService 用户业务代码
type UserService struct {
	repo repository.IUserRepository
}

// NewUserService create a UserService obj
func NewUserService(repo repository.IUserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

//Exist Determine if the user exists
func (us *UserService) Exist(email string) bool {
	user, _ := us.repo.FindByEmail(email)
	return user != nil
}

//Create user
func (us *UserService) Create(user *model.User) error {
	if us.Exist(user.Email) {
		return e.ERROR_EXIST
	}

	return us.repo.Create(user)
}

//Update user
func (us *UserService) Update(user *model.User) error {
	if !us.Exist(user.Email) {
		return e.ERROR_NOTEXIST
	}

	return us.repo.Update(user)
}

// FindByPage 分页查询user
func (us *UserService) FindByPage(page, size int) ([]*model.User, error) {
	return us.repo.FindByPage(page, size)
}

// Delete user
func (us *UserService) Delete(id uint32) error {
	return us.repo.Delete(id)
}

func (us *UserService) Login(userLogin *model.UserLoginParam) (*model.User, error) {
	user, err := us.repo.FindByEmail(userLogin.Email)
	if err == nil && user.Pwd != userLogin.Pwd {
		return nil, e.ERROR_WRONGPASSWORD
	}
	return user, err
}
