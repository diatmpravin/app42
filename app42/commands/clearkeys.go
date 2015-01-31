package commands

import (
	"github.com/codegangsta/cli"
	"github.com/diatmpravin/app42_client/app42/configuration"
	term "github.com/diatmpravin/app42_client/app42/terminal"
)

type ClearKeys struct {
}

func NewClearKeys() (ck ClearKeys) {
	return
}

func (ck ClearKeys) Run(c *cli.Context) {
	_, err := configuration.Load()
	if err != nil {
		term.Failed("Error loading configuration", err)
		return
	}

	ack := term.Ask(term.Red("Do you want to delete existing keys? [Yn]:"))

	if ack == "Y" || ack == "y" {
		err = configuration.DeleteKeys()

		if err != nil {
			term.Failed("Failed logging out", err)
			return
		}

		term.Ok()
	}

}
