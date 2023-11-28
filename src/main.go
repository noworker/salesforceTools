package main

import (
	"github.com/noworker/salesforceTools/controller"
	"github.com/noworker/salesforceTools/domain/usecase"
	"github.com/noworker/salesforceTools/infrastructure/router"
	"github.com/noworker/salesforceTools/infrastructure/salesforce"
)

func main() {
	// infrastructure
	salesforceClient := salesforce.NewSalesforceClient()
	// usecase
	salesforceUsecase := usecase.NewSalesforceUsecase(salesforceClient)
	// controller
	salesforceController := controller.NewSalesforceController(salesforceUsecase)
	// router
	echo := router.NewRouter(salesforceController)
	// execute router
	echo.Logger.Fatal(echo.Start(":8080"))
}
