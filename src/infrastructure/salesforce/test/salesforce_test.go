package test

import (
	"os"
	"testing"

	"github.com/noworker/salesforceTools/domain/model"
	"github.com/noworker/salesforceTools/infrastructure/salesforce"
	"github.com/noworker/salesforceTools/test"
)

func TestGetSalesforceTokenByUserAndPassword(t *testing.T) {
	test.Setup()
	salesforceClient := salesforce.NewSalesforceClient()
	sat := model.SalesforceAccessToken{}
	salesforceAccessTokenResponse, err := salesforceClient.GetSalesforceTokenByUserAndPassword(
		sat,
		os.Getenv("SALESFORCE_MY_DOMAIN"),
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
