package wallet

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
	"user-service/model"
)

func (r *Route) Create(c echo.Context) error {
	wallet := &model.Wallets{}

	if err := c.Bind(&wallet); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	validate := validator.New()
	if err := validate.Struct(wallet); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err := r.useCase.WalletUseCase.Create(wallet)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, wallet)
}
