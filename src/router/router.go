package router

import (
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
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
	auth.POST("/login", uc.Login)
	auth.POST("/logout", uc.Logout)

	// api process
	api := e.Group("/api/v1")
	api.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	api.GET("/debuglogs", sc.GetDebugLogs)
	api.PATCH("/update-user-salesforce-info", uc.UpdateUserSalesforceInfo)
	return e
}
