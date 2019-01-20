package handler

import (
	"github.com/suzuki-shunsuke/go-cliutil"
	"github.com/urfave/cli"

	"github.com/suzuki-shunsuke/gomic/internal/infra"
	"github.com/suzuki-shunsuke/gomic/internal/usecase/initcmd"
)

var initCommand = cli.Command{
	Name:   "init",
	Usage:  "create a configuration file if it doesn't exist",
	Action: initAction,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "dest, d",
			Usage: "created configuration file path",
			Value: ".gomic.yml",
		},
	},
}

func initAction(c *cli.Context) error {
	return cliutil.ConvErrToExitError(
		initcmd.Main(infra.FileSystem{}, c.String("dest")))
}
