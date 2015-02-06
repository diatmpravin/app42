package commands

import (
	"encoding/json"
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/diatmpravin/app42_client/app42"
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

	a.findAllApps(path)

}

func (a Apps) findAllApps(url string) {

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

	response := new(app42.AllApps)

	err := api.PerformRequestForBody(request, &response)

	fmt.Printf("%+v\n", response)

	if err != nil {
		fmt.Println("Failed", err)
		return
	}

	return
}
