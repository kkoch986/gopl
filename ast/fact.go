package ast

import (
	"fmt"
	"strings"
)

type Fact struct {
	Head string
	Args []Arg
}

func (f *Fact) GetType() string {
	return "Fact"
}

func (f *Fact) String() string {
	args := []string{}

	for _, v := range f.Args {
		args = append(args, v.String())
	}

	// TODO: handling pretty-printing lists (any fact with head == "|")
	if f.Head == "|" {
		if len(f.Args) == 0 {
			return "L[]"
		}
		return "L[" + prettyPrintList(f.Args) + "]"
	}

	return fmt.Sprintf("%s(%s)", f.Head, strings.Join(args, ","))
}

func prettyPrintList(a []Arg) string {
	if len(a) == 0 {
		return ""
	}

	left := a[0]
	right := a[1]

	if right.GetType() == "Fact" && right.(*Fact).Head == "|" {
		rightStr := prettyPrintList(right.(*Fact).Args)
		if rightStr == "" {
			return fmt.Sprintf("%s", left)
		}
		return fmt.Sprintf("%s,%s", left, rightStr)
	} else {
		return fmt.Sprintf("%s|%s", left, right)
	}
}
