package commands

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/diatmpravin/cli/app42/configuration"
	term "github.com/diatmpravin/cli/app42/terminal"
)

type ClearKeys struct {
}

func NewClearKeys() (ck ClearKeys) {
	return
}

func (ck ClearKeys) Run(c *cli.Context) {
	_, err := configuration.Load()
	if err != nil {
		fmt.Println("Error loading configuration", err)
		return
	}

	ack := term.Ask("Do you want to delete existing keys? [Yn]:")

	if ack == "Y" || ack == "y" {
		err = configuration.DeleteKeys()

		if err != nil {
			fmt.Println("Failed logging out", err)
			return
		}

		//ck.ui.Ok()
	}

}
