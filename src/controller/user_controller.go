package controller

import (
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/noworker/salesforceTools/domain/model"
	"github.com/noworker/salesforceTools/domain/usecase"
)

type IUserController interface {
	SignUp(c echo.Context) error
	Login(c echo.Context) error
	Logout(c echo.Context) error
	UpdateUserSalesforceInfo(c echo.Context) error
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

func (uc *UserController) Login(c echo.Context) error {
	userModel := &model.User{}
	err := c.Bind(userModel)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	tokenString, err := uc.uu.Login(userModel)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	cookie := &http.Cookie{}
	cookie.Name = "token"
	cookie.Value = tokenString
	cookie.Expires = time.Now().Add(time.Hour * 24)
	cookie.Path = "/"
	cookie.HttpOnly = true
	cookie.Domain = os.Getenv("API_DOMAIN")
	if os.Getenv("MODE") == "DEV" {
		cookie.SameSite = http.SameSiteLaxMode
	} else {
		cookie.SameSite = http.SameSiteNoneMode
		cookie.Secure = true
	}
	c.SetCookie(cookie)
	return c.NoContent(http.StatusOK)
}

func (uc *UserController) Logout(c echo.Context) error {
	cookie := &http.Cookie{}
	cookie.Name = "token"
	cookie.Value = ""
	cookie.Expires = time.Now()
	cookie.Path = "/"
	cookie.HttpOnly = true
	cookie.Domain = os.Getenv("API_DOMAIN")
	if os.Getenv("MODE") == "DEV" {
		cookie.SameSite = http.SameSiteLaxMode
	} else {
		cookie.SameSite = http.SameSiteNoneMode
		cookie.Secure = true
	}
	c.SetCookie(cookie)
	return c.NoContent(http.StatusOK)
}

func (uc *UserController) UpdateUserSalesforceInfo(c echo.Context) error {
	receivedUser := &model.User{}
	err := c.Bind(receivedUser)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	updatedUser, err := uc.uu.UpdateUserSalesforceInfo(receivedUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, updatedUser)
}
