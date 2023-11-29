package main

import (
	"fmt"

	"github.com/noworker/salesforceTools/db"
	"github.com/noworker/salesforceTools/domain/model"
	"github.com/noworker/salesforceTools/test"
)

func main() {
	test.LoadEnv("../.env.dev")
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.User{})
}
