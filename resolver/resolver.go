package resolver

import (
	"log"

	"github.com/kkoch986/gopl/ast"
)

type Bindings struct {
	B map[string]interface{}
}

type FactResolver interface {
	Resolve(*ast.Fact, *Bindings, chan<- *Bindings, chan<- bool)
}

type R struct {
	fr []FactResolver
}

func (r *R) AddFactResolver(nr FactResolver) {
	r.fr = append(r.fr, nr)
}

func New() *R {
	r := &R{}
	r.AddFactResolver(&Writeln{})
	r.AddFactResolver(&True{})
	r.AddFactResolver(&Fail{})
    r.AddFactResolver(&Load{})
	return r
}

func (r *R) ResolveStatementList(sl []ast.Statement, c *Bindings, out chan<- *Bindings) error {
	defer close(out)
	if len(sl) == 0 {
		out <- c
		return nil
	}

	// find all the bindings for the first statement
	headBindings := make(chan *Bindings, 2)
    tail := sl[1:]

	go r.ResolveStatement(sl[0], c, headBindings)
	for hb := range headBindings {
		// for each binding of the first element of the list, try to resolve the next
        tailBindings := make(chan *Bindings, 2)
        go r.ResolveStatementList(tail, hb, tailBindings)
        for ob := range tailBindings {
            out <- ob
        }
	}
	return nil
}

func (r *R) ResolveStatement(s ast.Statement, c *Bindings, out chan<- *Bindings) {
	log.Printf("[ResolveStatement] %s", s)

	t := s.GetType()
	switch t {
	case ast.T_Query:
		go r.ResolveQuery(s.(*ast.Query), c, out)
	case ast.T_Rule:
	case ast.T_Fact:
	default:
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
	headBindings := make(chan *Bindings, 2)
	tail := q.Tail()
	go r.ResolveFact(q.Head(), c, headBindings)

	for b := range headBindings {
		// find all resolutions of the tail and run them back to out
		tailBindings := make(chan *Bindings, 2)
		go r.ResolveQuery(tail, b, tailBindings)
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

	// TODO: if nothing matched, invoke the default behavior
}
