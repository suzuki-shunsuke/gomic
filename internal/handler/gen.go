package handler

import (
	"github.com/urfave/cli"

	"github.com/suzuki-shunsuke/gomic/internal/infra"
	"github.com/suzuki-shunsuke/gomic/internal/usecase/gencmd"
)

// GenCommand is the sub command "gen".
var GenCommand = cli.Command{
	Name:   "gen",
	Usage:  "generate mocks",
	Action: Gen,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "config, c",
			Usage: "configuration file path",
			Value: "",
		},
	},
}

// Gen is the sub command "gen".
func Gen(c *cli.Context) error {
	return wrapUsecase(
		gencmd.Main(infra.FileSystem{}, infra.Importer{},
			infra.CfgReader{}, c.String("config")))
}
