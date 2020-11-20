package main

import (
	//	"fmt"
	"log"
	"os"

	"github.com/kkoch986/gopl/app"
	//	"github.com/kkoch986/gopl/ast"
	//	"github.com/kkoch986/gopl/lexer"
	//	"github.com/kkoch986/gopl/parser"
)

func main() {
	err := app.App.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
	/*	l := lexer.NewFile("./test.npl")
		if bsrSet, errs := parser.Parse(l); len(errs) > 0 {
			fmt.Println("ERR", errs)
		} else {
			a := ast.BuildStatementList(bsrSet.GetRoot())
			fmt.Println(a)
		}*/
}
