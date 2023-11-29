package router

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/noworker/salesforceTools/controller"
)

func NewRouter(uc controller.IUserController, sc controller.ISalesforceController) *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"https://localhost:3000", os.Getenv("FS_URL")},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAccessControlAllowHeaders, echo.HeaderXCSRFToken},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
	}))
	// auth
	auth := e.Group("/auth")
	auth.POST("/signup", uc.SignUp)

	// api process
	api := e.Group("/api/v1")
	api.GET("/debuglogs", sc.GetDebugLogs)
	return e
}
