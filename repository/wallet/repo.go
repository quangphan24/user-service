package wallet

import (
	"gorm.io/gorm"
	"user-service/model"
)

type RepoWallet struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) IRepoWallet {
	return &RepoWallet{db: db}
}

//go:generate mockery --name IRepoWallet
type IRepoWallet interface {
	GetOneWallet(id string) (*model.Wallets, error)
	UpdateBalance(id string, amount int) error
	Create(wallet *model.Wallets) error
}

func (r *RepoWallet) GetOneWallet(id string) (*model.Wallets, error) {
	rt := &model.Wallets{}

	err := r.db.Table("wallets").Where("id = ?", id).Take(&rt).Error
	return rt, err
}
func (r *RepoWallet) UpdateBalance(id string, amount int) error {
	return r.db.Where("id = ?", id).Table("wallets").Updates(map[string]interface{}{"balance": amount}).Error
}
func (r *RepoWallet) Create(wallet *model.Wallets) error {
	return r.db.Table("wallets").Create(&wallet).Error
}
