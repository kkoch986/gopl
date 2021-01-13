package resolver

import "testing"
import "github.com/kkoch986/gopl/ast"

/**
 * TestFailPassthrough ensures that this does not respond to anything but fail/0
 */
func TestFailPassthrough(t *testing.T) {
	f := &Fail{}
	m := make(chan bool)
	out := make(chan *Bindings)

	go f.Resolve(&ast.Fact{Head: "something", Args: []ast.Term{}}, EmptyBindings(), out, m)

	select {
	case o := <-out:
		t.Errorf("Fail resolver wrote results for something/0 (got %s)", o)
	case matched := <-m:
		if matched {
			t.Errorf("Fail resolver matched for something/0")
		}
	}

	// make sure it doesnt match "fail/1"
	m = make(chan bool)
	out = make(chan *Bindings)
	go f.Resolve(&ast.Fact{Head: "fail", Args: []ast.Term{ast.CreateAtom("A")}}, EmptyBindings(), out, m)
	select {
	case o := <-out:
		t.Errorf("Fail resolver wrote results for fail/1 (got %s)", o)
	case matched := <-m:
		if matched {
			t.Errorf("Fail resolver matched for fail/1")
		}
	}

}

/**
 * TestNoResults affirms that fail matches fail/0 but does not return any results
 */
func TestFailNoResults(t *testing.T) {
	f := &Fail{}
	m := make(chan bool)
	out := make(chan *Bindings)

	go f.Resolve(&ast.Fact{Head: "fail", Args: []ast.Term{}}, EmptyBindings(), out, m)

	select {
	case o := <-out:
		t.Errorf("Fail resolver wrote results for fail/0 (got %s)", o)
	case matched := <-m:
		if !matched {
			t.Errorf("Fail resolver did not match for fail/0")
		}
	}
}
