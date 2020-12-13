package ast

import (
	"encoding/json"
	"errors"
	"fmt"
)

/**
 * Term is an interface for things that can be placed in an argument list
 * Primarily, we need to be able to identify the type so it can be cast
 * and handled correctly at evaluation time
 */
type Term interface {
	GetType() TermType
	String() string
}

func UnmarshalJSONTerm(b []byte) (Term, error) {
	rm := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &rm)
	if err != nil {
		return nil, err
	}

	// unmarshal the type to a string
	var t string
	err = json.Unmarshal(rm["t"], &t)
	if err != nil {
		return nil, err
	}

	// based on the type, populate S with the correct statement type
	switch t {
	case "var":
		v := &Variable{}
		err = json.Unmarshal(b, v)
		return v, err
    case "fact":
		f := &Fact{}
		err = json.Unmarshal(b, f)
		return f, err
	case "query":
		q := &Query{}
		err = json.Unmarshal(b, q)
		return q, err
	case "rule":
		r := &Rule{}
		err = json.Unmarshal(b, r)
		return r, err
	case "atom":
		v := &Atom{}
		err = json.Unmarshal(b, v)
		return v, err
	case "str":
		v := &StringLiteral{}
		err = json.Unmarshal(b, v)
		return v, err
	case "num":
		v := &NumericLiteral{}
		err = json.Unmarshal(b, v)
		return v, err
	default:
		return nil, errors.New(fmt.Sprintf("Unknown raw statement type: %s", t))
	}
}

/**
 * Variable
 */
type Variable struct {
	string
}

func (v *Variable) GetType() TermType {
	return T_Variable
}
func (v *Variable) String() string {
	return v.string
}
func (v *Variable) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{})
	m["t"] = "var"
	m["v"] = v.string
	return json.Marshal(m)
}

func (v *Variable) UnmarshalJSON(b []byte) error {
	rm := make(map[string]json.RawMessage)
	var s string
	err := json.Unmarshal(b, &rm)
	if err != nil {
		return err
	}
	err = json.Unmarshal(rm["v"], &s)
	if err != nil {
		return err
	}
	v.string = s
	return nil
}


/**
 * Atom
 */
type Atom struct {
	string
}

func (v *Atom) GetType() TermType {
	return T_Atom
}
func (v *Atom) String() string {
	return v.string
}
func (v *Atom) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{})
	m["t"] = "atom"
	m["v"] = v.string
	return json.Marshal(m)
}
func (v *Atom) UnmarshalJSON(b []byte) error {
	rm := make(map[string]json.RawMessage)
	var s string
	err := json.Unmarshal(b, &rm)
	if err != nil {
		return err
	}
	err = json.Unmarshal(rm["v"], &s)
	if err != nil {
		return err
	}
	v.string = s
	return nil
}

/**
 * StringLiteral
 */
type StringLiteral struct {
	string
}

func (v *StringLiteral) GetType() TermType {
	return T_String
}
func (v *StringLiteral) String() string {
	return v.string
}
func (v *StringLiteral) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{})
	m["t"] = "str"
	m["v"] = v.string
	return json.Marshal(m)
}
func (v *StringLiteral) UnmarshalJSON(b []byte) error {
	rm := make(map[string]json.RawMessage)
	var s string
	err := json.Unmarshal(b, &rm)
	if err != nil {
		return err
	}
	err = json.Unmarshal(rm["v"], &s)
	if err != nil {
		return err
	}
	v.string = s
	return nil
}

/**
 * NumericLiteral
 */
type NumericLiteral struct {
	float64
}

func (v *NumericLiteral) GetType() TermType {
	return T_Number
}

func (v *NumericLiteral) String() string {
	return fmt.Sprintf("%f", v.float64)
}

func (v *NumericLiteral) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{})
	m["t"] = "num"
	m["v"] = v.float64
	return json.Marshal(m)
}
func (v *NumericLiteral) UnmarshalJSON(b []byte) error {
	rm := make(map[string]json.RawMessage)
	var s float64
	err := json.Unmarshal(b, &rm)
	if err != nil {
		return err
	}
	err = json.Unmarshal(rm["v"], &s)
	if err != nil {
		return err
	}
	v.float64 = s
	return nil
}
