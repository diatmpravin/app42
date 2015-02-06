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

var params map[string]string

type SetupInfra struct {
}

func NewSetupInfra() (s SetupInfra) {
	return
}

func (s SetupInfra) Run(c *cli.Context) {
	appName := base.AskAppName()
	s.checkAppAvailability(appName)
	fmt.Println(appName)
}

func (s SetupInfra) checkAppAvailability(appName string) {
	path := constant.Host + constant.Version + "/app/availability" + "?appName=" + appName
	secretKey, basicParams := base.Params()

	params = make(map[string]string)
	_ = json.Unmarshal([]byte(basicParams), &params)

	params["appName"] = appName

	request := api.NewGetRequest("GET", path)

	queryParams, err := json.Marshal(params)

	signature := util.Sign(secretKey, string(queryParams))

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("params", string(queryParams))
	request.Header.Set("apiKey", params["apiKey"])
	request.Header.Set("version", constant.Version)
	request.Header.Set("timeStamp", params["timeStamp"])
	request.Header.Set("signature", signature)

	response := new(app42.AppAvailability)

	err = api.PerformRequestForBody(request, &response)

	fmt.Printf("%+v\n", response)

	if err != nil {
		fmt.Println("Failed", err)
		return
	}
}
