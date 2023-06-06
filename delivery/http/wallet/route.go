package wallet

import (
	"github.com/labstack/echo/v4"
	"user-service/usecase"
)

type Route struct {
	useCase *usecase.UseCase
}

func Init(group *echo.Group, useCase *usecase.UseCase) {
	r := &Route{useCase: useCase}

	group.GET("/:id", r.GetOneWallet)
	group.PUT("/:id", r.UpdateBalance)
	group.POST("", r.Create)
}
