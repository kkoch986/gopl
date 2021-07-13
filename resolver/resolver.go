package resolver

import (
	"errors"
	"fmt"
	"log"

	"github.com/kkoch986/gopl/ast"
	"github.com/kkoch986/gopl/indexer"
)

var (
	ErrUnboundVariable    = errors.New("Unbound variable in MathExpr")
	ErrNonNumericVariable = errors.New("Non-numeric variable assignment in MathExpr")
	ErrUnknownMathExprOp  = errors.New("Unknown MathExpr Operation")
	ErrUnknownMultOp      = errors.New("Unknown Mult Operation")
	ErrInvalidFactor      = errors.New("MathExpr Factor found with no value assigned")
)

type FactResolver interface {
	Resolve(*ast.Fact, *Bindings, chan<- *Bindings, chan<- bool)
}

type R struct {
	fr []FactResolver
	i  indexer.Indexer
}

func (r *R) AddFactResolver(nr FactResolver) {
	r.fr = append(r.fr, nr)
}

func (r *R) AddFactResolvers(rs []FactResolver) {
	for _, v := range rs {
		r.AddFactResolver(v)
	}
}

func New(i indexer.Indexer) *R {
	r := &R{
		i: i,
	}
	r.AddFactResolvers([]FactResolver{
		&Equals{},
		&Writeln{},
		&True{},
		&Fail{},
		&Assert{i},
	})
	return r
}

func (r *R) ResolveStatementList(sl []ast.Statement, c *Bindings, out chan<- *Bindings) {
	defer close(out)
	if len(sl) == 0 {
		out <- c
		return
	}

	// find all the bindings for the first statement
	headBindings := make(chan *Bindings, 1)
	tail := sl[1:]

	go r.ResolveStatement(sl[0], c, headBindings)
	for hb := range headBindings {
		// for each binding of the first element of the list, try to resolve the next
		tailBindings := make(chan *Bindings, 1)
		go r.ResolveStatementList(tail, hb, tailBindings)
		for ob := range tailBindings {
			out <- ob
		}
	}
}

func (r *R) ResolveStatement(s ast.Statement, c *Bindings, out chan<- *Bindings) {
	t := s.GetType()
	log.Printf("[ResolveStatement] %s (%s)", s, t)

	switch t {
	case ast.T_Query:
		go r.ResolveQuery(s.(*ast.Query), c, out)
	case ast.T_Rule:
		fallthrough
	case ast.T_Fact:
		fallthrough
	default:
		log.Printf("[WARN] Unknown resolution input type: %v", s)
		close(out)
	}
}

func (r *R) ResolveQuery(q *ast.Query, c *Bindings, out chan<- *Bindings) {
	defer close(out)
	log.Printf("[DEBUG][ResolveQuery] %s", q)

	// If there are no statements in the list, accept the current binding
	if q.Empty() {
		out <- c
		return
	}

	// A query is an array of facts, recursively loop over each to DFS all possible bindings
	headBindings := make(chan *Bindings, 1)
	tail := q.Tail()
	headType := q.Head().GetType()
	if headType == ast.T_Fact {
		go r.ResolveFact(q.Head().(*ast.Fact), c, headBindings)
	} else if headType == ast.T_MathAssignment {
		go r.ResolveMathAssignment(q.Head().(*ast.MathAssignment), c, headBindings)
	} else {
		// should really never get here...
		panic(fmt.Sprintf("Can't resolve query item (not a fact or math assignment): %s", headType))
	}

	for hb := range headBindings {
		// find all resolutions of the tail and run them back to out
		tailBindings := make(chan *Bindings, 1)
		go r.ResolveQuery(tail, hb, tailBindings)
		for ob := range tailBindings {
			out <- ob
		}
	}
}

func (r *R) ResolveMathAssignment(ma *ast.MathAssignment, c *Bindings, out chan<- *Bindings) {
	defer close(out)
	log.Printf("[ResolveMathAssignment] %s", ma)

	// check the variable on the LHS
	//  if its bound, we should fail
	if ma.LHS != c.Ground(ma.LHS) {
		return
	}

	// Now try to derive a numeric value for the RHS
	// Resolving a math expression cannot create new bindings, but i can contain variables
	//  so it requires the current bindings to be able to resolve it
	//  if any unbound variables are encountered, the resolution will fail
	val, err := r.ResolveMathExpr(ma.RHS, c)

	// if there were issues resolving, just fail
	if err != nil {
		log.Printf("[ResolveMathExpr] Failed; %s", err)
		return
	}

	// if there were no errors, bind the numeric value to the variable in the LHS
	output := c.Clone()
	output.Bind(ma.LHS.String(), ast.CreateNumericLiteral(val))
	out <- output
}

func (r *R) ResolveMathExpr(me *ast.MathExpr, c *Bindings) (float64, error) {
	// resolve the LHS no matter what
	lhs, err := r.ResolveMult(me.LHS, c)
	if err != nil {
		return 0, err
	}

	// If the op is no-op, just return the lhs value
	op := me.Operator
	if op == ast.OP_MathExprNoOp {
		return lhs, nil
	}

	// resolve the RHS
	rhs, err := r.ResolveMult(me.RHS, c)
	if err != nil {
		return 0, err
	}

	switch op {
	case ast.OP_Add:
		return (lhs + rhs), nil
	case ast.OP_Subtract:
		return (lhs - rhs), nil
	}
	return 0, ErrUnknownMathExprOp
}

func (r *R) ResolveMult(m *ast.Mult, c *Bindings) (float64, error) {
	// resolve the LHS
	lhs, err := r.ResolveFactor(m.LHS, c)
	if err != nil {
		return 0, err
	}

	// if the op is no op, return lhs
	op := m.Operator
	if op == ast.OP_MultNoOp {
		return lhs, nil
	}

	// resolve the rhs
	rhs, err := r.ResolveFactor(m.RHS, c)
	if err != nil {
		return 0, nil
	}

	switch op {
	case ast.OP_Mult:
		return (lhs * rhs), nil
	case ast.OP_Divide:
		return (lhs / rhs), nil
	}
	return 0, ErrUnknownMultOp
}

func (r *R) ResolveFactor(f *ast.Factor, c *Bindings) (float64, error) {
	if f.Num != nil {
		return f.Num.Value(), nil
	} else if f.Var != nil {
		// try to dereference the variable, if its not bound or not bound to a number, return an error
		v := c.Dereference(f.Var)
		t := v.GetType()

		if t == ast.T_Variable {
			return 0, ErrUnboundVariable
		} else if t != ast.T_Number {
			return 0, ErrNonNumericVariable
		} else {
			return v.(*ast.NumericLiteral).Value(), nil
		}
	} else if f.Expr != nil {
		return r.ResolveMathExpr(f.Expr, c)
	}
	return 0, ErrInvalidFactor
}

func (r *R) ResolveFact(f *ast.Fact, c *Bindings, out chan<- *Bindings) {
	defer close(out)
	log.Printf("[ResolveFact] %s (as %s)\n", c.Ground(f), f)

	// loop over all the resolvers one at a time until one matches (indicated by closing `out`)
	rChan := make(chan *Bindings, 1)
	mChan := make(chan bool, 1)
	for _, resolver := range r.fr {
		go resolver.Resolve(f, c, rChan, mChan)
	ResultLoop:
		for {
			select {
			case b, ok := <-rChan:
				if !ok {
					return
				}
				out <- b
			case m := <-mChan:
				if m {
					return
				}
				break ResultLoop
			}
		}
	}

	// If we didnt find a matching resolver, follow the default behavior
	// Find all of the statements that match the signature
	matching := r.i.StatementsForSignature(f.Signature())
	// attempt to unify the input fact with each of the matching statements
	// return each one that does unify as a result binding
	for _, s := range matching {
		t := s.GetType()
		if t == ast.T_Fact {
			newBinding := unifyFacts(s.(*ast.Fact), f, c)
			if newBinding != nil {
				out <- newBinding
			}
		} else if t == ast.T_Rule {
			rule := s.(*ast.Rule)
			// We are trying to unify a Fact (the query) and a Rule (the base)
			// To unify a fact with a rule, follow this procedure:
			//    1. Ground the fact with respect to the current bindings
			//    1. create an initial "stack frame" by
			//        unifying the grounded query with the head of the base rule using fresh bindings
			//    2. With that binding, resolve the body of the rule as if it were a query.
			//    3. With each resulting binding:
			//       - extract the variables from the head of the query fact
			//       - return a new binding which binds each variable from the query fact to its corresponding
			//         value in the "stack frame" binding.
			// not sure this comment is enlightening but hopefully it and the code make sense together...
			initialBinding := unifyFacts(rule.Head, c.Ground(f).(*ast.Fact), EmptyBindings())

			if initialBinding == nil {
				continue
			}

			variablesToProve := initialBinding.Ground(f).(*ast.Fact).ExtractVariables()

			// set up a channel to receive valid resolutions of the body of the rule
			discoveredBindings := make(chan *Bindings, 1)
			go r.ResolveStatementList([]ast.Statement{rule.Body}, initialBinding, discoveredBindings)
			for db := range discoveredBindings {
				// find all of the variables defined by the head of the rule
				// lookup their values in the discovered binding
				// if it is bound to a non-variable, add the same binding to a clone of `c` and return that
				outBinding := c.Clone()
				valid := true
				for _, variable := range variablesToProve {
					deref := db.Dereference(variable)
					if deref == nil {
						continue
					}

					derefType := deref.GetType()
					if derefType != ast.T_Variable {
						if !outBinding.Bind(variable.String(), deref) {
							valid = false
							continue
						}
					}
				}
				if valid {
					out <- outBinding
				}
			}
		} else {
			fmt.Println("TODO: other types of fact unification")
		}
	}
}
