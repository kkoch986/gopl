package resolver

import (
	"reflect"
	"testing"

	"github.com/kkoch986/gopl/ast"
	"github.com/kkoch986/gopl/indexer"
)

type resolverTestCase struct {
	ExistingStatements []ast.Statement
	Input              ast.Statement
	ExistingBindings   *Bindings
	ExpectedBindings   []*Bindings
}

func TestResolverCases(t *testing.T) {
	cases := []resolverTestCase{
		// Basic test:
		//   f(a,b).
		//   ?- f(A,B).
		// expect one binding A: a, B: b
		resolverTestCase{
			[]ast.Statement{
				&ast.Fact{Head: "f", Args: []ast.Term{ast.CreateAtom("a"), ast.CreateAtom("b")}}, // f(a,b)
			},
			&ast.Query{
				&ast.Fact{Head: "f", Args: []ast.Term{ast.CreateVariable("A"), ast.CreateVariable("B")}}, // f(A,B)
			},
			EmptyBindings(),
			[]*Bindings{
				CreateBindings(map[string]ast.Term{
					"A": ast.CreateAtom("a"),
					"B": ast.CreateAtom("b"),
				}),
			},
		},
	}

	for _, v := range cases {
		// create an indexer and index the original statements
		i := indexer.NewDefault()
		for _, s := range v.ExistingStatements {
			i.IndexStatement(s)
		}

		// create a resolver and resolve the input statement
		r := New(i)
		out := make(chan *Bindings, 1)
		r.ResolveStatementList([]ast.Statement{v.Input}, v.ExistingBindings, out)

		results := []*Bindings{}
		for outputBinding := range out {
			results = append(results, outputBinding)
		}

		if len(results) != len(v.ExpectedBindings) {
			t.Errorf("Incorrect number of bindings found. expected %s, got %s", v.ExpectedBindings, results)
		}

		// verify that all of the bindings in expected bindings are present in results
		for _, e := range v.ExpectedBindings {
			found := false
			for _, r := range results {
				if reflect.DeepEqual(*r, *e) {
					found = true
					break
				}
			}

			if !found {
				t.Errorf("expected binding not found (%s)", e)
			}
		}
	}
}
