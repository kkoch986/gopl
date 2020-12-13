package resolver

import (
	"os"

	"github.com/kkoch986/gopl/ast"
	"github.com/kkoch986/gopl/indexer"
	"github.com/kkoch986/gopl/raw"
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

	// TODO: ground any input terms i.e. variables etc..

	if fact.Args[0].GetType() == ast.T_String {
		w.indexFile(fact.Args[0].String())
	} else {
		w.idx.IndexStatement(fact.Args[0])
	}

	out <- c
	m <- true
}

func (w *Assert) indexFile(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		// TODO: better error handling
		panic(err)
	}

	statements := make(chan ast.Statement)
	go raw.Deserialize(f, statements)
	for s := range statements {
		w.idx.IndexStatement(s)
	}
}
