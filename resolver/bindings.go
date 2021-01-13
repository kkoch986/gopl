package resolver

import (
	"fmt"
	"log"

	"github.com/kkoch986/gopl/ast"
)

type Bindings struct {
	B map[string]ast.Term
}

func EmptyBindings() *Bindings {
	b := Bindings{}
	b.B = make(map[string]ast.Term)
	return &b
}

func CreateBindings(m map[string]ast.Term) *Bindings {
	return &Bindings{m}
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
	log.Printf("[BIND] %s -> %s", k, v)
	if b.B[k] != nil && b.B[k] != v {
		log.Printf("[BIND] %s -> %s   FAIL (already bound to %s)\n", k, v, b.B[k])
		return false
	}
	log.Printf("[BIND] %s -> %s     SUCCESS", k, v)
	b.B[k] = v
	return true
}

func (b *Bindings) String() string {
	ret := "Bindings: \n"
	for k, v := range b.B {
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
