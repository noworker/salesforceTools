package test

import (
	"os"
	"testing"

	"github.com/noworker/salesforceTools/domain/model"
	"github.com/noworker/salesforceTools/infrastructure/salesforce"
	"github.com/noworker/salesforceTools/test"
)

func TestGetSalesforceTokenByUserAndPassword(t *testing.T) {
	test.LoadEnv("../../../.evn.dev")
	salesforceClient := salesforce.NewSalesforceClient()
	sat := model.SalesforceAccessTokenModel{}
	salesforceAccessTokenResponse, err := salesforceClient.GetSalesforceTokenByUserAndPassword(
		sat,
		os.Getenv("SALESFORCE_LOGIN_DOMAIN"),
		os.Getenv("SALESFORCE_CLIENT_ID"),
		os.Getenv("SALESFORCE_CLIENT_SECRET"),
		os.Getenv("SALESFORCE_USERNAME"),
		os.Getenv("SALESFORCE_PASSWORD"),
	)
	if err != nil {
		t.Log(err)
	}
	t.Log(salesforceAccessTokenResponse.AccessToken)
}

func TestGetSalesforceDescribeGlobal(t *testing.T) {
	test.LoadEnv("../../../.env.dev")
	salesforceClient := salesforce.NewSalesforceClient()
	model := model.SalesforceDescribeGlobalModel{}
	retModel, err := salesforceClient.GetSalesforceDescribeGlobal(model, os.Getenv("SALESFORCE_MY_DOMAIN"), os.Getenv("SALESFORCE_ACCESS_TOKEN"))
	if err != nil {
		t.Log(err)
	}
	t.Log(retModel.Encoding)
	for _, v := range retModel.Sobjects {
		t.Log(v.Name)
		t.Log(v.Label)
	}

}
