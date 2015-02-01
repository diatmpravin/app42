package commands

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/diatmpravin/app42_client/app42/api"
	"github.com/diatmpravin/app42_client/app42/base"
	"github.com/diatmpravin/app42_client/app42/constant"
	"net/http"
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

func (a Apps) findAllApps(url string) (response *http.Response) {

	secretKey, params := base.Params()
	response = api.NewGetRequest("GET", url, secretKey, string(params))

	return
}
