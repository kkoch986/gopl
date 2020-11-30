package resolver

import (
	"github.com/kkoch986/gopl/ast"
	"github.com/kkoch986/gopl/indexer"
)

/**
 * Assert/1 will take the given fact and insert it into the current indexed universe.
 */
type Assert struct {
	idx indexer.Indexer
}

func (w *Assert) Resolve(fact *ast.Fact, c *Bindings, out chan<- *Bindings, m chan<- bool) {
	if fact.Signature().String() != "assert/1" {
		m <- false
		return
	}
	defer close(out)
	defer close(m)

	w.idx.IndexStatement(fact.Args[0])

	out <- c
	m <- true
}
