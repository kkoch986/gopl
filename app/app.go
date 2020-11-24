package app

import (
	"errors"
	"log"
	"os"
	//"bufio"
	"fmt"

	"github.com/urfave/cli/v2"

	"github.com/kkoch986/gopl/ast"
	"github.com/kkoch986/gopl/lexer"
	"github.com/kkoch986/gopl/parser"
	"github.com/kkoch986/gopl/raw"
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
					cli.ShowAppHelp(c)
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
			Flags:   []cli.Flag{},
			Action: func(c *cli.Context) error {
				shell := &QueryCLI{}
				err := shell.Run()
				if err != nil {
					log.Fatal(err)
					return err
				}
				return nil
			},
		},
	},
}
