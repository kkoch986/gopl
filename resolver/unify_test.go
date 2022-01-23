package resolver_test

import (
	"testing"

	"github.com/kkoch986/gopl/ast"
	"github.com/kkoch986/gopl/resolver"
)

type testCase struct {
	F                *ast.Fact
	ShouldMatch      bool
	InitialBindings  *resolver.Bindings
	ExpectedBindings []*resolver.Bindings
}

func runUnifyTestCase(c *testCase, t *testing.T) {
	r := &resolver.Equals{}
	m := make(chan bool)
	out := make(chan *resolver.Bindings)
	go r.Resolve(c.F, c.InitialBindings, out, m)

	bindings := []*resolver.Bindings{}
	done := false
	for !done {
		select {
		case b := <-out:
			bindings = append(bindings, b)
		case matched := <-m:
			done = true
			if matched != c.ShouldMatch {
				t.Errorf("unexpected matched value, expected %v, got %v", c.ShouldMatch, matched)
			}
		}
	}
	if len(bindings) != len(c.ExpectedBindings) {
		t.Fatalf(
			"incorrect number of bindings returned, expected %d, got %d",
			len(c.ExpectedBindings),
			len(bindings),
		)
	}

	// loop over the bindings and make sure they match up
	for i, b := range bindings {
		expected := c.ExpectedBindings[i]
		if !b.Equals(expected) {
			t.Errorf("failed asserting that \n%v equals expected \n%v", b, expected)
		}
	}
}

/**
 * TestMatching ensures that Equals/2 only matches the correct terms
 **/
func TestMatching(t *testing.T) {
	testCases := []*testCase{
		// notequals shouldnt match =/2
		{
			F:                ast.CreateFact("notequals"),
			ShouldMatch:      false,
			InitialBindings:  resolver.EmptyBindings(),
			ExpectedBindings: []*resolver.Bindings{},
		},
		// dont match =/0
		{
			F:                ast.CreateFact("="),
			ShouldMatch:      false,
			InitialBindings:  resolver.EmptyBindings(),
			ExpectedBindings: []*resolver.Bindings{},
		},
		// dont match =/1
		{
			F:                ast.CreateFact("=", ast.CreateAtom("a")),
			ShouldMatch:      false,
			InitialBindings:  resolver.EmptyBindings(),
			ExpectedBindings: []*resolver.Bindings{},
		},
		// dont match equals/2
		{
			F:                ast.CreateFact("equals", ast.CreateAtom("a"), ast.CreateAtom("b")),
			ShouldMatch:      false,
			InitialBindings:  resolver.EmptyBindings(),
			ExpectedBindings: []*resolver.Bindings{},
		},
		// dont match =/3
		{
			F:                ast.CreateFact("=", ast.CreateAtom("a"), ast.CreateAtom("b"), ast.CreateAtom("c")),
			ShouldMatch:      false,
			InitialBindings:  resolver.EmptyBindings(),
			ExpectedBindings: []*resolver.Bindings{},
		},
		// match =/2
		{
			F:                ast.CreateFact("=", ast.CreateAtom("a"), ast.CreateAtom("b")),
			ShouldMatch:      true,
			InitialBindings:  resolver.EmptyBindings(),
			ExpectedBindings: []*resolver.Bindings{},
		},
	}

	for _, c := range testCases {
		runUnifyTestCase(c, t)
	}
}

/**
 * TestBasicTermResolution tests basic cases of ATOM = ATOM
 * tests both when the atoms are equal and when they are not.
 * Also covers numeric literals and string literals
 **/
func TestBasicTermResolution(t *testing.T) {
	a := ast.CreateAtom("a")
	a2 := ast.CreateAtom("a")
	b := ast.CreateAtom("b")
	one := ast.CreateNumericLiteral(1)
	oneAlt := ast.CreateNumericLiteral(1)
	onepointone := ast.CreateNumericLiteral(1.1)
	ten := ast.CreateNumericLiteral(10)
	stringA := ast.CreateStringLiteral("a")
	stringA2 := ast.CreateStringLiteral("a")
	string10 := ast.CreateStringLiteral("10")

	testCases := []*testCase{
		// test that 2 identical atoms unify
		{
			F:               ast.CreateFact("=", a, a),
			ShouldMatch:     true,
			InitialBindings: resolver.EmptyBindings(),
			ExpectedBindings: []*resolver.Bindings{
				resolver.EmptyBindings(),
			},
		},
		// test that 2 equal atoms
		{
			F:               ast.CreateFact("=", a, a2),
			ShouldMatch:     true,
			InitialBindings: resolver.EmptyBindings(),
			ExpectedBindings: []*resolver.Bindings{
				resolver.EmptyBindings(),
			},
		},
		// test that 2 different atoms dont unify
		{
			F:                ast.CreateFact("=", a, b),
			ShouldMatch:      true,
			InitialBindings:  resolver.EmptyBindings(),
			ExpectedBindings: []*resolver.Bindings{},
		},
		// test that 2 numeric literals unify
		{
			F:               ast.CreateFact("=", one, oneAlt),
			ShouldMatch:     true,
			InitialBindings: resolver.EmptyBindings(),
			ExpectedBindings: []*resolver.Bindings{
				resolver.EmptyBindings(),
			},
		},
		// test that 2 different numeric literals dont unify
		{
			F:                ast.CreateFact("=", one, onepointone),
			ShouldMatch:      true,
			InitialBindings:  resolver.EmptyBindings(),
			ExpectedBindings: []*resolver.Bindings{},
		},
		// test that a numeric literal and matching string representation dont unify
		{
			F:                ast.CreateFact("=", ten, string10),
			ShouldMatch:      true,
			InitialBindings:  resolver.EmptyBindings(),
			ExpectedBindings: []*resolver.Bindings{},
		},
		// test that matching string literals unify
		{
			F:               ast.CreateFact("=", stringA, stringA2),
			ShouldMatch:     true,
			InitialBindings: resolver.EmptyBindings(),
			ExpectedBindings: []*resolver.Bindings{
				resolver.EmptyBindings(),
			},
		},
		// test that a string literal with the same value as an atom can unify
		{
			F:               ast.CreateFact("=", stringA, a),
			ShouldMatch:     true,
			InitialBindings: resolver.EmptyBindings(),
			ExpectedBindings: []*resolver.Bindings{
				resolver.EmptyBindings(),
			},
		},
	}

	for _, c := range testCases {
		runUnifyTestCase(c, t)
	}
}

/**
 * TestSimpleTermVariableBinding tests unifying a single variable with a simple term like an atom or literal.
 */
func TestSimpleTermVariableBinding(t *testing.T) {
	testCases := []*testCase{
		// basic atom unification VAR = ATOM
		{
			F:               ast.CreateFact("=", ast.CreateVariable("A"), ast.CreateAtom("a")),
			ShouldMatch:     true,
			InitialBindings: resolver.EmptyBindings(),
			ExpectedBindings: []*resolver.Bindings{
				resolver.CreateBindings(map[string]ast.Term{
					"A": ast.CreateAtom("a"),
				}),
			},
		},
		// test that order doesnt matter (ATOM = VAR)
		{
			F:               ast.CreateFact("=", ast.CreateAtom("a"), ast.CreateVariable("A")),
			ShouldMatch:     true,
			InitialBindings: resolver.EmptyBindings(),
			ExpectedBindings: []*resolver.Bindings{
				resolver.CreateBindings(map[string]ast.Term{
					"A": ast.CreateAtom("a"),
				}),
			},
		},
		// test that VAR = VAR results in a ground term when the LHS is already bound
		{
			F:           ast.CreateFact("=", ast.CreateVariable("A"), ast.CreateVariable("B")),
			ShouldMatch: true,
			InitialBindings: resolver.CreateBindings(map[string]ast.Term{
				"A": ast.CreateAtom("a"),
			}),
			ExpectedBindings: []*resolver.Bindings{
				resolver.CreateBindings(map[string]ast.Term{
					"A": ast.CreateAtom("a"),
					"B": ast.CreateAtom("a"),
				}),
			},
		},
		// test that VAR = VAR results in a ground term when the RHS is already bound
		{
			F:           ast.CreateFact("=", ast.CreateVariable("A"), ast.CreateVariable("B")),
			ShouldMatch: true,
			InitialBindings: resolver.CreateBindings(map[string]ast.Term{
				"B": ast.CreateAtom("a"),
			}),
			ExpectedBindings: []*resolver.Bindings{
				resolver.CreateBindings(map[string]ast.Term{
					"A": ast.CreateAtom("a"),
					"B": ast.CreateAtom("a"),
				}),
			},
		},
		// test that VAR = VAR results in the lowest alpha VAR term bound when both are unbound
		{
			F:               ast.CreateFact("=", ast.CreateVariable("A"), ast.CreateVariable("B")),
			ShouldMatch:     true,
			InitialBindings: resolver.CreateBindings(map[string]ast.Term{}),
			ExpectedBindings: []*resolver.Bindings{
				resolver.CreateBindings(map[string]ast.Term{
					"A": ast.CreateVariable("B"),
				}),
			},
		},
		// test that VAR = VAR results in the lowest alpha VAR term bound when both are unbound
		{
			F:               ast.CreateFact("=", ast.CreateVariable("B"), ast.CreateVariable("A")),
			ShouldMatch:     true,
			InitialBindings: resolver.CreateBindings(map[string]ast.Term{}),
			ExpectedBindings: []*resolver.Bindings{
				resolver.CreateBindings(map[string]ast.Term{
					"A": ast.CreateVariable("B"),
				}),
			},
		},
		// Test that VAR = VAR results in the lowest alpha VAR term bound when there is a chain of bound VARs
		// For example if A = B and C = D, then binding B = C should result in B bound to D
		//  since we bind the lowest alpha variable to the most ground form of the other term.
		{
			F:           ast.CreateFact("=", ast.CreateVariable("C"), ast.CreateVariable("B")),
			ShouldMatch: true,
			InitialBindings: resolver.CreateBindings(map[string]ast.Term{
				"A": ast.CreateVariable("B"),
				"C": ast.CreateVariable("D"),
			}),
			ExpectedBindings: []*resolver.Bindings{
				resolver.CreateBindings(map[string]ast.Term{
					"A": ast.CreateVariable("B"),
					"B": ast.CreateVariable("D"),
					"C": ast.CreateVariable("D"),
				}),
			},
		},
		// Test that when a variable is bound to another variable, it is fully grounded first
		// For example if you have the binding:
		//   A -> B, B -> C, C -> "a"
		//  and you bind D = A, the binding created should be D -> "a"
		// technically, binding it directly to the VAR (D -> A) would be sufficient, but the more
		//   ground terms in the bindings means less hops in grounding / dereferencing.
		{
			F:           ast.CreateFact("=", ast.CreateVariable("D"), ast.CreateVariable("A")),
			ShouldMatch: true,
			InitialBindings: resolver.CreateBindings(map[string]ast.Term{
				"A": ast.CreateVariable("B"),
				"B": ast.CreateVariable("C"),
				"C": ast.CreateAtom("a"),
			}),
			ExpectedBindings: []*resolver.Bindings{
				resolver.CreateBindings(map[string]ast.Term{
					"A": ast.CreateVariable("B"),
					"B": ast.CreateVariable("C"),
					"C": ast.CreateAtom("a"),
					"D": ast.CreateAtom("a"),
				}),
			},
		},
		// If you bind a variable that is referenced by other bound vars, they should not be updated in the
		//  bindings directly. Grounding and dereferencing is intended to handle that resolution chain
		{
			F:           ast.CreateFact("=", ast.CreateAtom("a"), ast.CreateVariable("B")),
			ShouldMatch: true,
			InitialBindings: resolver.CreateBindings(map[string]ast.Term{
				"A": ast.CreateVariable("B"),
			}),
			ExpectedBindings: []*resolver.Bindings{
				resolver.CreateBindings(map[string]ast.Term{
					"A": ast.CreateVariable("B"),
					"B": ast.CreateAtom("a"),
				}),
			},
		},
	}

	for _, c := range testCases {
		runUnifyTestCase(c, t)
	}
}

/**
 * TestFactUnification tests various cases around unifying more compelx terms
 */
func TestFactUnification(t *testing.T) {
	testCases := []*testCase{
		// test that facts of different arity cannot unify a/0 vs a/1
		{
			F: ast.CreateFact("=",
				ast.CreateFact("a"),
				ast.CreateFact("a", ast.CreateAtom("a")),
			),
			ShouldMatch:      true,
			InitialBindings:  resolver.CreateBindings(map[string]ast.Term{}),
			ExpectedBindings: []*resolver.Bindings{},
		},
		// test that facts with different args do not unify `a(b)` vs `a(c)`
		{
			F: ast.CreateFact("=",
				ast.CreateFact("a", ast.CreateAtom("b")),
				ast.CreateFact("a", ast.CreateAtom("c")),
			),
			ShouldMatch:      true,
			InitialBindings:  resolver.CreateBindings(map[string]ast.Term{}),
			ExpectedBindings: []*resolver.Bindings{},
		},

		// test that simple facts unify a/0
		{
			F: ast.CreateFact("=",
				ast.CreateFact("a"),
				ast.CreateFact("a"),
			),
			ShouldMatch:     true,
			InitialBindings: resolver.CreateBindings(map[string]ast.Term{}),
			ExpectedBindings: []*resolver.Bindings{
				resolver.CreateBindings(map[string]ast.Term{}),
			},
		},
		// test that simple facts unify a/1
		{
			F: ast.CreateFact("=",
				ast.CreateFact("a", ast.CreateAtom("b")),
				ast.CreateFact("a", ast.CreateAtom("b")),
			),
			ShouldMatch:     true,
			InitialBindings: resolver.CreateBindings(map[string]ast.Term{}),
			ExpectedBindings: []*resolver.Bindings{
				resolver.CreateBindings(map[string]ast.Term{}),
			},
		},
		// TODO: test unifying facts with complex terms
		// TODO: test unifying facts with variables & complex terms
	}

	for _, c := range testCases {
		runUnifyTestCase(c, t)
	}
}
