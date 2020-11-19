package main

import (
	"fmt"

	"github.com/kkoch986/gopl/ast"
	"github.com/kkoch986/gopl/lexer"
	"github.com/kkoch986/gopl/parser"
)

func main() {
	l := lexer.NewFile("./test.npl")
	if bsrSet, errs := parser.Parse(l); len(errs) > 0 {
		fmt.Println("ERR", errs)
	} else {
		a := ast.BuildStatementList(bsrSet.GetRoot())
		fmt.Println(a)
	}
}
