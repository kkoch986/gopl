package resolver

import (
	"fmt"
	"log"

	"github.com/kkoch986/gopl/ast"
	"github.com/kkoch986/gopl/indexer"
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
		fmt.Println("how to resolve?", t)
		close(out)
	}
}

func (r *R) ResolveQuery(q *ast.Query, c *Bindings, out chan<- *Bindings) {
	defer close(out)
	log.Printf("[ResolveQuery] %s", q)

	// If there are no statements in the list, accept the current binding
	if q.Empty() {
		out <- c
		return
	}

	// A query is an array of facts, recursively loop over each to DFS all possible bindings
	headBindings := make(chan *Bindings, 1)
	tail := q.Tail()
	go r.ResolveFact(q.Head(), c, headBindings)

	for hb := range headBindings {
		// find all resolutions of the tail and run them back to out
		tailBindings := make(chan *Bindings, 1)
		go r.ResolveQuery(tail, hb, tailBindings)
		for ob := range tailBindings {
			out <- ob
		}
	}
}

func (r *R) ResolveFact(f *ast.Fact, c *Bindings, out chan<- *Bindings) {
	defer close(out)
	log.Printf("[ResolveFact] %s\n", f)

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
			//    1. create an initial "stack frame" by unifying the query with the head of the base rule
			//    2. With that binding, resolve the body of the rule as if it were a query.
			//    3. With each resulting binding:
			//       - extract the variables from the head of the query fact
			//       - return a new binding which binds each variable from the query fact to its corresponding
			//         value in the "stack frame" binding.
			// not sure this comment is enlightening but hopefully it and the code make sense together...
			initialBinding := unifyFacts(rule.Head, f, c)

			if initialBinding == nil {
				continue
			}

			// set up a channel to receive valid resolutions of the body of the rule
			discoveredBindings := make(chan *Bindings, 1)
			q := ast.Query(rule.Body)
			go r.ResolveStatementList([]ast.Statement{&q}, initialBinding, discoveredBindings)
			for db := range discoveredBindings {
				// find all of the variables defined by the head of the rule
				// lookup their values in the discovered binding
				// if it is bound to a non-variable, add the same binding to a clone of `c` and return that
				outBinding := c.Clone()
				valid := true
				for _, variable := range f.ExtractVariables() {
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
