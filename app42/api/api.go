package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Param struct {
	ApiKey    string `json:"apiKey"`
	Version   string `json:"version"`
	TimeStamp string `json:"timeStamp"`
}

func NewGetRequest(method, path string) (request *http.Request) {
	request, err := http.NewRequest("GET", path, nil)

	if err != nil {
		return
	}

	return
}

func PerformRequestForBody(request *http.Request, response interface{}) (err error) {
	cli := &http.Client{}

	rawResponse, err := cli.Do(request)

	if err != nil {
		err = errors.New(fmt.Sprintf("Error performing request: %s", err.Error()))
	}

	jsonBytes, err := ioutil.ReadAll(rawResponse.Body)
	rawResponse.Body.Close()

	if err != nil {
		err = errors.New(fmt.Sprintf("Could not read response body: %s", err.Error()))
	}

	err = json.Unmarshal(jsonBytes, &response)

	if err != nil {
		err = errors.New(fmt.Sprintf("Invalid JSON response from server: %s", err.Error()))
	}
	return
}
