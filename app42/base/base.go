package base

import (
	"encoding/json"
	"fmt"
	"github.com/diatmpravin/app42_client/app42/configuration"
	"github.com/diatmpravin/app42_client/app42/constant"
	term "github.com/diatmpravin/app42_client/app42/terminal"
	"strings"
	"time"
)

var basicParams map[string]string

func Params() (secretKey string, params []byte) {

	config, err := configuration.Load()

	if err != nil {
		term.Failed("File is invalid", err)
		return

	}

	basicParams = make(map[string]string)

	basicParams["apiKey"] = config.ApiKey
	basicParams["version"] = constant.Version
	basicParams["timeStamp"] = TimeStampUTC()

	params, err = json.Marshal(basicParams)
	secretKey = config.SecretKey

	if err != nil {
		fmt.Println("error:", err)
	}

	return
}

func TimeStampUTC() string {
	currentTime := time.Now().UTC()
	formatedTime := currentTime.Format(time.RFC3339)
	dateTime := strings.Split(formatedTime, "T")
	stampMilli := currentTime.Format(time.StampMilli)
	timeArray := strings.Split(stampMilli, " ")

	s := []string{dateTime[0], timeArray[len(timeArray)-1]}
	final := strings.Join(s, "T")
	return final + "Z"
}

func AskAppName() (appName string) {
	appName = term.Ask(term.Yellow("Enter App Name:"))
	return
}
