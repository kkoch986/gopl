package ast

import (
	"fmt"
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

/**
 * Atom
 */
type Atom struct {
	string
}

func (v *Atom) GetType() string {
	return "Atom"
}
func (v *Atom) String() string {
	return v.string
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
