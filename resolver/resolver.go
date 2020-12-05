package resolver

import (
	"fmt"
	"log"

	"github.com/kkoch986/gopl/ast"
	"github.com/kkoch986/gopl/indexer"
)

// TODO: stop using interface{} in bindings
type Bindings struct {
	B map[string]ast.Term
}

func (b *Bindings) Empty() bool {
    return len(b.B) == 0
}

func (b *Bindings) Clone() *Bindings {
	newMap := make(map[string]ast.Term)
	for i, v := range b.B {
		newMap[i] = v
	}
	return &Bindings{newMap}
}

func (b *Bindings) Bind(k string, v ast.Term) bool {
    if b.B[k] != nil {
        return false
    }
    b.B[k] = v
    return true
}

func (b *Bindings) String() string {
    ret := "Bindings: \n"
    for k, v := range(b.B) {
        ret = ret + fmt.Sprintf("\t%s: %s\n", k, v)
    }
    return ret
}

/**
 * Derefernce takes a term and returns a term.
 * If the term is a variable, and there is a binding present, it will return that term
 */
func (b *Bindings) Dereference(t ast.Term) ast.Term {
    termType := t.GetType()
    if termType == ast.T_Variable {
        d := b.B[t.(*ast.Variable).String()]
        // if there is no binding for this term, just return it as is
        if d == nil {
            return t
        }
        // if what was returned is a variable, dereference that again
        if d.GetType() == ast.T_Variable {
            return b.Dereference(d)
        }

        // otherwise, return whatever we got
        return d
    }
    return t
}

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
		&Writeln{r},
		&True{},
		&Fail{},
		&Load{},
		&Assert{i},
	})
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

	for hb := range headBindings {
		// find all resolutions of the tail and run them back to out
		tailBindings := make(chan *Bindings, 2)
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

	// TODO: if nothing matched, invoke the default behavior
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
		} else {
			fmt.Println("TODO: other types of fact unification")
		}
	}
}
