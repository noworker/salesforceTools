package test

import (
	"encoding/json"
	"testing"

	"github.com/noworker/salesforceTools/db"
	"github.com/noworker/salesforceTools/domain/model"
	"github.com/noworker/salesforceTools/infrastructure/repositories"
	"github.com/noworker/salesforceTools/test"
)

func TestUpdateUser(t *testing.T) {
	test.LoadEnv("../../../.env.dev")
	dbInstance := db.NewDB()
	repository := repositories.NewUserRepository(dbInstance)
	user := &model.User{
		Id:   "adae4fd225d9432ea26e4e5831a86c63",
		Name: "updateTest",
	}
	err := repository.UpdateUser(user)
	if err != nil {
		t.Log(err)
	}
	json, _ := json.MarshalIndent(user, "", " ")
	jsonString := string(json)
	t.Log(jsonString)
}
