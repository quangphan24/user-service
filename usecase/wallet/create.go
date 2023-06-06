package wallet

import "user-service/model"

func (u *WalletUseCase) Create(wallet *model.Wallets) error {
	return u.repo.Create(wallet)
}
