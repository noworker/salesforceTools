package test

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/joho/godotenv"
)

// 環境変数の読み取り テスト前処理
func LoadEnv(envPath string) {
	fmt.Println("Load env")
	err := godotenv.Load(envPath)
	if err != nil {
		log.Println(err)
	}
}

// HTTPリクエストのダンプを出力
func HttpRequestDump(req *http.Request) {
	dump, err := httputil.DumpRequest(req, true)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("---httpRequestDump---")
	fmt.Printf("%s\n", dump)
	fmt.Println("---httpRequestDump---")
}

func HttpResponseDump(res *http.Response) {
	dump, err := httputil.DumpResponse(res, true)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("---httpResponseDump---")
	fmt.Printf("%s\n", dump)
	fmt.Println("---httpResponseDump---")
}
