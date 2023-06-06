package wallet

import "user-service/model"

func (u *WalletUseCase) GetOneWallet(id string) (*model.Wallets, error) {
	return u.repo.GetOneWallet(id)
}
