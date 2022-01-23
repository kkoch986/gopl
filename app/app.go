package app

import (
	"log"
	"os"

	"github.com/hashicorp/logutils"
	"github.com/urfave/cli/v2"
)

func enableLogger(ctx *cli.Context) {
	out := os.Stdout
	flags := log.LstdFlags | log.Lmicroseconds | log.LUTC | log.Llongfile
	log.SetOutput(&logutils.LevelFilter{
		Levels:   []logutils.LogLevel{"VERBOSE", "DEBUG", "INFO", "WARN", "ERROR"},
		MinLevel: logutils.LogLevel(ctx.String("log-level")),
		Writer:   out,
	})

	log.SetFlags(flags)
}

func handleLogger(af cli.ActionFunc) cli.ActionFunc {
	return func(ctx *cli.Context) error {
		enableLogger(ctx)
		return af(ctx)
	}
}

var App = &cli.App{
	Name: "GOPL - Go Prolog",
	Commands: []*cli.Command{
		{
			Name:      "compile",
			Aliases:   []string{"c"},
			Usage:     "compile the given input file",
			ArgsUsage: "<filename>",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "outfile",
					Aliases: []string{"o"},
					Value:   "out.P",
				},
			},
			Action: handleLogger(compile),
		},
		{
			Name:    "shell",
			Aliases: []string{"s", ""},
			Usage:   "Enter the interactive query shell",
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:    "verbose",
					Aliases: []string{"vv"},
					Value:   false,
				},
			},
			Action: handleLogger(interactive),
		},
	},
}
