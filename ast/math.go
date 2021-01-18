package ast

import (
	"fmt"
)

type Factor struct {
	Var  *Variable
	Num  *NumericLiteral
	Expr *MathExpr
}

func (f *Factor) Anonymize(bindings *map[string]string, start int) (*Factor, int) {
	if f.Num != nil {
		return f, 0
	}
	if f.Var != nil {
		varName := f.Var.String()
		bound := (*bindings)[varName]
		if bound != "" {
			return &Factor{CreateVariable(bound), nil, nil}, 0
		} else {
			newVar := fmt.Sprintf("_h%d", start)
			(*bindings)[varName] = newVar
			return &Factor{CreateVariable(newVar), nil, nil}, 1
		}
	}
	if f.Expr != nil {
		newExpr, used := f.Expr.Anonymize(bindings, start)
		return &Factor{Expr: newExpr}, used
	}
	// i dont think we should normally get down here
	return f, 0
}

func (f *Factor) String() string {
	if f.Var != nil {
		return f.Var.String()
	} else if f.Num != nil {
		return f.Num.String()
	} else {
		return f.Expr.String()
	}
}

type MultOperator int

const (
	OP_Mult MultOperator = iota
	OP_Divide
	OP_MultNoOp
)

type Mult struct {
	LHS      *Factor
	Operator MultOperator
	RHS      *Factor
}

func (m *Mult) Anonymize(bindings *map[string]string, start int) (*Mult, int) {
	lhs, lused := m.LHS.Anonymize(bindings, start)
	rhs, rused := m.RHS.Anonymize(bindings, start+lused)
	return &Mult{lhs, m.Operator, rhs}, (rused + lused)
}

func (m *Mult) String() string {
	switch m.Operator {
	case OP_MultNoOp:
		return m.LHS.String()
	case OP_Mult:
		return fmt.Sprintf("(%s * %s)", m.LHS.String(), m.RHS.String())
	case OP_Divide:
		return fmt.Sprintf("(%s / %s)", m.LHS.String(), m.RHS.String())
	default:
		return fmt.Sprintf("Unknown Mult operation: %s", m.Operator)
	}
}

type MathExprOperator int

const (
	OP_Add MathExprOperator = iota
	OP_Subtract
	OP_MathExprNoOp
)

type MathExpr struct {
	LHS      *Mult
	Operator MathExprOperator
	RHS      *Mult
}

func (m *MathExpr) GetType() TermType {
	return T_MathExpr
}

func (m *MathExpr) String() string {
	switch m.Operator {
	case OP_MathExprNoOp:
		return m.LHS.String()
	case OP_Subtract:
		return fmt.Sprintf("(%s - %s)", m.LHS.String(), m.RHS.String())
	case OP_Add:
		return fmt.Sprintf("(%s + %s)", m.LHS.String(), m.RHS.String())
	default:
		return fmt.Sprintf("Unknown MathExpr operation: %s", m.Operator)
	}
}

func (m *MathExpr) Anonymize(bindings *map[string]string, start int) (*MathExpr, int) {
	lhs, used := m.LHS.Anonymize(bindings, start)
	start = start + used
	rhs, rused := m.RHS.Anonymize(bindings, start+used)

	return &MathExpr{lhs, m.Operator, rhs}, (used + rused)
}

type MathAssignment struct {
	LHS *Variable
	RHS *MathExpr
}

func (m *MathAssignment) GetType() TermType {
	return T_MathAssignment
}

func (m *MathAssignment) String() string {
	return fmt.Sprintf("%s is %s", m.LHS.String(), m.RHS.String())
}

func (m *MathAssignment) Anonymize(start int) (*MathAssignment, int) {
	bindings := make(map[string]string)
	lhsVar := fmt.Sprintf("_h%d", start)
	bindings[m.LHS.String()] = lhsVar
	rhs, count := m.RHS.Anonymize(&bindings, start+1)
	return &MathAssignment{CreateVariable(lhsVar), rhs}, (count + 1)
}
