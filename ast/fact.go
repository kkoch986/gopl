package ast

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Signature struct {
	Functor string
	Arity   int
}

func (s *Signature) String() string {
	return fmt.Sprintf("%s/%d", s.Functor, s.Arity)
}

type Fact struct {
	Head string `json:"f"`
	Args []Term `json:"a"`
}

func (f *Fact) GetType() TermType {
	return T_Fact
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

func (f *Fact) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{})
	m["t"] = "fact"
	m["f"] = f.Head
	m["a"] = f.Args
	return json.Marshal(m)
}

func (f *Fact) Signature() *Signature {
	return &Signature{f.Head, len(f.Args)}
}

func prettyPrintList(a []Term) string {
	if len(a) == 0 {
		return ""
	}

	left := a[0]
	right := a[1]

	if right.GetType() == T_Fact && right.(*Fact).Head == "|" {
		rightStr := prettyPrintList(right.(*Fact).Args)
		if rightStr == "" {
			return fmt.Sprintf("%s", left)
		}
		return fmt.Sprintf("%s,%s", left, rightStr)
	} else {
		return fmt.Sprintf("%s|%s", left, right)
	}
}
