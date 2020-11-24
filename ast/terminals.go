package ast

import (
	"encoding/json"
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

/**
 * Atom
 */
type Atom struct {
	string `json:"v"`
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
	m["type"] = "num"
	m["v"] = v.float64
	return json.Marshal(m)
}
