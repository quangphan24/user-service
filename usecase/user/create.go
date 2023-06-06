package user

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"user-service/model"
	"user-service/util"
)

func (uc *UserUseCase) Create(user model.User) (*model.User, error) {
	var (
		oldUser *model.User
		err     error
	)
	if oldUser, err = uc.repo.GetOneUserByEmail(user.Email); err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, err
		}
	}
	if oldUser != nil {
		return nil, errors.New(user.Email + ` is not available`)
	}

	//hash password
	user.Password, err = util.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	//create
	if err := uc.repo.CreateUser(&user); err != nil {
		return nil, err
	}
	return &user, nil
}
