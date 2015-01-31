package commands

import (
	"github.com/codegangsta/cli"
	"github.com/diatmpravin/app42_client/app42/configuration"
	term "github.com/diatmpravin/app42_client/app42/terminal"
)

type AddKeys struct {
}

func NewKeys() (k AddKeys) {
	return
}

func (k AddKeys) Run(c *cli.Context) {
	apiKey := term.Ask(term.Yellow("Enter API Key::"))
	secretKey := term.Ask(term.Yellow("Enter Secret Key::"))

	_, err := k.saveKeys(apiKey, secretKey)

	if err != nil {
		term.Failed("Error saving configuration", err)
		return
	}

	term.Say(term.Green("Adding keys...done"))
}

func (k AddKeys) saveKeys(apiKey, secretKey string) (config *configuration.Keys, err error) {
	config = new(configuration.Keys)
	config.ApiKey = apiKey
	config.SecretKey = secretKey
	err = config.Save()
	return
}
