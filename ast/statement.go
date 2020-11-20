package ast

import (
	"fmt"
	"strings"
    "encoding/json"
)

// Statement can be a Query, Rule or Fact.
//    Its primary tenant is that is will be evaluatable as a top-level executable statement
type Statement interface {
}

type Query []Fact

func (q *Query) String() string {
	stringList := []string{}

	for _, v := range *q {
		stringList = append(stringList, v.String())
	}

	return fmt.Sprintf("?- %s", strings.Join(stringList, ","))
}

func (q *Query) MarshalJSON() ([]byte, error) {
    m := make(map[string]interface{})
    m["t"] = "query"
    m["b"] = *q
    return json.Marshal(m)
}


type Rule struct {
	Head Fact
	Body []Fact
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
