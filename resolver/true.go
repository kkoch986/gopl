package resolver

import (
	"github.com/kkoch986/gopl/ast"
)

/**
 * True will terminate with no matched bindings.
 */
type True struct{}

func (w *True) Resolve(fact *ast.Fact, c *Bindings, out chan<- *Bindings, m chan<- bool) {
	if fact.Signature().String() != "true/0" {
		m <- false
		return
	}
	defer close(out)
	defer close(m)
	out <- c
	m <- true
}
