package resolver

import (
	"github.com/kkoch986/gopl/ast"
)

/**
 * load/1 will be used to load source files or compiled source files into the execution environment for querying
 */
type Load struct{}

func (w *Load) Resolve(fact *ast.Fact, c *Bindings, out chan<- *Bindings, m chan<- bool) {
	if fact.Signature().String() != "load/1" {
		m <- false
		return
	}
	defer close(out)
	defer close(m)

	// TODO: load the given file and index it

	m <- true
	out <- c
}
