package user

import "user-service/model"

func (uc *UserUseCase) GetOneUser(id string) (*model.User, error) {
	return uc.repo.GetOne(id)
}
