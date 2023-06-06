package payload

type UpdateBalanceReq struct {
	Id     string `json:"wallet_id" param:"wallet_id" validate:"required"`
	Amount int    `json:"amount"  validate:"required"`
}
