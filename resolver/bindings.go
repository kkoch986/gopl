package resolver

import (
	"fmt"

	"github.com/kkoch986/gopl/ast"
)

type Bindings struct {
	B map[string]ast.Term
}

func CreateBindings() *Bindings {
	b := Bindings{}
	b.B = make(map[string]ast.Term)
	return &b
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
	fmt.Printf("[BIND] %s -> %s", k, v)
	if b.B[k] != nil {
		fmt.Printf("     FAIL (already bound to %s)\n", b.B[k])
		return false
	}
	fmt.Println("    SUCCESS")
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
