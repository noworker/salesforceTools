package salesforce

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/noworker/salesforceTools/domain/model"
	"github.com/noworker/salesforceTools/test"
)

type ISalesforceClient interface {
	GetSalesforceTokenByUserAndPassword(model model.SalesforceAccessTokenModel, salesforceUrl string, clientId string, clientSecret string, username string, password string) (model.SalesforceAccessTokenModel, error)
	GetSalesforceDescribeGlobal(model model.SalesforceDescribeGlobalModel, salesforceMyDomain string, salesforceAccessToken string) (model.SalesforceDescribeGlobalModel, error)
	GetDebugLogs() error
}

type SalesforceClient struct {
}

func NewSalesforceClient() ISalesforceClient {
	return &SalesforceClient{}
}

// salesforceのアクセストークンを取得する関数
func (cl *SalesforceClient) GetSalesforceTokenByUserAndPassword(
	model model.SalesforceAccessTokenModel,
	salesforceUrl string,
	clientId string,
	clientSecret string,
	username string,
	password string,
) (model.SalesforceAccessTokenModel, error) {
	// クライアントを作成
	client := &http.Client{}
	// リクエストボディを作成
	data := url.Values{}
	// エンドポイントを定義
	endpoint := fmt.Sprintf("%s%s", salesforceUrl, "/services/oauth2/token")
	data.Set("client_id", clientId)
	data.Set("client_secret", clientSecret)
	data.Set("username", username)
	data.Set("password", password)
	data.Set("grant_type", "password")
	// リクエスト内容を作成
	req, err := http.NewRequest(
		echo.POST,
		endpoint,
		strings.NewReader(data.Encode()),
	)
	if err != nil {
		return model, err
	}
	// ヘッダーを作成
	header := http.Header{}
	header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header = header
	test.HttpRequestDump(req)
	res, err := client.Do(req)
	test.HttpResponseDump(res)
	if err != nil {
		return model, err
	}
	if res.StatusCode != 200 {
		bodyBytes, _ := io.ReadAll(res.Body)
		return model, fmt.Errorf("[salesforce return error], %s", string(bodyBytes))
	}
	err = json.NewDecoder(res.Body).Decode(&model)
	if err != nil {
		return model, err
	}

	return model, nil
}

/*
# salesforceのSオブジェクトの一覧とメタ情報を取得する関数
*/
func (cl *SalesforceClient) GetSalesforceDescribeGlobal(model model.SalesforceDescribeGlobalModel, salesforceMyDomain string, salesforceAccessToken string) (model.SalesforceDescribeGlobalModel, error) {
	client := &http.Client{}
	header := http.Header{}
	header.Set("Content-Type", "application/json")
	header.Set("Authorization", "Bearer "+salesforceAccessToken)
	endpoint := fmt.Sprintf("%s%s", salesforceMyDomain, "/services/data/v56.0/sobjects/")
	req, _ := http.NewRequest(
		echo.GET,
		endpoint,
		nil,
	)
	req.Header = header
	test.HttpRequestDump(req)
	res, err := client.Do(req)
	test.HttpResponseDump(res)
	if err != nil {
		return model, err
	}
	if res.StatusCode != 200 {
		bodyBytes, _ := io.ReadAll(res.Body)
		return model, fmt.Errorf("[salesforce return error] %s", string(bodyBytes))
	}
	err = json.NewDecoder(res.Body).Decode(&model)
	if err != nil {
		return model, err
	}
	return model, nil
}

func (cl *SalesforceClient) GetDebugLogs() error {
	client := &http.Client{}
	header := http.Header{}
	req, _ := http.NewRequest(
		echo.POST,
		"localhost",
		nil,
	)
	req.Header = header
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	fmt.Println(res.Body)
	return nil
}
