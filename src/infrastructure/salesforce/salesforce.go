package salesforce

type ISalesforceClient interface {
	GetDebugLogs()
}

type SalesforceClient struct {
}

func NewSalesforceClient() ISalesforceClient {
	return &SalesforceClient{}
}

func (client *SalesforceClient) GetDebugLogs() {

}
