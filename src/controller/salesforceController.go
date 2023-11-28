package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/noworker/salesforceTools/domain/usecase"
)

type ISalesforceController interface {
	GetDebugLogs(c echo.Context) error
}

type SalesforceController struct {
	su usecase.ISalesforceUsecase
}

func NewSalesforceController(su usecase.ISalesforceUsecase) ISalesforceController {
	return &SalesforceController{su}
}

func (sc *SalesforceController) GetDebugLogs(c echo.Context) error {
	return c.JSON(200, nil)
}
