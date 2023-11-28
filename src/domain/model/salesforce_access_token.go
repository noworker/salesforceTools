package model

type SalesforceAccessToken struct {
	AccessToken string `json:"access_token"`
	InstanceUrl string `json:"instance_url"`
	Id          string `json:"id"`
	TokenType   string `json:"token_type"`
	IssuedAt    string `json:"issued_at"`
	Signature   string `json:"signature"`
}

/**
example.
{
    "access_token": "00D5j00000DhxaS!AR0AQNzaJfMiNAikAgsWADX2EdoMJu780It1p55vjvZJJGT1hCmgEuEuv1m1apt7ymQD5YqLSlJ0EYXrtyrIt36jDRewrh8X",
    "instance_url": "htrailblaze.my.salesforce.com",
    "id": "https://login.salesforce.com/id/00D5j00000DhxaSEAR/0055j00000AZAQlAAP",
    "token_type": "Bearer",
    "issued_at": "1701160824576",
    "signature": "9fjva8Vcz34aAylIMl9Vhg4dL7bcve6DkBFCyhN8tog="
}
**/
