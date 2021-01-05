package ast

import (
	"encoding/json"
	"fmt"
	"strings"
)

type TermType int

const (
	T_Query TermType = iota
	T_Rule
	T_Fact
	T_Variable
	T_Atom
	T_String
	T_Number
)

func (s TermType) String() string {
	return []string{"Query", "Rule", "Fact", "Variable", "Atom", "String", "Number"}[s]
}

// Statement can be a Query, Rule or Fact.
//    Its primary tenant is that is will be evaluatable as a top-level executable statement
type Statement interface {
	GetType() TermType
	String() string
}

type Query []*Fact

func (q *Query) String() string {
	stringList := []string{}

	for _, v := range *q {
		stringList = append(stringList, v.String())
	}

	return fmt.Sprintf("?- %s", strings.Join(stringList, ","))
}

func (q *Query) GetType() TermType {
	return T_Query
}

func (q *Query) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{})
	m["t"] = "query"
	m["b"] = *q
	return json.Marshal(m)
}

func (q *Query) UnmarshalJSON(b []byte) error {
	rm := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &rm)
	if err != nil {
		return err
	}

	// parse the args
	var rmArgs []json.RawMessage
	err = json.Unmarshal(rm["b"], &rmArgs)
	if err != nil {
		return err
	}

	for _, v := range rmArgs {
		f := &Fact{}
		err = json.Unmarshal(v, f)
		if err != nil {
			return err
		}
		*q = append(*q, f)
	}

	return nil
}

func (q *Query) Empty() bool {
	return len(*q) == 0
}

func (q *Query) Head() *Fact {
	if q.Empty() {
		return nil
	}
	return (*q)[0]
}

func (q *Query) Tail() *Query {
	t := (*q)[1:]
	return &t
}

type Rule struct {
	Head *Fact
	Body []*Fact
}

func (r *Rule) GetType() TermType {
	return T_Rule
}

func (r *Rule) Signature() *Signature {
	return r.Head.Signature()
}

func (q *Rule) String() string {
	stringList := []string{}

	for _, v := range q.Body {
		stringList = append(stringList, (v).String())
	}

	return fmt.Sprintf("%s :- %s", q.Head, strings.Join(stringList, ","))
}

func (q *Rule) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{})
	m["t"] = "rule"
	m["h"] = q.Head
	m["b"] = q.Body
	return json.Marshal(m)
}

func (r *Rule) UnmarshalJSON(b []byte) error {
	rm := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &rm)
	if err != nil {
		return err
	}

	// unmarshal the head
	err = json.Unmarshal(rm["h"], &r.Head)
	if err != nil {
		return err
	}

	// TODO: unmarshal the body
	err = json.Unmarshal(rm["b"], &r.Body)
	if err != nil {
		return err
	}

	return nil
}
