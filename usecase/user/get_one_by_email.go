package user

import "user-service/model"

func (uc *UserUseCase) GetOneByEmail(email string) (*model.User, error) {
	return uc.repo.GetOneUserByEmail(email)
}
