package commands

import (
	"github.com/codegangsta/cli"
	"github.com/diatmpravin/cli/app42/configuration"
)

func Keys(c *cli.Context) {
	configuration.ShowKeys()
}
