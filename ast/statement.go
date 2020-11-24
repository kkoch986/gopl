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

type Query []Fact

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

func (q *Query) Empty() bool {
	return len(*q) == 0
}

func (q *Query) Head() *Fact {
	if q.Empty() {
		return nil
	}
	return &(*q)[0]
}

func (q *Query) Tail() *Query {
	t := (*q)[1:]
	return &t
}

type Rule struct {
	Head Fact
	Body []Fact
}

func (r *Rule) GetType() TermType {
	return T_Rule
}

func (q *Rule) String() string {
	stringList := []string{}

	for _, v := range q.Body {
		stringList = append(stringList, (&v).String())
	}

	return fmt.Sprintf("%s :- %s", &q.Head, strings.Join(stringList, ","))
}

func (q *Rule) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{})
	m["t"] = "rule"
	m["h"] = q.Head
	m["b"] = q.Body
	return json.Marshal(m)
}
