package router

import (
	"github.com/labstack/echo/v4"
	"github.com/noworker/salesforceTools/controller"
)

func NewRouter(sc controller.ISalesforceController) *echo.Echo {
	e := echo.New()
	api := e.Group("/api/v1")
	api.GET("/debuglogs", sc.GetDebugLogs)
	return e
}
