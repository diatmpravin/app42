package commands

import (
	"encoding/json"
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/diatmpravin/app42_client/app42/api"
	"github.com/diatmpravin/app42_client/app42/base"
	"github.com/diatmpravin/app42_client/app42/constant"
	"github.com/diatmpravin/app42_client/app42/util"
)

type Apps struct {
	Name string
}

type Param struct {
	ApiKey    string `json:"apiKey"`
	Version   string `json:"version"`
	TimeStamp string `json:"timeStamp"`
}

func NewApps() (a Apps) {
	return
}

func (a Apps) Run(c *cli.Context) {

	path := constant.Host + constant.Version + "/app"

	apps := a.findAllApps(path)
	fmt.Println("Response====>", apps)

}

func (a Apps) findAllApps(url string) (response interface{}) {

	secretKey, params := base.Params()
	request := api.NewGetRequest("GET", url)

	signature := util.Sign(secretKey, string(params))

	var pro Param
	_ = json.Unmarshal([]byte(params), &pro)

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("params", string(params))
	request.Header.Set("apiKey", pro.ApiKey)
	request.Header.Set("version", constant.Version)
	request.Header.Set("timeStamp", pro.TimeStamp)
	request.Header.Set("signature", signature)

	response, err := api.PerformRequestForBody(request, nil)
	if err != nil {
		return
	}

	return
}
