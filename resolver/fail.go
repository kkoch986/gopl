package resolver

import (
	"github.com/kkoch986/gopl/ast"
)

/**
 * Fail will terminate with no matched bindings.
 */
type Fail struct{}

func (w *Fail) Resolve(fact *ast.Fact, c *Bindings, out chan<- *Bindings, m chan<- bool) {
	if fact.Signature().String() != "fail/0" {
		m <- false
		return
	}
	defer close(out)
	defer close(m)
	m <- true
}
