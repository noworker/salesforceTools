package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/noworker/salesforceTools/controller"
	"github.com/noworker/salesforceTools/db"
	"github.com/noworker/salesforceTools/domain/usecase"
	"github.com/noworker/salesforceTools/infrastructure/repositories"
	"github.com/noworker/salesforceTools/infrastructure/salesforce"
	"github.com/noworker/salesforceTools/router"
)

func main() {
	// Load env
	if os.Getenv("MODE") == "DEV" {
		fmt.Println("Load env")
		err := godotenv.Load(".env.dev")
		if err != nil {
			log.Println(err)
		}
	}

	// db
	db := db.NewDB()
	// infrastructure
	salesforceClient := salesforce.NewSalesforceClient()
	userRepository := repositories.NewUserRepository(db)
	// usecase
	salesforceUsecase := usecase.NewSalesforceUsecase(salesforceClient)
	userUsecase := usecase.NewUserUsecase(userRepository)
	// controller
	salesforceController := controller.NewSalesforceController(salesforceUsecase)
	userController := controller.NewUserController(userUsecase)
	// router
	echo := router.NewRouter(
		userController,
		salesforceController,
	)
	// execute router
	echo.Logger.Fatal(echo.Start(":8080"))
}
