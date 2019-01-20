package handler

import (
	"os"

	"github.com/urfave/cli"

	"github.com/suzuki-shunsuke/gomic/internal/domain"
)

// Gomic calls a command.
func Gomic() {
	app := cli.NewApp()
	app.Name = "gomic"
	app.Version = domain.Version
	app.Author = "suzuki-shunsuke https://github.com/suzuki-shunsuke"
	app.Usage = "generate golang's mock for test"

	app.Commands = []cli.Command{
		initCommand,
		genCommand,
	}
	app.Run(os.Args)
}
