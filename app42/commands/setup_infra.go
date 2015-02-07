package commands

import (
	"encoding/json"
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/diatmpravin/app42_client/app42"
	"github.com/diatmpravin/app42_client/app42/api"
	"github.com/diatmpravin/app42_client/app42/base"
	"github.com/diatmpravin/app42_client/app42/constant"
	term "github.com/diatmpravin/app42_client/app42/terminal"
	"github.com/diatmpravin/app42_client/app42/util"
	"strconv"
)

var params map[string]string

type SetupInfra struct {
}

func NewSetupInfra() (s SetupInfra) {
	return
}

func (s SetupInfra) IsAppAvailable(appName string) (available bool, mes string) {
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
	if err != nil {
		fmt.Println("Failed", err)
	}
	available = response.App42.Response.Success
	mes = response.App42.Response.Description
	return
}

func (s SetupInfra) GetAppAndCheckAvailability(counter int) (appName string) {
	if counter <= 3 {
		counter++
		appName = base.AskAppName()
		available, mes := s.IsAppAvailable(appName)
		if available {
			return
		} else {
			term.Say(term.Red(mes))
			s.GetAppAndCheckAvailability(counter)
		}
	}

	return
}

func (s SetupInfra) chooseVMType(vmTypes []string) (vmType string) {
	for i, _ := range vmTypes {
		term.Say("%s: %s", term.Green(strconv.Itoa(i+1)), vmTypes[i])
	}

	index, err := strconv.Atoi(term.Ask("Select Instance Type>"))

	if err != nil || index > len(vmTypes) {
		term.Failed("Invalid number", err)
		return s.chooseVMType(vmTypes)
	}

	return vmTypes[index-1]
}

func (s SetupInfra) GetVMType(appName string) (vmType string) {
	path := constant.Host + constant.Version + "/info/subscription/app"
	secretKey, basicParams := base.Params()
	params = make(map[string]string)
	_ = json.Unmarshal([]byte(basicParams), &params)
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

	response := new(app42.AppSubscription)
	err = api.PerformRequestForBody(request, &response)
	if err != nil {
		fmt.Println("Failed", err)
	}

	vmType = s.chooseVMType(response.App42.Response.DeploymentType)
	return
}

func (s SetupInfra) chooseIaaS(iaass []string) (iaas string) {
	for i, _ := range iaass {
		term.Say("%s: %s", term.Green(strconv.Itoa(i+1)), iaass[i])
	}

	index, err := strconv.Atoi(term.Ask("Select IaaS Provider>"))

	if err != nil || index > len(iaass) {
		term.Failed("Invalid number", err)
		return s.chooseIaaS(iaass)
	}

	return iaass[index-1]
}

func (s SetupInfra) GetIaaSProviders(vmType string) (iaasId string) {
	path := constant.Host + constant.Version + "/info/iaasproviders/" + vmType
	secretKey, basicParams := base.Params()
	params = make(map[string]string)
	_ = json.Unmarshal([]byte(basicParams), &params)
	params["type"] = vmType
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

	response := new(app42.IaaSProviders)
	err = api.PerformRequestForBody(request, &response)
	if err != nil {
		fmt.Println("Failed", err)
	}

	iaass := []string{}
	iaasMap := make(map[string]string)

	for i := 0; i < len(response.App42.Response.Iaas); i++ {
		iaasMap[response.App42.Response.Iaas[i].Id] = response.App42.Response.Iaas[i].Name + " " + response.App42.Response.Iaas[i].Zone
		iaass = append(iaass, response.App42.Response.Iaas[i].Name+" "+response.App42.Response.Iaas[i].Zone)
	}

	iaas := s.chooseIaaS(iaass)

	for i := range iaasMap {
		if iaasMap[i] == iaas {
			iaasId = i
		}
	}

	return
}

func (s SetupInfra) Run(c *cli.Context) {
	appName := s.GetAppAndCheckAvailability(1)
	vmType := s.GetVMType(appName)
	iaasId := s.GetIaaSProviders(vmType)
	fmt.Println(iaasId)
}
