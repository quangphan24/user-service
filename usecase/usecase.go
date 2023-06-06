package usecase

import (
	"user-service/repository"
	"user-service/usecase/user"
	"user-service/usecase/wallet"
)

type UseCase struct {
	UserUseCase   user.IUserUseCase
	WalletUseCase wallet.IWalletUseCase
}

func New(repo *repository.Repository) *UseCase {
	return &UseCase{
		UserUseCase:   user.New(repo),
		WalletUseCase: wallet.New(repo),
	}
}
