package ast

import (
	"encoding/json"
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

func (f *Factor) GetType() TermType {
	return T_Factor
}

func (f *Factor) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{})
	m["t"] = "mf"
	if f.Var != nil {
		m["v"] = f.Var
	} else if f.Num != nil {
		m["n"] = f.Num
	} else if f.Expr != nil {
		m["e"] = f.Expr
	}
	return json.Marshal(m)
}

func (f *Factor) UnmarshalJSON(b []byte) error {
	rm := make(map[string]json.RawMessage)
	var v *Variable
	var n *NumericLiteral
	var e *MathExpr
	var val json.RawMessage
	var ok bool

	err := json.Unmarshal(b, &rm)
	if err != nil {
		return err
	}

	if val, ok = rm["v"]; ok {
		err = json.Unmarshal(val, v)
		if err != nil {
			return err
		}
		f.Var = v
	} else if val, ok = rm["n"]; ok {
		err = json.Unmarshal(val, n)
		if err != nil {
			return err
		}
		f.Num = n
	} else if val, ok = rm["e"]; ok {
		err = json.Unmarshal(val, e)
		if err != nil {
			return err
		}
		f.Expr = e
	}

	return nil
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

func (m *Mult) GetType() TermType {
	return T_Mult
}

func (me *Mult) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{})
	m["t"] = "mu"
	m["l"] = me.LHS
	if me.RHS != nil {
		m["r"] = me.RHS
	}
	m["o"] = me.Operator
	return json.Marshal(m)
}

func (mu *Mult) UnmarshalJSON(b []byte) error {
	rm := make(map[string]json.RawMessage)
	var lhs, rhs Factor
	var op MultOperator

	err := json.Unmarshal(b, &rm)
	if err != nil {
		return err
	}
	err = json.Unmarshal(rm["l"], &lhs)
	if err != nil {
		return err
	}

	err = json.Unmarshal(rm["r"], &rhs)
	if err != nil {
		return err
	}

	err = json.Unmarshal(rm["o"], &op)
	if err != nil {
		return err
	}

	// put it all back together
	mu.LHS = &lhs
	mu.RHS = &rhs
	mu.Operator = op

	return nil
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

func (me *MathExpr) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{})
	m["t"] = "me"
	m["l"] = me.LHS
	if me.RHS != nil {
		m["r"] = me.RHS
	}
	m["o"] = me.Operator
	return json.Marshal(m)
}

func (me *MathExpr) UnmarshalJSON(b []byte) error {
	rm := make(map[string]json.RawMessage)
	var lhs, rhs Mult
	var op MathExprOperator

	err := json.Unmarshal(b, &rm)
	if err != nil {
		return err
	}
	err = json.Unmarshal(rm["l"], &lhs)
	if err != nil {
		return err
	}

	err = json.Unmarshal(rm["r"], &rhs)
	if err != nil {
		return err
	}

	err = json.Unmarshal(rm["o"], &op)
	if err != nil {
		return err
	}

	// put it all back together
	me.LHS = &lhs
	me.RHS = &rhs
	me.Operator = op

	return nil
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

func (ma *MathAssignment) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{})
	m["t"] = "ma"
	m["v"] = ma.LHS
	m["e"] = ma.RHS
	return json.Marshal(m)
}

func (ma *MathAssignment) UnmarshalJSON(b []byte) error {
	fmt.Println("UNMARSHAL MATHASSIGN")
	rm := make(map[string]json.RawMessage)
	var v Variable
	var rhs MathExpr

	err := json.Unmarshal(b, &rm)
	if err != nil {
		return err
	}
	err = json.Unmarshal(rm["v"], &v)
	if err != nil {
		return err
	}

	err = json.Unmarshal(rm["e"], &rhs)
	if err != nil {
		return err
	}

	// put it all back together
	ma.LHS = &v
	ma.RHS = &rhs

	return nil
}
