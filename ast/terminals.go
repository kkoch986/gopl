package ast

import (
	"fmt"
    "encoding/json"
)

/**
 * Arg is an interface for things that can be placed in an argument list
 * Primarily, we need to be able to identify the type so it can be cast
 * and handled correctly at evaluation time
 */
type Arg interface {
	GetType() string
	String() string
}

/**
 * Variable
 */
type Variable struct {
	string
}

func (v *Variable) GetType() string {
	return "Variable"
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

func (v *Atom) GetType() string {
	return "Atom"
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

func (v *StringLiteral) GetType() string {
	return "String"
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

func (v *NumericLiteral) GetType() string {
	return "Number"
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
