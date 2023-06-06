package wallet

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (r *Route) GetOneWallet(c echo.Context) error {
	id := c.Param("id")

	wallet, err := r.useCase.UserUseCase.GetOneUser(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, wallet)
}
