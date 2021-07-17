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
	T_MathExpr
	T_MathAssignment
	T_Mult
	T_Factor
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

// TODO: converting query to a statement list for now, should wrap this in some protection
//        to ensure it only ever contains Fact or MathAssignements.
type Query []Statement

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
		// parse into a map first and check the type
		// we might expect either a fact or a MathAssignment
		temp := make(map[string]json.RawMessage)
		err = json.Unmarshal(v, &temp)
		if err != nil {
			return err
		}
		t := string(temp["t"])
		if t == "ma" {
			ma := &MathAssignment{}
			err = json.Unmarshal(v, ma)
			if err != nil {
				return err
			}
			*q = append(*q, ma)
		} else if t == "fact" {
			f := &Fact{}
			err = json.Unmarshal(v, f)
			if err != nil {
				return err
			}
			*q = append(*q, f)
		}
	}
	return nil
}

func CreateRule(h *Fact, q ...Statement) *Rule {
	query := Query(q)
	return &Rule{
		h,
		&query,
	}
}

func (q *Query) Empty() bool {
	return len(*q) == 0
}

func (q *Query) Head() Statement {
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
	Body *Query
}

func (r *Rule) GetType() TermType {
	return T_Rule
}

func (r *Rule) Signature() *Signature {
	return r.Head.Signature()
}

func (r *Rule) String() string {
	stringList := []string{}

	for _, v := range *r.Body {
		stringList = append(stringList, (v).String())
	}

	return fmt.Sprintf("%s :- %s", r.Head, strings.Join(stringList, ","))
}

func (r *Rule) Anonymize(start int, prefix string) (*Rule, map[string]string, int) {
	used := 0
	anonymousBody := Query{}
	existing := make(map[string]string)

	// anonymize the head of the fact and set those bindings
	anonymousHead, moreUsed := r.Head.Anonymize(start+used, prefix, &existing)
	used = used + moreUsed

	// now start anonymizing the body
	for _, f := range *r.Body {
		/// TODO: handle this for mathassignments
		af, u := f.(*Fact).Anonymize(start+used, prefix, &existing)
		used = used + u
		anonymousBody = append(anonymousBody, af)
	}

	return &Rule{anonymousHead, &anonymousBody}, existing, used
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
	// unmarshal the body
	err = json.Unmarshal(rm["b"], &r.Body)
	if err != nil {
		return err
	}

	return nil
}
