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
	path := constant.Host + constant.Version + "/app/availability"
	secretKey, params := base.Params()
	request := api.NewGetRequest("GET", path)

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

	fmt.Println("--------------->", response)
}
