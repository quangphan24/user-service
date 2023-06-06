package repository

import (
	"gorm.io/gorm"
	"user-service/repository/user"
	"user-service/repository/wallet"
)

type Repository struct {
	RepoUser   user.IRepoUser
	RepoWallet wallet.IRepoWallet
}

func New(db *gorm.DB) *Repository {
	return &Repository{
		RepoUser:   user.NewRepo(db),
		RepoWallet: wallet.NewRepo(db),
	}
}
