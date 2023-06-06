package user

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (r *Route) GetOneUser(c echo.Context) error {

	newUser, err := r.useCase.UserUseCase.GetOneUser(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, newUser)
}
