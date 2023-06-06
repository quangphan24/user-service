package wallet

func (u *WalletUseCase) UpdateBalance(id string, amount int) error {
	return u.repo.UpdateBalance(id, amount)
}
