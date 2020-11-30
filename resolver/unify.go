package resolver

import (
	"github.com/kkoch986/gopl/ast"
)

/**
 * attempt to unify the 2 given facts
 * 2 facts unify iff:
 *   1. their Head and arity are the same
 *   2. each of their args can unify
 * If a nil binding is returned, the two terms do not unify.
 * Otherwise, the new binding returned is the resultant binding.
 */
func unifyFacts(base *ast.Fact, query *ast.Fact, b *Bindings) *Bindings {
	// check that the signatures are the same
	if base.Signature().String() != query.Signature().String() {
		return nil
	}

	// check that each of the args unify, passing the binding from the first arg onto the next one etc..
	testBindings := b.Clone()
	for i, b := range base.Args {
		q := query.Args[i]
		testBindings = unifyTerms(b, q, testBindings)

		if testBindings == nil {
			return nil
		}
	}

	return testBindings
}

/**
 * Attempt to unify any 2 terms
 * TODO: flesh out this comment some more
 */
func unifyTerms(base ast.Term, query ast.Term, b *Bindings) *Bindings {
	baseType := base.GetType()
	queryType := query.GetType()
    
    

	// if either is a variable, derefence it first
    base = b.Dereference(base)
    query = b.Dereference(query)

    // UNIFY TERMINALS
	// The most basic case is if both are primitive terms (string, atom, number)
	// Basically are they the same value?
	if (baseType == ast.T_String || baseType == ast.T_Atom) && (queryType == ast.T_String || queryType == ast.T_Atom) {
		if base.String() == query.String() {
			return b
		}
		return nil
	}

	if baseType == ast.T_Number && queryType == baseType {
		if base.(*ast.NumericLiteral).String() == query.(*ast.NumericLiteral).String() {
			return b
		}
		return nil
	}

    // VARIABLES
    // If both are variables, pick one and assign it to the other
    //   this needs to be done deterministically so for now just go with alphabetical order by the variable name
    // If only one is a variable, attempt to bind that variable to the other
    if baseType == ast.T_Variable && queryType == ast.T_Variable {
        // if both terms represent the same variable, do not do any binding
        // but _do_ return the current bindings since they do technically unify
        if(base.String() == query.String()) {
            return b
        } else if base.String() < query.String() {
            test := b.Clone()
            if test.Bind(base.String(), query) {
                return test
            }
            return nil
        } else {
            test := b.Clone()
            if test.Bind(query.String(), base) {
                return test
            }
            return nil
        }
    } else if baseType == ast.T_Variable {
        // create a copy of the bindings so we can test things out and return it if its ok
        test := b.Clone()
        if test.Bind(base.String(), query) {
            return test
        }
        return nil
    } else if queryType == ast.T_Variable {
        // create a copy of the bindings so we can test things out and return it if its ok
        test := b.Clone()
        if test.Bind(query.String(), base) {
            return test
        }
        return nil
    }

    // UNIFY FACTS
    if (baseType == ast.T_Fact && queryType == ast.T_Fact) {
        return unifyFacts(base.(*ast.Fact), query.(*ast.Fact), b)
    }

    // If we fall all the way through, assume there are no bindings
	return nil
}
