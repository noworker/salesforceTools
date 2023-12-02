package usecase

import (
	salesforce "github.com/noworker/salesforceTools/infrastructure/salesforceService"
)

type ISalesforceUsecase interface {
	GetDebugLogs() error
}

type SalesforceUsecase struct {
	sc salesforce.ISalesforceClient
}

func NewSalesforceUsecase(sc salesforce.ISalesforceClient) ISalesforceUsecase {
	return &SalesforceUsecase{sc}
}

func (su *SalesforceUsecase) GetDebugLogs() error {
	return nil
}
