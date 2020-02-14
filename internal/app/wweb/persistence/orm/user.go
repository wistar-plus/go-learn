package orm

import (
	"go-learn/internal/app/wweb/model"
	"go-learn/internal/app/wweb/repository"
)

type userRepository struct{}

//NewUserRepository return userRepository obj
func NewUserRepository() repository.IUserRepository {
	return &userRepository{}
}

func (ur *userRepository) FindByPage(page, size int) ([]*model.User, error) {
	var res []*model.User
	q := db.Offset(page * size).Limit(size).Find(&res)
	if q.Error != nil && q.RecordNotFound() {
		return nil, nil
	}
	return res, q.Error
}

func (ur *userRepository) FindByEmail(email string) (*model.User, error) {
	res := new(model.User)
	q := db.Where("email = ?", email).First(res)
	if q.Error != nil {
		if q.RecordNotFound() {
			return nil, q.Error
		}
	}

	return res, q.Error
}

func (ur *userRepository) Create(user *model.User) error {
	return db.Create(user).Error
}

func (ur *userRepository) Update(user *model.User) error {
	return db.Model(&user).Update(map[string]interface{}{
		"nick_name": user.NickName,
		"pwd":       user.Pwd,
	}).Error
}

func (ur *userRepository) Delete(id uint32) error {
	return db.Delete(&model.User{Uid: id}).Error
}
