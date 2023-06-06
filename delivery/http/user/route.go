package user

import (
	"github.com/labstack/echo/v4"
	"user-service/usecase"
)

type Route struct {
	useCase *usecase.UseCase
}

func Init(group *echo.Group, useCase *usecase.UseCase) {
	r := &Route{useCase: useCase}

	group.POST("/create", r.Create)
	group.GET("/:id", r.GetOneUser)
	group.POST("/login", r.Login)
	group.GET("/refresh", r.RefreshToken)
}
