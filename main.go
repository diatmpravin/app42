package main

import (
	"encoding/json"
	"fmt"
	"github.com/diatmpravin/cli_sample/util"
	"net/http"
)

const (
	version   = "1.0"
	apiKey    = "ea71af1b0d732835af33e7bb1a7b7984ee705169c7c395b79ad4d6a06cd8f246"
	secretKey = "57aeabbba4864e3c5b39f2b2f39c27d1dcfa4a0fef9e8c17bcc67bd0ee815d29"
	host      = "https://paashq.shephertz.com/paas/1.0/app"
)

type Param struct {
	ApiKey    string `json:"apiKey"`
	Version   string `json:"version"`
	TimeStamp string `json:"timeStamp"`
}

func main() {
	time := util.TimeStamp()
	fmt.Println(time)

	p := &Param{
		ApiKey:    apiKey,
		Version:   version,
		TimeStamp: time}

	params, err := json.Marshal(p)

	if err != nil {
		fmt.Println("Json Encoding Error:", err)
	}

	signature := util.Sign(secretKey, string(params))

	fmt.Println("Final Signature==>", signature)
	fmt.Println("Params ==========>", string(params))

	request, err := http.NewRequest("GET", host, nil)

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("params", string(params))
	request.Header.Set("apiKey", apiKey)
	request.Header.Set("version", version)
	request.Header.Set("timeStamp", time)
	request.Header.Set("signature", signature)

	fmt.Println(request)
	fmt.Println(err)

	cli := &http.Client{}

	rawResponse, err := cli.Do(request)
	fmt.Println("rawResponse====>", rawResponse)
	fmt.Println(err)

	//return &AuthorizedRequest{request}, err

}
