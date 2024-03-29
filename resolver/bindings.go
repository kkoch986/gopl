package resolver

import (
	"fmt"
	"log"
	"sort"
	"strings"

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

func (b *Bindings) Equals(c *Bindings) bool {
	// quick length check
	if len(b.B) != len(c.B) {
		return false
	}

	// go key by key from b and make sure it exists and matches
	for key, bv := range b.B {
		cv := c.B[key]
		if cv == nil {
			return false
		}

		if bv.GetType() != cv.GetType() {
			return false
		}

		// TODO: theres probably a better way to do this...
		if bv.String() != cv.String() {
			return false
		}
	}

	return true
}

func (b *Bindings) Clone() *Bindings {
	newMap := make(map[string]ast.Term)
	for i, v := range b.B {
		newMap[i] = v
	}
	return &Bindings{newMap}
}

func (b *Bindings) Bind(k string, v ast.Term) bool {
	// TODO: when binding, if the variable is already bound, but its bound to a variable we need to handle that differently...
	target := b.B[k]
	if target != nil {
		target = b.Dereference(b.B[k])
	}
	if target != nil && target.GetType() != ast.T_Variable {
		if b.B[k] != nil && b.B[k] != v {
			log.Printf("[DEBUG][BIND] %s -> %s   FAIL (already bound to %s [%s])\n", k, v, b.B[k], b.Dereference(b.B[k]))
			return false
		}
	} else if target != nil {
		// bind v to a variable bound to another variable
		return b.Bind(target.String(), v)
	}
	log.Printf("[VERBOSE][BIND] %s -> %s     SUCCESS", k, v)
	b.B[k] = v
	return true
}

func (b *Bindings) String() string {
	ret := "Bindings: \n"
	s := make([]string, len(b.B))
	for k, v := range b.B {
		s = append(s, fmt.Sprintf("\t%s: %s\n", k, b.Ground(v)))
	}
	sort.Strings(s)
	return ret + strings.Join(s, "")
}

func (b *Bindings) ShortString() string {
	ret := ""
	s := make([]string, len(b.B))
	i := 0
	for k, v := range b.B {
		s[i] = fmt.Sprintf("%s:%s", k, b.Ground(v))
		i = i + 1
	}
	sort.Strings(s)
	ret = ret + strings.Join(s, ", ")
	return ret
}

/**
 * Ground is similar to dereference, but it will go deep into nested facts and dereference
 * all terms it can find.
 */
func (b *Bindings) Ground(t ast.Term) ast.Term {
	termType := t.GetType()
	switch termType {
	case ast.T_Fact:
		f := t.(*ast.Fact)
		argc := len(f.Args)
		// in case of a fact, loop over each of the args and ground them
		newArgs := make([]ast.Term, argc)
		for i, v := range f.Args {
			newArgs[i] = b.Ground(v)
		}
		return &ast.Fact{Head: f.Head, Args: newArgs}
	case ast.T_Variable:
		// when grouding a variable, dont change it if the deref returns a variable.
		// this prevents accidentally swapping for the wrong variable when resolving facts with rules
		d := b.Dereference(t)
		if d.GetType() == ast.T_Variable {
			return t
		}
		return d
	case ast.T_Atom:
		fallthrough
	case ast.T_String:
		fallthrough
	case ast.T_Number:
		fallthrough
	default:
		return b.Dereference(t)
	}
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
