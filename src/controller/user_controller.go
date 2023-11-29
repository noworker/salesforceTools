package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/noworker/salesforceTools/domain/model"
	"github.com/noworker/salesforceTools/domain/usecase"
)

type IUserController interface {
	SignUp(c echo.Context) error
}

type UserController struct {
	uu usecase.IUserUsecase
}

func NewUserController(ur usecase.IUserUsecase) IUserController {
	return &UserController{ur}
}

func (uc *UserController) SignUp(c echo.Context) error {
	userModel := &model.User{}
	err := c.Bind(userModel)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	userRes, err := uc.uu.SignUp(userModel)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, userRes)
}
