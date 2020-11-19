package ast

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/kkoch986/gopl/parser/bsr"
	"github.com/kkoch986/gopl/parser/symbols"
	"github.com/kkoch986/gopl/token"
	// "github.com/kkoch986/gopl/parser/slot"
)

// TODO: Error handling across the board
//       seems like you can add extra things that parse but dont get included in the AST

func BuildStatementList(b bsr.BSR) []Statement {
	sl := b.GetNTChildI(0)
	t := sl.Label.Head().String()

	if t == "Statement" {
		f := BuildStatement(sl)
		// TODO: this nil check is not really right since Statement is an interface
		if f != nil {
			return []Statement{f}
		}
	} else if t == "StatementList" {
		l := BuildStatementList(sl)

		ret := []Statement{}
		for _, v := range l {
			ret = append(ret, v)
		}

		// the second part of this might be an arg
		statementBST := b.GetNTChild(symbols.NT_Statement, 0)
		statement := BuildStatement(statementBST)

		if statement == nil {
			panic("Unable to parse statement")
		}

		ret = append(ret, statement)
		return ret
	} else {
		panic("Unknown type found in StatementList: " + t)
	}
	return []Statement{}
}

func BuildStatement(b bsr.BSR) Statement {
	sl := b.GetNTChildI(0)
	t := sl.Label.Head().String()

	if t == "Query" {
		return BuildQuery(sl.GetNTChild(symbols.NT_Concatenation, 0))
	} else if t == "Fact" {
		return BuildFact(sl)
	} else if t == "Rule" {
		return BuildRule(sl)
	} else {
		panic("Unknown Statement Type: " + t)
	}

	return nil
}

func BuildRule(b bsr.BSR) *Rule {
	head := BuildFact(b.GetNTChild(symbols.NT_Fact, 0))
	if head == nil {
		panic("Unable to construct rule head")
	}

	bodybsr := b.GetNTChild(symbols.NT_FactList, 0)
	body := BuildFactList(bodybsr)

	return &Rule{*head, body}
}

func BuildFactList(b bsr.BSR) []Fact {
	sl := b.GetNTChildI(0)
	t := sl.Label.Head().String()

	if t == "Fact" {
		f := BuildFact(sl)
		if f != nil {
			return []Fact{*f}
		}
	} else if t == "FactList" {
		l := BuildFactList(sl)

		ret := []Fact{}
		for _, v := range l {
			ret = append(ret, v)
		}

		// the second part of this might be an arg
		factBST := b.GetNTChild(symbols.NT_Fact, 0)
		fact := BuildFact(factBST)

		if fact == nil {
			panic("Unable to parse fact")
		}

		ret = append(ret, *fact)
		return ret
	} else {
		panic("Unknown type found in FactList: " + t)
	}
	return []Fact{}
}

func BuildQuery(b bsr.BSR) *Query {
	c := b.GetNTChildI(0)
	t := c.Label.Head().String()

	if t == "Fact" {
		f := BuildFact(c)
		if f != nil {
			return &Query{*f}
		}
	} else if t == "Concatenation" {
		cl := BuildQuery(c)

		ret := Query{}
		for _, v := range *cl {
			ret = append(ret, v)
		}

		// the second part of this might be an arg
		factBST := b.GetNTChild(symbols.NT_Fact, 0)
		fact := BuildFact(factBST)

		if fact == nil {
			panic("Unable to parse fact")
		}

		ret = append(ret, *fact)
		return &ret
	} else {
		panic("Unknown type found in Query: " + t)
	}
	return &Query{}
}

func BuildFact(b bsr.BSR) *Fact {
	// TODO: Error handling?
	// The first thing should be either an atom or string lit
	c := b.GetTChildI(0)
	id := c.Type().ID()
	var head string
	if id == "atom" || id == "string_lit" {
		head = string(c.Literal())
	} else {
		panic(fmt.Sprintf("Unknown Fact first parameter %s", id))
	}

	// Handle blank arg lists
	if len(b.Label.Symbols()) <= 2 {
		return &Fact{head, []Arg{}}
	}

	// Next is a NT ArgList
	al := b.GetNTChild(symbols.NT_ArgList, 0)
	argList := BuildArgList(al)

	return &Fact{head, argList}
}

/**
 * BuildList builds a fact from a list parse tree.
 * Lists are just syntactic sugar for a Fact which is provided a special name `|`
 */
func BuildList(b bsr.BSR) *Fact {
	// The first alternative is just an empty list
	// TODO: maybe a better way to do this, i dont like using Alternate because its sensitive
	//       to changes in the order things are written in the grammar
	if b.Alternate() == 0 {
		return buildArgListIntoListFact([]Arg{}, []Arg{})
	}

	al := b.GetNTChildI(1)
	t := al.Label.Head().String()

	// Handle just an arg list
	if t == "ArgList" {
		// parse the arglist into an array of args
		// then, recursively build that into nested |/2 and |/0
		argSlice := BuildArgList(al)

		if len(argSlice) == 0 {
			panic("Unexpected empty arglist in List parse")
		}

		return buildArgListIntoListFact(argSlice, []Arg{})
	} else if t == "Cons" {
		// Build the LHS arg list
		lhs := BuildArgList(al.GetNTChildI(0))
		rhs := BuildArgList(al.GetNTChildI(2))

		return buildArgListIntoListFact(lhs, rhs)
	} else {
		panic("Unknown list type: " + t)
	}
	return nil
}

/**
 * take an array of Args and construct them into nested |/2 and |/0
 * for example:
 *  1,2,3 becomes `|(1, |(2, |(3, |())))`
 */
func buildArgListIntoListFact(lhs []Arg, rhs []Arg) *Fact {
	if len(lhs) == 0 {
		if len(rhs) == 0 {
			return &Fact{"|", []Arg{}}
		} else {
			return buildArgListIntoListFact(rhs, []Arg{})
		}
	} else if len(lhs) == 1 && len(rhs) == 1 {
		// If the left hand side only has one element and theres a RHS with one element
		// fill the |/2 with each one
		// think of the case like `[1,2|X]`
		// the outcome should be `|(1, |(2, X))`
		// If there are more than one things on the RHS, we will resolve that as a list itself
		// above with a call to buildArgListIntoListFact
		return &Fact{"|", []Arg{lhs[0], rhs[0]}}
	}

	// Build the tail
	tail := buildArgListIntoListFact(lhs[1:], rhs)
	if tail == nil {
		panic("error building list tail")
	}

	return &Fact{"|", []Arg{lhs[0], tail}}
}

func BuildArgList(b bsr.BSR) []Arg {
	al := b.GetNTChildI(0)
	t := al.Label.Head().String()

	if t == "Arg" {
		a := BuildArg(al)
		if a != nil {
			return []Arg{a}
		}
	} else if t == "ArgList" {
		argList := BuildArgList(al)

		ret := []Arg{}
		for _, v := range argList {
			ret = append(ret, v)
		}

		// the second part of this might be an arg
		argBST := b.GetNTChild(symbols.NT_Arg, 0)
		arg := BuildArg(argBST)

		if arg == nil {
			panic("Unable to parse arg")
		}

		return append(ret, arg)
	} else {
		panic("Unknown ArgList type " + t)
	}
	return []Arg{}
}

func BuildArg(b bsr.BSR) Arg {
	t := b.Label.Symbols()[0].String()

	var ret Arg
	switch t {
	case "atom":
		ret = BuildAtom(b.GetTChildI(0))
	case "string_lit":
		ret = BuildStringLiteral(b.GetTChildI(0))
	case "num_lit":
		ret = BuildNumericLiteral(b.GetTChildI(0))
	case "var":
		ret = BuildVariable(b.GetTChildI(0))
	case "Fact":
		ret = BuildFact(b.GetNTChild(symbols.NT_Fact, 0))
	case "List":
		ret = BuildList(b.GetNTChild(symbols.NT_List, 0))
	default:
		panic("Unknown Arg type: " + t)
	}

	if reflect.ValueOf(ret).IsNil() {
		return nil
	}

	return ret
}

func BuildVariable(t *token.Token) *Variable {
	return &Variable{string(t.Literal())}
}

func BuildAtom(t *token.Token) *Atom {
	return &Atom{string(t.Literal())}
}

func BuildStringLiteral(t *token.Token) *StringLiteral {
	return &StringLiteral{string(t.Literal())}
}

func BuildNumericLiteral(t *token.Token) *NumericLiteral {
	s := string(t.Literal())
	n, _ := strconv.ParseFloat(s, 64)
	return &NumericLiteral{n}
}
