package wallet

func (u *WalletUseCase) Payment(walletId string, amount int) error {
	wallet, err := u.GetOneWallet(walletId)
	if err != nil {
		return err
	}
	return u.UpdateBalance(walletId, wallet.Balance-amount)
}
