package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/diatmpravin/app42_client/app42/constant"
	"github.com/diatmpravin/app42_client/app42/util"
	"io/ioutil"
	"net/http"
)

type Param struct {
	ApiKey    string `json:"apiKey"`
	Version   string `json:"version"`
	TimeStamp string `json:"timeStamp"`
}

func NewGetRequest(method, path, secretKey, params string) (response *http.Response) {
	signature := util.Sign(secretKey, params)

	var pro Param
	_ = json.Unmarshal([]byte(params), &pro)

	request, err := http.NewRequest("GET", path, nil)

	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("params", string(params))
	request.Header.Set("apiKey", pro.ApiKey)
	request.Header.Set("version", constant.Version)
	request.Header.Set("timeStamp", pro.TimeStamp)
	request.Header.Set("signature", signature)

	response, err = PerformRequestForBody(request, nil)
	if err != nil {
		return
	}
	return
}

func PerformRequestForBody(request *http.Request, response interface{}) (rawResponse *http.Response, err error) {
	cli := &http.Client{}

	rawResponse, err = cli.Do(request)

	if err != nil {
		err = errors.New(fmt.Sprintf("Error performing request: %s", err.Error()))
		return
	}

	jsonBytes, err := ioutil.ReadAll(rawResponse.Body)
	rawResponse.Body.Close()

	if err != nil {
		err = errors.New(fmt.Sprintf("Could not read response body: %s", err.Error()))
		return
	}

	err = json.Unmarshal(jsonBytes, &response)

	fmt.Println("response==============>", response)

	if err != nil {
		err = errors.New(fmt.Sprintf("Invalid JSON response from server: %s", err.Error()))
	}

	return
}
