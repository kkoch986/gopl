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

func (f *Fact) Anonymize(start int, existing *map[string]string) (*Fact, int) {
	used := 0
	anonymousBody := []Term{}
	for _, v := range f.Args {
		t := v.GetType()
		switch t {
		case T_Variable:
			if (*existing)[v.String()] != "" {
				anonymousBody = append(anonymousBody, &Variable{(*existing)[v.String()]})
			} else {
				newVal := fmt.Sprintf("_h%d", start+used)
				anonymousBody = append(anonymousBody, &Variable{newVal})
				(*existing)[v.String()] = newVal
				used = used + 1
			}
		case T_Fact:
			af, u := v.(*Fact).Anonymize(start+used, existing)
			used = used + u
			anonymousBody = append(anonymousBody, af)
		default:
			anonymousBody = append(anonymousBody, v)
		}
	}

	return &Fact{f.Head, anonymousBody}, used
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

func (f *Fact) ExtractVariables() []*Variable {
	ret := []*Variable{}

	for _, v := range f.Args {
		t := v.GetType()
		switch t {
		case T_Fact:
			for _, v2 := range v.(*Fact).ExtractVariables() {
				ret = append(ret, v2)
			}
		case T_Variable:
			ret = append(ret, v.(*Variable))
		}
	}

	return ret
}

func (f *Fact) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{})
	m["t"] = "fact"
	m["f"] = f.Head
	m["a"] = f.Args
	return json.Marshal(m)
}

func (f *Fact) UnmarshalJSON(b []byte) error {
	rm := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &rm)
	if err != nil {
		return err
	}

	// parse the head
	err = json.Unmarshal(rm["f"], &f.Head)
	if err != nil {
		return err
	}

	// parse the args
	var rmArgs []json.RawMessage
	err = json.Unmarshal(rm["a"], &rmArgs)
	if err != nil {
		return err
	}

	for _, v := range rmArgs {
		t, err := UnmarshalJSONTerm(v)
		if err != nil {
			return err
		}
		f.Args = append(f.Args, t)
	}

	return nil
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
