package resolver

import (
	"fmt"

	"github.com/kkoch986/gopl/ast"
)

type Writeln struct {
	r *R
}

func (w *Writeln) Resolve(fact *ast.Fact, c *Bindings, out chan<- *Bindings, m chan<- bool) {
	if fact.Signature().String() != "writeln/1" {
		m <- false
		return
	}

	defer close(out)
	defer close(m)

	// TODO: use r to resolve the arg

	fmt.Println(fact.Args[0])
	out <- c
	m <- true
}
