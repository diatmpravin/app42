package commands

import (
	"github.com/codegangsta/cli"
	"github.com/diatmpravin/app42_client/app42/configuration"
)

func Keys(c *cli.Context) {
	configuration.ShowKeys()
}
