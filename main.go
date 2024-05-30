package main

import (
	"context"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	app := &cli.Command{
		Name:                  "cli",
		EnableShellCompletion: true,
	}

	cmd := &cli.Command{
		Name:                  "command",
		EnableShellCompletion: true,
	}

	flags := []cli.Flag{}

	flags = append(flags, &cli.StringFlag{
		Name:    "flag",
		Usage:   "flag",
		Aliases: []string{"f"},
	})

	subcmd := &cli.Command{
		Name:                  "subcommand",
		EnableShellCompletion: true,
		Flags:                 flags,
	}

	cmd.Commands = append(cmd.Commands, subcmd)
	app.Commands = append(app.Commands, cmd)

	app.Run(context.Background(), os.Args)
}
