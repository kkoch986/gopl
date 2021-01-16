package app

import (
	"errors"
	"fmt"
	"log"
	"os"
    "io/ioutil"

	"github.com/urfave/cli/v2"

	"github.com/kkoch986/gopl/ast"
	"github.com/kkoch986/gopl/indexer"
	"github.com/kkoch986/gopl/lexer"
	"github.com/kkoch986/gopl/parser"
	"github.com/kkoch986/gopl/raw"
	"github.com/kkoch986/gopl/resolver"
)

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
			Action: func(c *cli.Context) error {
				filename := c.Args().First()
				outfile := c.String("outfile")

				if filename == "" {
					_ = cli.ShowAppHelp(c)
					return errors.New("Filename is required")
				}
				fmt.Printf("Compiling %s\n", filename)

				// Parse the file
				l := lexer.NewFile(filename)
				if bsrSet, errs := parser.Parse(l); len(errs) > 0 {
					log.Fatal(errs)
					return errors.New("Parser Error")
				} else {
					a := ast.BuildStatementList(bsrSet.GetRoot())
					fmt.Println(a)

					// open the file for writing
					f, err := os.Create(outfile)
					if err != nil {
						return err
					}
					defer f.Close()

					err = raw.Serialize(a, f)
					if err != nil {
						return err
					}
				}
				return nil
			},
		},
		{
			Name:    "shell",
			Aliases: []string{"s", ""},
			Usage:   "Enter the interactive query shell",
			Flags:   []cli.Flag{
                &cli.BoolFlag{
                    Name: "verbose",
                    Aliases: []string{"vv"},
                    Value: false,
                },
            },
			Action: func(c *cli.Context) error {
				i := indexer.NewDefault()
                if !c.Bool("verbose") {
                    log.SetOutput(ioutil.Discard)
                }
				// TODO: let flags define these
				h, err := NewHistory(os.Getenv("HOME")+"/.gopl_history", 1000)

				if err != nil {
					log.Fatal(err)
					return err
				}

				shell := &QueryCLI{
					I: i,
					R: resolver.New(i),
					H: h,
				}
				err = shell.Run()
				if err != nil {
					log.Fatal(err)
					return err
				}
				return nil
			},
		},
	},
}
