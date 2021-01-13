package resolver

import "testing"
import "github.com/kkoch986/gopl/ast"

/**
 * TestTruePassthrough ensures that this does not respond to anything but true/0
 */
func TestTruePassthrough(t *testing.T) {
	f := &True{}
	m := make(chan bool)
	out := make(chan *Bindings)

	go f.Resolve(&ast.Fact{Head: "something", Args: []ast.Term{}}, EmptyBindings(), out, m)

	select {
	case o := <-out:
		t.Errorf("True resolver wrote results for something/0 (got %s)", o)
	case matched := <-m:
		if matched {
			t.Errorf("True resolver matched for something/0")
		}
	}

	// make sure it doesnt match "fail/1"
	m = make(chan bool)
	out = make(chan *Bindings)
	go f.Resolve(&ast.Fact{Head: "true", Args: []ast.Term{ast.CreateAtom("A")}}, EmptyBindings(), out, m)
	select {
	case o := <-out:
		t.Errorf("True resolver wrote results for true/1 (got %s)", o)
	case matched := <-m:
		if matched {
			t.Errorf("True resolver matched for true/1")
		}
	}

}

/**
 * TestTrueResults affirms that true matches true/0 and returns the given bindings
 */
func TestTrueResults(t *testing.T) {
	f := &True{}
	m := make(chan bool)
	out := make(chan *Bindings)
	inputBindings := EmptyBindings()

	go f.Resolve(&ast.Fact{Head: "true", Args: []ast.Term{}}, inputBindings, out, m)

	foundResults := false
	select {
	case o := <-out:
		foundResults = true
		if o != inputBindings {
			t.Errorf("True resolver wrote incorrect results for true/0 (expected %s, got %s)", inputBindings, o)
		}
	case matched := <-m:
		if !matched {
			t.Errorf("True resolver did not match for true/0")
		}
	}

	if !foundResults {
		t.Errorf("True resolver did not write results for true/0")
	}
}
