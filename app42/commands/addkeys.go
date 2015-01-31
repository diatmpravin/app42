package commands

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/diatmpravin/cli/app42/configuration"
	term "github.com/diatmpravin/cli/app42/terminal"
)

type AddKeys struct {
}

func NewKeys() (k AddKeys) {
	return
}

func (k AddKeys) Run(c *cli.Context) {
	apiKey := term.Ask("Enter API Key::")
	secretKey := term.Ask("Enter Secret Key::")

	_, err := k.saveKeys(apiKey, secretKey)

	if err != nil {
		fmt.Println("Error saving configuration", err)
	}

	term.Say("Adding keys...done")
}

func (k AddKeys) saveKeys(apiKey, secretKey string) (config *configuration.Keys, err error) {
	config = new(configuration.Keys)
	config.ApiKey = apiKey
	config.SecretKey = secretKey
	err = config.Save()
	return
}
