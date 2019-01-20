package handler

import (
	"github.com/suzuki-shunsuke/go-cliutil"
	"github.com/urfave/cli"

	"github.com/suzuki-shunsuke/gomic/internal/infra"
	"github.com/suzuki-shunsuke/gomic/internal/usecase/gencmd"
)

var genCommand = cli.Command{
	Name:   "gen",
	Usage:  "generate mocks",
	Action: genAction,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "config, c",
			Usage: "configuration file path",
			Value: "",
		},
	},
}

func genAction(c *cli.Context) error {
	return cliutil.ConvErrToExitError(
		gencmd.Main(infra.FileSystem{}, infra.Importer{},
			infra.CfgReader{}, c.String("config")))
}
