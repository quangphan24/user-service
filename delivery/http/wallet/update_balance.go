package wallet

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"user-service/payload"
)

func (r *Route) UpdateBalance(c echo.Context) error {
	var req payload.UpdateBalanceReq
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	err := r.useCase.WalletUseCase.UpdateBalance(req.Id, req.Amount)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, "Successfully!")
}
