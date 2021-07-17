package resolver_test

import (
	"reflect"
	"testing"

	"github.com/kkoch986/gopl/ast"
	"github.com/kkoch986/gopl/indexer"
	"github.com/kkoch986/gopl/resolver"
)

type resolverTestCase struct {
	Label              string
	ExistingStatements []ast.Statement
	Input              ast.Statement
	ExistingBindings   *resolver.Bindings
	ExpectedBindings   []*resolver.Bindings
}

func runTestCase(t *testing.T, v resolverTestCase) {
	t.Logf("Starting Test %s", v.Label)
	// create an indexer and index the original statements
	i := indexer.NewDefault()
	for _, s := range v.ExistingStatements {
		i.IndexStatement(s)
	}

	// create a resolver and resolve the input statement
	r := resolver.New(i)
	out := make(chan *resolver.Bindings, 1)
	go r.ResolveStatementList([]ast.Statement{v.Input}, v.ExistingBindings, out)

	results := []*resolver.Bindings{}
	for outputBinding := range out {
		results = append(results, outputBinding)
	}

	if len(results) != len(v.ExpectedBindings) {
		t.Errorf("Incorrect number of bindings found. expected %d, got %d (results: \n%v)", len(v.ExpectedBindings), len(results), results)
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

func TestResolverCases(t *testing.T) {
	cases := []resolverTestCase{
		// Basic test:
		//   f(a,b).
		//   ?- f(A,B).
		// expect one binding A: a, B: b
		resolverTestCase{
			"Basic f(A,B)",
			[]ast.Statement{
				ast.CreateFact("f", ast.CreateAtom("a"), ast.CreateAtom("b")), // f(a,b).
			},
			&ast.Query{
				ast.CreateFact("f", ast.CreateVariable("A"), ast.CreateVariable("B")), // f(A,B).
			},
			resolver.EmptyBindings(),
			[]*resolver.Bindings{
				resolver.CreateBindings(map[string]ast.Term{
					"A": ast.CreateAtom("a"),
					"B": ast.CreateAtom("b"),
				}),
			},
		},
		// shallow friend of a friend test:
		// f(a,b).
		// f(b,c).
		// f(c,d).
		// f(d,e).
		// fof(A,B) := f(A,C), f(C,B).
		// ?- fof(A,C).
		// expect two bindings:
		//   A: a, B: c
		//   A: b, B: d
		//   A: c, B: e
		resolverTestCase{
			"Shallow FOF",
			[]ast.Statement{
				ast.CreateFact("f", ast.CreateAtom("a"), ast.CreateAtom("b")),
				ast.CreateFact("f", ast.CreateAtom("b"), ast.CreateAtom("c")),
				ast.CreateFact("f", ast.CreateAtom("c"), ast.CreateAtom("d")),
				ast.CreateFact("f", ast.CreateAtom("d"), ast.CreateAtom("e")),
				ast.CreateRule(
					ast.CreateFact("fof", ast.CreateVariable("A"), ast.CreateVariable("B")),
					// :=
					ast.CreateFact("f", ast.CreateVariable("A"), ast.CreateVariable("C")),
					ast.CreateFact("f", ast.CreateVariable("C"), ast.CreateVariable("B")),
				),
			},
			&ast.Query{
				ast.CreateFact("fof", ast.CreateVariable("A"), ast.CreateVariable("B")),
			},
			resolver.EmptyBindings(),
			[]*resolver.Bindings{
				resolver.CreateBindings(map[string]ast.Term{
					"A": ast.CreateAtom("a"),
					"B": ast.CreateAtom("c"),
				}),
				resolver.CreateBindings(map[string]ast.Term{
					"A": ast.CreateAtom("b"),
					"B": ast.CreateAtom("d"),
				}),
				resolver.CreateBindings(map[string]ast.Term{
					"A": ast.CreateAtom("c"),
					"B": ast.CreateAtom("e"),
				}),
			},
		},
		// Recursive fof
		// tests that we can resolve deeply recursive rules
		// f(a,b). f(b,c). f(c,d). f(d,e).
		// fof(A,C) := f(A,C).
		// fof(A,C) := f(A,B), fof(B,C).
		// expect these bindings:
		// A: a, C: b
		// A: a, C: c
		// A: a, C: d
		// A: a, C: e
		// A: b, C: c
		// A: b, C: d
		// A: b, C: e
		// A: c, C: d
		// A: c, C: e
		// A: d, C: e
		resolverTestCase{
			"Recusrive FOF",
			[]ast.Statement{
				ast.CreateFact("f", ast.CreateAtom("a"), ast.CreateAtom("b")),
				ast.CreateFact("f", ast.CreateAtom("b"), ast.CreateAtom("c")),
				ast.CreateFact("f", ast.CreateAtom("c"), ast.CreateAtom("d")),
				ast.CreateFact("f", ast.CreateAtom("d"), ast.CreateAtom("e")),
				ast.CreateRule(
					ast.CreateFact("fof", ast.CreateVariable("A"), ast.CreateVariable("C")),
					// :=
					ast.CreateFact("f", ast.CreateVariable("A"), ast.CreateVariable("C")),
				),
				ast.CreateRule(
					ast.CreateFact("fof", ast.CreateVariable("A"), ast.CreateVariable("C")),
					// :=
					ast.CreateFact("f", ast.CreateVariable("A"), ast.CreateVariable("B")),
					ast.CreateFact("fof", ast.CreateVariable("B"), ast.CreateVariable("C")),
				),
			},
			&ast.Query{
				ast.CreateFact("fof", ast.CreateVariable("A"), ast.CreateVariable("C")),
			},
			resolver.EmptyBindings(),
			[]*resolver.Bindings{
				resolver.CreateBindings(map[string]ast.Term{"A": ast.CreateAtom("a"), "C": ast.CreateAtom("b")}),
				resolver.CreateBindings(map[string]ast.Term{"A": ast.CreateAtom("a"), "C": ast.CreateAtom("c")}),
				resolver.CreateBindings(map[string]ast.Term{"A": ast.CreateAtom("a"), "C": ast.CreateAtom("d")}),
				resolver.CreateBindings(map[string]ast.Term{"A": ast.CreateAtom("a"), "C": ast.CreateAtom("e")}),
				resolver.CreateBindings(map[string]ast.Term{"A": ast.CreateAtom("b"), "C": ast.CreateAtom("c")}),
				resolver.CreateBindings(map[string]ast.Term{"A": ast.CreateAtom("b"), "C": ast.CreateAtom("d")}),
				resolver.CreateBindings(map[string]ast.Term{"A": ast.CreateAtom("b"), "C": ast.CreateAtom("e")}),
				resolver.CreateBindings(map[string]ast.Term{"A": ast.CreateAtom("c"), "C": ast.CreateAtom("d")}),
				resolver.CreateBindings(map[string]ast.Term{"A": ast.CreateAtom("c"), "C": ast.CreateAtom("e")}),
				resolver.CreateBindings(map[string]ast.Term{"A": ast.CreateAtom("d"), "C": ast.CreateAtom("e")}),
			},
		},
	}

	for _, v := range cases {
		runTestCase(t, v)
	}
}
