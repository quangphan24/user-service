package wallet

import (
	"user-service/model"
	"user-service/repository"
	"user-service/repository/wallet"
)

type WalletUseCase struct {
	repo wallet.IRepoWallet
}
type IWalletUseCase interface {
	GetOneWallet(id string) (*model.Wallets, error)
	UpdateBalance(id string, amount int) error
	Payment(id string, amount int) error
	Create(wallet *model.Wallets) error
}

func New(repo *repository.Repository) IWalletUseCase {
	return &WalletUseCase{
		repo: repo.RepoWallet,
	}
}
