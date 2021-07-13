package resolver_test

import "testing"
import "github.com/kkoch986/gopl/ast"
import "github.com/kkoch986/gopl/resolver"

/**
 * TestEmpty will test that the empty function returns true when there is nothing bound
 * and false otherwise
 */
func TestEmpty(t *testing.T) {
	// a new binding should be empty
	b := resolver.EmptyBindings()
	if !b.Empty() {
		t.Errorf("New binding is not empty")
	}

	// set a binding
	b.Bind("a", &ast.Fact{Head: "a", Args: []ast.Term{}})

	// not empty anymore
	if b.Empty() {
		t.Errorf("Binding still empty after set")
	}
}

/**
 * TestDeref will test various cases of bindings and dereferences
 */
func TestDeref(t *testing.T) {
	b := resolver.EmptyBindings()

	// dereferencing something that is not a variable should just return that thing
	var bases = []ast.Term{
		ast.CreateAtom("a"),
		ast.CreateStringLiteral("b"),
		ast.CreateNumericLiteral(10),
		&ast.Fact{Head: "b", Args: []ast.Term{}},
	}
	for _, base := range bases {
		v := b.Dereference(base)
		if v != base {
			t.Errorf("non-variable deref failed, expected %s, got %s", base, v)
		}
	}

	// bind the variable "A" to the atom "a"
	// ensure that it dereferences correctly, then make sure atoms and strings "A" do not derefernce to atom "a"
	aBinding := ast.CreateAtom("a")
	b.Bind("A", aBinding)

	v := b.Dereference(ast.CreateVariable("A"))
	if v != aBinding {
		t.Errorf("Dereferencing variable failed. expected %s, got %s", aBinding, v)
	}

	bases = []ast.Term{
		ast.CreateAtom("A"),
		ast.CreateStringLiteral("A"),
	}
	for _, base := range bases {
		v := b.Dereference(base)
		if v != base {
			t.Errorf("Non-variable deferenced as bound variable. expected %s, got %s", base, v)
		}
	}

	// test a chain of variable bindings
	finalBinding := ast.CreateAtom("final")
	b.Bind("C", ast.CreateVariable("D"))
	b.Bind("D", ast.CreateVariable("E"))
	b.Bind("E", finalBinding)

	bases = []ast.Term{
		ast.CreateVariable("C"),
		ast.CreateVariable("D"),
		ast.CreateVariable("E"),
		finalBinding,
	}
	for _, base := range bases {
		v := b.Dereference(base)
		if v != finalBinding {
			t.Errorf("variable chain resolution failed for %s. Expected %s, got %s", base, finalBinding, v)
		}
	}
}

/**
 * TestClone will test that a cloned binding does not affect the original object
 */
func TestClone(t *testing.T) {
	b := resolver.EmptyBindings()

	// bind some variables
	b.Bind("a", ast.CreateVariable("A"))

	v := b.Dereference(ast.CreateVariable("a"))
	if v.(*ast.Variable).String() != "A" {
		t.Errorf("initial binding failed deref. Expected %s, got %s", "A", v)
	}

	// clone it
	c := b.Clone()

	// assert that dereferencing the things set return the same results
	v = c.Dereference(ast.CreateVariable("a"))
	if v.(*ast.Variable).String() != "A" {
		t.Errorf("cloned binding failed deref. Expected %s, got %s", "A", v)
	}

	// change something on b and assert its not changed in c
	b.Bind("b", ast.CreateVariable("B"))

	// check its good in b
	v = b.Dereference(ast.CreateVariable("b"))
	if v.(*ast.Variable).String() != "B" {
		t.Errorf("initial binding failed deref. Expected %s, got %s", "B", v)
	}

	// and not bound in c
	v = c.Dereference(ast.CreateVariable("b"))
	if v.(*ast.Variable).String() != "b" {
		t.Errorf("cloned binding failed deref. Expected %s, got %s", "b", v)
	}

	// change something on c and assert its not changed in b
	c.Bind("c", ast.CreateVariable("C"))

	// check its unbound in b
	v = b.Dereference(ast.CreateVariable("c"))
	if v.(*ast.Variable).String() != "c" {
		t.Errorf("initial binding failed deref. Expected %s, got %s", "c", v)
	}

	// and bound in c
	v = c.Dereference(ast.CreateVariable("c"))
	if v.(*ast.Variable).String() != "C" {
		t.Errorf("cloned binding failed deref. Expected %s, got %s", "C", v)
	}
}

/**
 * TestEquals will test various cases of equality testing between 2 bindings
 **/
func TestEquals(t *testing.T) {
	// Test empty == empty
	if !resolver.EmptyBindings().Equals(resolver.EmptyBindings()) {
		t.Errorf("Empty bindings !== empty bindings")
	}

	// Test 2 equal bindings
	b1 := resolver.CreateBindings(map[string]ast.Term{
		"A": ast.CreateAtom("a"),
		"B": ast.CreateAtom("b"),
	})
	b2 := resolver.CreateBindings(map[string]ast.Term{
		"A": ast.CreateAtom("a"),
		"B": ast.CreateAtom("b"),
	})
	if !b1.Equals(b2) {
		t.Errorf("failed asserting that \n%v equals \n%v", b1, b2)
	}

	// Test 2 not-equal bindings with the same values but different keys
	b1 = resolver.CreateBindings(map[string]ast.Term{
		"A": ast.CreateAtom("a"),
	})
	b2 = resolver.CreateBindings(map[string]ast.Term{
		"B": ast.CreateAtom("a"),
	})
	if b1.Equals(b2) {
		t.Errorf("incorrectly asserted that \n%v equals \n%v", b1, b2)
	}

	// Test 2 not-equal bindings with the same number keys but different values
	b1 = resolver.CreateBindings(map[string]ast.Term{
		"A": ast.CreateAtom("a"),
	})
	b2 = resolver.CreateBindings(map[string]ast.Term{
		"A": ast.CreateAtom("b"),
	})
	if b1.Equals(b2) {
		t.Errorf("incorrectly asserted that\n %v equals \n%v", b1, b2)
	}

	// TODO: test complex types
}
