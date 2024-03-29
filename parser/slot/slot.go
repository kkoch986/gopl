// Package slot is generated by gogll. Do not edit.
package slot

import (
	"bytes"
	"fmt"

	"github.com/kkoch986/gopl/parser/symbols"
)

type Label int

const (
	Arg0R0 Label = iota
	Arg0R1
	Arg1R0
	Arg1R1
	Arg2R0
	Arg2R1
	Arg3R0
	Arg3R1
	Arg4R0
	Arg4R1
	ArgList0R0
	ArgList0R1
	ArgList0R2
	ArgList0R3
	ArgList1R0
	ArgList1R1
	Concatenation0R0
	Concatenation0R1
	Concatenation0R2
	Concatenation0R3
	Concatenation1R0
	Concatenation1R1
	Concatenation1R2
	Concatenation1R3
	Concatenation2R0
	Concatenation2R1
	Concatenation3R0
	Concatenation3R1
	Cons0R0
	Cons0R1
	Cons0R2
	Cons0R3
	Fact0R0
	Fact0R1
	Fact1R0
	Fact1R1
	Fact2R0
	Fact2R1
	Fact2R2
	Fact3R0
	Fact3R1
	Fact3R2
	Fact4R0
	Fact4R1
	Fact4R2
	Fact4R3
	Fact4R4
	Fact5R0
	Fact5R1
	Fact5R2
	Fact5R3
	Fact5R4
	FactList0R0
	FactList0R1
	FactList0R2
	FactList0R3
	FactList1R0
	FactList1R1
	Factor0R0
	Factor0R1
	Factor1R0
	Factor1R1
	Factor2R0
	Factor2R1
	Factor2R2
	Factor2R3
	Infix0R0
	Infix0R1
	Infix0R2
	Infix0R3
	List0R0
	List0R1
	List1R0
	List1R1
	List1R2
	List1R3
	List2R0
	List2R1
	List2R2
	List2R3
	MathAssignment0R0
	MathAssignment0R1
	MathAssignment0R2
	MathAssignment0R3
	MathExpr0R0
	MathExpr0R1
	MathExpr0R2
	MathExpr0R3
	MathExpr1R0
	MathExpr1R1
	MathExpr1R2
	MathExpr1R3
	MathExpr2R0
	MathExpr2R1
	Mult0R0
	Mult0R1
	Mult0R2
	Mult0R3
	Mult1R0
	Mult1R1
	Mult1R2
	Mult1R3
	Mult2R0
	Mult2R1
	Query0R0
	Query0R1
	Query0R2
	Rule0R0
	Rule0R1
	Rule0R2
	Rule0R3
	Statement0R0
	Statement0R1
	Statement0R2
	Statement1R0
	Statement1R1
	Statement1R2
	Statement2R0
	Statement2R1
	Statement2R2
	StatementList0R0
	StatementList0R1
	StatementList0R2
	StatementList1R0
	StatementList1R1
)

type Slot struct {
	NT      symbols.NT
	Alt     int
	Pos     int
	Symbols symbols.Symbols
	Label   Label
}

type Index struct {
	NT  symbols.NT
	Alt int
	Pos int
}

func GetAlternates(nt symbols.NT) []Label {
	alts, exist := alternates[nt]
	if !exist {
		panic(fmt.Sprintf("Invalid NT %s", nt))
	}
	return alts
}

func GetLabel(nt symbols.NT, alt, pos int) Label {
	l, exist := slotIndex[Index{nt, alt, pos}]
	if exist {
		return l
	}
	panic(fmt.Sprintf("Error: no slot label for NT=%s, alt=%d, pos=%d", nt, alt, pos))
}

func (l Label) EoR() bool {
	return l.Slot().EoR()
}

func (l Label) Head() symbols.NT {
	return l.Slot().NT
}

func (l Label) Index() Index {
	s := l.Slot()
	return Index{s.NT, s.Alt, s.Pos}
}

func (l Label) Alternate() int {
	return l.Slot().Alt
}

func (l Label) Pos() int {
	return l.Slot().Pos
}

func (l Label) Slot() *Slot {
	s, exist := slots[l]
	if !exist {
		panic(fmt.Sprintf("Invalid slot label %d", l))
	}
	return s
}

func (l Label) String() string {
	return l.Slot().String()
}

func (l Label) Symbols() symbols.Symbols {
	return l.Slot().Symbols
}

func (s *Slot) EoR() bool {
	return s.Pos >= len(s.Symbols)
}

func (s *Slot) String() string {
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, "%s : ", s.NT)
	for i, sym := range s.Symbols {
		if i == s.Pos {
			fmt.Fprintf(buf, "∙")
		}
		fmt.Fprintf(buf, "%s ", sym)
	}
	if s.Pos >= len(s.Symbols) {
		fmt.Fprintf(buf, "∙")
	}
	return buf.String()
}

var slots = map[Label]*Slot{
	Arg0R0: {
		symbols.NT_Arg, 0, 0,
		symbols.Symbols{
			symbols.T_18,
		},
		Arg0R0,
	},
	Arg0R1: {
		symbols.NT_Arg, 0, 1,
		symbols.Symbols{
			symbols.T_18,
		},
		Arg0R1,
	},
	Arg1R0: {
		symbols.NT_Arg, 1, 0,
		symbols.Symbols{
			symbols.T_17,
		},
		Arg1R0,
	},
	Arg1R1: {
		symbols.NT_Arg, 1, 1,
		symbols.Symbols{
			symbols.T_17,
		},
		Arg1R1,
	},
	Arg2R0: {
		symbols.NT_Arg, 2, 0,
		symbols.Symbols{
			symbols.T_14,
		},
		Arg2R0,
	},
	Arg2R1: {
		symbols.NT_Arg, 2, 1,
		symbols.Symbols{
			symbols.T_14,
		},
		Arg2R1,
	},
	Arg3R0: {
		symbols.NT_Arg, 3, 0,
		symbols.Symbols{
			symbols.T_19,
		},
		Arg3R0,
	},
	Arg3R1: {
		symbols.NT_Arg, 3, 1,
		symbols.Symbols{
			symbols.T_19,
		},
		Arg3R1,
	},
	Arg4R0: {
		symbols.NT_Arg, 4, 0,
		symbols.Symbols{
			symbols.NT_Fact,
		},
		Arg4R0,
	},
	Arg4R1: {
		symbols.NT_Arg, 4, 1,
		symbols.Symbols{
			symbols.NT_Fact,
		},
		Arg4R1,
	},
	ArgList0R0: {
		symbols.NT_ArgList, 0, 0,
		symbols.Symbols{
			symbols.NT_ArgList,
			symbols.T_5,
			symbols.NT_Arg,
		},
		ArgList0R0,
	},
	ArgList0R1: {
		symbols.NT_ArgList, 0, 1,
		symbols.Symbols{
			symbols.NT_ArgList,
			symbols.T_5,
			symbols.NT_Arg,
		},
		ArgList0R1,
	},
	ArgList0R2: {
		symbols.NT_ArgList, 0, 2,
		symbols.Symbols{
			symbols.NT_ArgList,
			symbols.T_5,
			symbols.NT_Arg,
		},
		ArgList0R2,
	},
	ArgList0R3: {
		symbols.NT_ArgList, 0, 3,
		symbols.Symbols{
			symbols.NT_ArgList,
			symbols.T_5,
			symbols.NT_Arg,
		},
		ArgList0R3,
	},
	ArgList1R0: {
		symbols.NT_ArgList, 1, 0,
		symbols.Symbols{
			symbols.NT_Arg,
		},
		ArgList1R0,
	},
	ArgList1R1: {
		symbols.NT_ArgList, 1, 1,
		symbols.Symbols{
			symbols.NT_Arg,
		},
		ArgList1R1,
	},
	Concatenation0R0: {
		symbols.NT_Concatenation, 0, 0,
		symbols.Symbols{
			symbols.NT_Concatenation,
			symbols.T_5,
			symbols.NT_Fact,
		},
		Concatenation0R0,
	},
	Concatenation0R1: {
		symbols.NT_Concatenation, 0, 1,
		symbols.Symbols{
			symbols.NT_Concatenation,
			symbols.T_5,
			symbols.NT_Fact,
		},
		Concatenation0R1,
	},
	Concatenation0R2: {
		symbols.NT_Concatenation, 0, 2,
		symbols.Symbols{
			symbols.NT_Concatenation,
			symbols.T_5,
			symbols.NT_Fact,
		},
		Concatenation0R2,
	},
	Concatenation0R3: {
		symbols.NT_Concatenation, 0, 3,
		symbols.Symbols{
			symbols.NT_Concatenation,
			symbols.T_5,
			symbols.NT_Fact,
		},
		Concatenation0R3,
	},
	Concatenation1R0: {
		symbols.NT_Concatenation, 1, 0,
		symbols.Symbols{
			symbols.NT_Concatenation,
			symbols.T_5,
			symbols.NT_MathAssignment,
		},
		Concatenation1R0,
	},
	Concatenation1R1: {
		symbols.NT_Concatenation, 1, 1,
		symbols.Symbols{
			symbols.NT_Concatenation,
			symbols.T_5,
			symbols.NT_MathAssignment,
		},
		Concatenation1R1,
	},
	Concatenation1R2: {
		symbols.NT_Concatenation, 1, 2,
		symbols.Symbols{
			symbols.NT_Concatenation,
			symbols.T_5,
			symbols.NT_MathAssignment,
		},
		Concatenation1R2,
	},
	Concatenation1R3: {
		symbols.NT_Concatenation, 1, 3,
		symbols.Symbols{
			symbols.NT_Concatenation,
			symbols.T_5,
			symbols.NT_MathAssignment,
		},
		Concatenation1R3,
	},
	Concatenation2R0: {
		symbols.NT_Concatenation, 2, 0,
		symbols.Symbols{
			symbols.NT_Fact,
		},
		Concatenation2R0,
	},
	Concatenation2R1: {
		symbols.NT_Concatenation, 2, 1,
		symbols.Symbols{
			symbols.NT_Fact,
		},
		Concatenation2R1,
	},
	Concatenation3R0: {
		symbols.NT_Concatenation, 3, 0,
		symbols.Symbols{
			symbols.NT_MathAssignment,
		},
		Concatenation3R0,
	},
	Concatenation3R1: {
		symbols.NT_Concatenation, 3, 1,
		symbols.Symbols{
			symbols.NT_MathAssignment,
		},
		Concatenation3R1,
	},
	Cons0R0: {
		symbols.NT_Cons, 0, 0,
		symbols.Symbols{
			symbols.NT_ArgList,
			symbols.T_20,
			symbols.NT_ArgList,
		},
		Cons0R0,
	},
	Cons0R1: {
		symbols.NT_Cons, 0, 1,
		symbols.Symbols{
			symbols.NT_ArgList,
			symbols.T_20,
			symbols.NT_ArgList,
		},
		Cons0R1,
	},
	Cons0R2: {
		symbols.NT_Cons, 0, 2,
		symbols.Symbols{
			symbols.NT_ArgList,
			symbols.T_20,
			symbols.NT_ArgList,
		},
		Cons0R2,
	},
	Cons0R3: {
		symbols.NT_Cons, 0, 3,
		symbols.Symbols{
			symbols.NT_ArgList,
			symbols.T_20,
			symbols.NT_ArgList,
		},
		Cons0R3,
	},
	Fact0R0: {
		symbols.NT_Fact, 0, 0,
		symbols.Symbols{
			symbols.NT_Infix,
		},
		Fact0R0,
	},
	Fact0R1: {
		symbols.NT_Fact, 0, 1,
		symbols.Symbols{
			symbols.NT_Infix,
		},
		Fact0R1,
	},
	Fact1R0: {
		symbols.NT_Fact, 1, 0,
		symbols.Symbols{
			symbols.NT_List,
		},
		Fact1R0,
	},
	Fact1R1: {
		symbols.NT_Fact, 1, 1,
		symbols.Symbols{
			symbols.NT_List,
		},
		Fact1R1,
	},
	Fact2R0: {
		symbols.NT_Fact, 2, 0,
		symbols.Symbols{
			symbols.T_14,
			symbols.T_1,
		},
		Fact2R0,
	},
	Fact2R1: {
		symbols.NT_Fact, 2, 1,
		symbols.Symbols{
			symbols.T_14,
			symbols.T_1,
		},
		Fact2R1,
	},
	Fact2R2: {
		symbols.NT_Fact, 2, 2,
		symbols.Symbols{
			symbols.T_14,
			symbols.T_1,
		},
		Fact2R2,
	},
	Fact3R0: {
		symbols.NT_Fact, 3, 0,
		symbols.Symbols{
			symbols.T_18,
			symbols.T_1,
		},
		Fact3R0,
	},
	Fact3R1: {
		symbols.NT_Fact, 3, 1,
		symbols.Symbols{
			symbols.T_18,
			symbols.T_1,
		},
		Fact3R1,
	},
	Fact3R2: {
		symbols.NT_Fact, 3, 2,
		symbols.Symbols{
			symbols.T_18,
			symbols.T_1,
		},
		Fact3R2,
	},
	Fact4R0: {
		symbols.NT_Fact, 4, 0,
		symbols.Symbols{
			symbols.T_14,
			symbols.T_0,
			symbols.NT_ArgList,
			symbols.T_2,
		},
		Fact4R0,
	},
	Fact4R1: {
		symbols.NT_Fact, 4, 1,
		symbols.Symbols{
			symbols.T_14,
			symbols.T_0,
			symbols.NT_ArgList,
			symbols.T_2,
		},
		Fact4R1,
	},
	Fact4R2: {
		symbols.NT_Fact, 4, 2,
		symbols.Symbols{
			symbols.T_14,
			symbols.T_0,
			symbols.NT_ArgList,
			symbols.T_2,
		},
		Fact4R2,
	},
	Fact4R3: {
		symbols.NT_Fact, 4, 3,
		symbols.Symbols{
			symbols.T_14,
			symbols.T_0,
			symbols.NT_ArgList,
			symbols.T_2,
		},
		Fact4R3,
	},
	Fact4R4: {
		symbols.NT_Fact, 4, 4,
		symbols.Symbols{
			symbols.T_14,
			symbols.T_0,
			symbols.NT_ArgList,
			symbols.T_2,
		},
		Fact4R4,
	},
	Fact5R0: {
		symbols.NT_Fact, 5, 0,
		symbols.Symbols{
			symbols.T_18,
			symbols.T_0,
			symbols.NT_ArgList,
			symbols.T_2,
		},
		Fact5R0,
	},
	Fact5R1: {
		symbols.NT_Fact, 5, 1,
		symbols.Symbols{
			symbols.T_18,
			symbols.T_0,
			symbols.NT_ArgList,
			symbols.T_2,
		},
		Fact5R1,
	},
	Fact5R2: {
		symbols.NT_Fact, 5, 2,
		symbols.Symbols{
			symbols.T_18,
			symbols.T_0,
			symbols.NT_ArgList,
			symbols.T_2,
		},
		Fact5R2,
	},
	Fact5R3: {
		symbols.NT_Fact, 5, 3,
		symbols.Symbols{
			symbols.T_18,
			symbols.T_0,
			symbols.NT_ArgList,
			symbols.T_2,
		},
		Fact5R3,
	},
	Fact5R4: {
		symbols.NT_Fact, 5, 4,
		symbols.Symbols{
			symbols.T_18,
			symbols.T_0,
			symbols.NT_ArgList,
			symbols.T_2,
		},
		Fact5R4,
	},
	FactList0R0: {
		symbols.NT_FactList, 0, 0,
		symbols.Symbols{
			symbols.NT_FactList,
			symbols.T_5,
			symbols.NT_Fact,
		},
		FactList0R0,
	},
	FactList0R1: {
		symbols.NT_FactList, 0, 1,
		symbols.Symbols{
			symbols.NT_FactList,
			symbols.T_5,
			symbols.NT_Fact,
		},
		FactList0R1,
	},
	FactList0R2: {
		symbols.NT_FactList, 0, 2,
		symbols.Symbols{
			symbols.NT_FactList,
			symbols.T_5,
			symbols.NT_Fact,
		},
		FactList0R2,
	},
	FactList0R3: {
		symbols.NT_FactList, 0, 3,
		symbols.Symbols{
			symbols.NT_FactList,
			symbols.T_5,
			symbols.NT_Fact,
		},
		FactList0R3,
	},
	FactList1R0: {
		symbols.NT_FactList, 1, 0,
		symbols.Symbols{
			symbols.NT_Fact,
		},
		FactList1R0,
	},
	FactList1R1: {
		symbols.NT_FactList, 1, 1,
		symbols.Symbols{
			symbols.NT_Fact,
		},
		FactList1R1,
	},
	Factor0R0: {
		symbols.NT_Factor, 0, 0,
		symbols.Symbols{
			symbols.T_17,
		},
		Factor0R0,
	},
	Factor0R1: {
		symbols.NT_Factor, 0, 1,
		symbols.Symbols{
			symbols.T_17,
		},
		Factor0R1,
	},
	Factor1R0: {
		symbols.NT_Factor, 1, 0,
		symbols.Symbols{
			symbols.T_19,
		},
		Factor1R0,
	},
	Factor1R1: {
		symbols.NT_Factor, 1, 1,
		symbols.Symbols{
			symbols.T_19,
		},
		Factor1R1,
	},
	Factor2R0: {
		symbols.NT_Factor, 2, 0,
		symbols.Symbols{
			symbols.T_0,
			symbols.NT_MathExpr,
			symbols.T_2,
		},
		Factor2R0,
	},
	Factor2R1: {
		symbols.NT_Factor, 2, 1,
		symbols.Symbols{
			symbols.T_0,
			symbols.NT_MathExpr,
			symbols.T_2,
		},
		Factor2R1,
	},
	Factor2R2: {
		symbols.NT_Factor, 2, 2,
		symbols.Symbols{
			symbols.T_0,
			symbols.NT_MathExpr,
			symbols.T_2,
		},
		Factor2R2,
	},
	Factor2R3: {
		symbols.NT_Factor, 2, 3,
		symbols.Symbols{
			symbols.T_0,
			symbols.NT_MathExpr,
			symbols.T_2,
		},
		Factor2R3,
	},
	Infix0R0: {
		symbols.NT_Infix, 0, 0,
		symbols.Symbols{
			symbols.NT_Arg,
			symbols.T_15,
			symbols.NT_Arg,
		},
		Infix0R0,
	},
	Infix0R1: {
		symbols.NT_Infix, 0, 1,
		symbols.Symbols{
			symbols.NT_Arg,
			symbols.T_15,
			symbols.NT_Arg,
		},
		Infix0R1,
	},
	Infix0R2: {
		symbols.NT_Infix, 0, 2,
		symbols.Symbols{
			symbols.NT_Arg,
			symbols.T_15,
			symbols.NT_Arg,
		},
		Infix0R2,
	},
	Infix0R3: {
		symbols.NT_Infix, 0, 3,
		symbols.Symbols{
			symbols.NT_Arg,
			symbols.T_15,
			symbols.NT_Arg,
		},
		Infix0R3,
	},
	List0R0: {
		symbols.NT_List, 0, 0,
		symbols.Symbols{
			symbols.T_12,
		},
		List0R0,
	},
	List0R1: {
		symbols.NT_List, 0, 1,
		symbols.Symbols{
			symbols.T_12,
		},
		List0R1,
	},
	List1R0: {
		symbols.NT_List, 1, 0,
		symbols.Symbols{
			symbols.T_11,
			symbols.NT_Cons,
			symbols.T_13,
		},
		List1R0,
	},
	List1R1: {
		symbols.NT_List, 1, 1,
		symbols.Symbols{
			symbols.T_11,
			symbols.NT_Cons,
			symbols.T_13,
		},
		List1R1,
	},
	List1R2: {
		symbols.NT_List, 1, 2,
		symbols.Symbols{
			symbols.T_11,
			symbols.NT_Cons,
			symbols.T_13,
		},
		List1R2,
	},
	List1R3: {
		symbols.NT_List, 1, 3,
		symbols.Symbols{
			symbols.T_11,
			symbols.NT_Cons,
			symbols.T_13,
		},
		List1R3,
	},
	List2R0: {
		symbols.NT_List, 2, 0,
		symbols.Symbols{
			symbols.T_11,
			symbols.NT_ArgList,
			symbols.T_13,
		},
		List2R0,
	},
	List2R1: {
		symbols.NT_List, 2, 1,
		symbols.Symbols{
			symbols.T_11,
			symbols.NT_ArgList,
			symbols.T_13,
		},
		List2R1,
	},
	List2R2: {
		symbols.NT_List, 2, 2,
		symbols.Symbols{
			symbols.T_11,
			symbols.NT_ArgList,
			symbols.T_13,
		},
		List2R2,
	},
	List2R3: {
		symbols.NT_List, 2, 3,
		symbols.Symbols{
			symbols.T_11,
			symbols.NT_ArgList,
			symbols.T_13,
		},
		List2R3,
	},
	MathAssignment0R0: {
		symbols.NT_MathAssignment, 0, 0,
		symbols.Symbols{
			symbols.T_19,
			symbols.T_16,
			symbols.NT_MathExpr,
		},
		MathAssignment0R0,
	},
	MathAssignment0R1: {
		symbols.NT_MathAssignment, 0, 1,
		symbols.Symbols{
			symbols.T_19,
			symbols.T_16,
			symbols.NT_MathExpr,
		},
		MathAssignment0R1,
	},
	MathAssignment0R2: {
		symbols.NT_MathAssignment, 0, 2,
		symbols.Symbols{
			symbols.T_19,
			symbols.T_16,
			symbols.NT_MathExpr,
		},
		MathAssignment0R2,
	},
	MathAssignment0R3: {
		symbols.NT_MathAssignment, 0, 3,
		symbols.Symbols{
			symbols.T_19,
			symbols.T_16,
			symbols.NT_MathExpr,
		},
		MathAssignment0R3,
	},
	MathExpr0R0: {
		symbols.NT_MathExpr, 0, 0,
		symbols.Symbols{
			symbols.NT_Mult,
			symbols.T_4,
			symbols.NT_Mult,
		},
		MathExpr0R0,
	},
	MathExpr0R1: {
		symbols.NT_MathExpr, 0, 1,
		symbols.Symbols{
			symbols.NT_Mult,
			symbols.T_4,
			symbols.NT_Mult,
		},
		MathExpr0R1,
	},
	MathExpr0R2: {
		symbols.NT_MathExpr, 0, 2,
		symbols.Symbols{
			symbols.NT_Mult,
			symbols.T_4,
			symbols.NT_Mult,
		},
		MathExpr0R2,
	},
	MathExpr0R3: {
		symbols.NT_MathExpr, 0, 3,
		symbols.Symbols{
			symbols.NT_Mult,
			symbols.T_4,
			symbols.NT_Mult,
		},
		MathExpr0R3,
	},
	MathExpr1R0: {
		symbols.NT_MathExpr, 1, 0,
		symbols.Symbols{
			symbols.NT_Mult,
			symbols.T_6,
			symbols.NT_Mult,
		},
		MathExpr1R0,
	},
	MathExpr1R1: {
		symbols.NT_MathExpr, 1, 1,
		symbols.Symbols{
			symbols.NT_Mult,
			symbols.T_6,
			symbols.NT_Mult,
		},
		MathExpr1R1,
	},
	MathExpr1R2: {
		symbols.NT_MathExpr, 1, 2,
		symbols.Symbols{
			symbols.NT_Mult,
			symbols.T_6,
			symbols.NT_Mult,
		},
		MathExpr1R2,
	},
	MathExpr1R3: {
		symbols.NT_MathExpr, 1, 3,
		symbols.Symbols{
			symbols.NT_Mult,
			symbols.T_6,
			symbols.NT_Mult,
		},
		MathExpr1R3,
	},
	MathExpr2R0: {
		symbols.NT_MathExpr, 2, 0,
		symbols.Symbols{
			symbols.NT_Mult,
		},
		MathExpr2R0,
	},
	MathExpr2R1: {
		symbols.NT_MathExpr, 2, 1,
		symbols.Symbols{
			symbols.NT_Mult,
		},
		MathExpr2R1,
	},
	Mult0R0: {
		symbols.NT_Mult, 0, 0,
		symbols.Symbols{
			symbols.NT_Factor,
			symbols.T_3,
			symbols.NT_Factor,
		},
		Mult0R0,
	},
	Mult0R1: {
		symbols.NT_Mult, 0, 1,
		symbols.Symbols{
			symbols.NT_Factor,
			symbols.T_3,
			symbols.NT_Factor,
		},
		Mult0R1,
	},
	Mult0R2: {
		symbols.NT_Mult, 0, 2,
		symbols.Symbols{
			symbols.NT_Factor,
			symbols.T_3,
			symbols.NT_Factor,
		},
		Mult0R2,
	},
	Mult0R3: {
		symbols.NT_Mult, 0, 3,
		symbols.Symbols{
			symbols.NT_Factor,
			symbols.T_3,
			symbols.NT_Factor,
		},
		Mult0R3,
	},
	Mult1R0: {
		symbols.NT_Mult, 1, 0,
		symbols.Symbols{
			symbols.NT_Factor,
			symbols.T_8,
			symbols.NT_Factor,
		},
		Mult1R0,
	},
	Mult1R1: {
		symbols.NT_Mult, 1, 1,
		symbols.Symbols{
			symbols.NT_Factor,
			symbols.T_8,
			symbols.NT_Factor,
		},
		Mult1R1,
	},
	Mult1R2: {
		symbols.NT_Mult, 1, 2,
		symbols.Symbols{
			symbols.NT_Factor,
			symbols.T_8,
			symbols.NT_Factor,
		},
		Mult1R2,
	},
	Mult1R3: {
		symbols.NT_Mult, 1, 3,
		symbols.Symbols{
			symbols.NT_Factor,
			symbols.T_8,
			symbols.NT_Factor,
		},
		Mult1R3,
	},
	Mult2R0: {
		symbols.NT_Mult, 2, 0,
		symbols.Symbols{
			symbols.NT_Factor,
		},
		Mult2R0,
	},
	Mult2R1: {
		symbols.NT_Mult, 2, 1,
		symbols.Symbols{
			symbols.NT_Factor,
		},
		Mult2R1,
	},
	Query0R0: {
		symbols.NT_Query, 0, 0,
		symbols.Symbols{
			symbols.T_10,
			symbols.NT_Concatenation,
		},
		Query0R0,
	},
	Query0R1: {
		symbols.NT_Query, 0, 1,
		symbols.Symbols{
			symbols.T_10,
			symbols.NT_Concatenation,
		},
		Query0R1,
	},
	Query0R2: {
		symbols.NT_Query, 0, 2,
		symbols.Symbols{
			symbols.T_10,
			symbols.NT_Concatenation,
		},
		Query0R2,
	},
	Rule0R0: {
		symbols.NT_Rule, 0, 0,
		symbols.Symbols{
			symbols.NT_Fact,
			symbols.T_9,
			symbols.NT_Concatenation,
		},
		Rule0R0,
	},
	Rule0R1: {
		symbols.NT_Rule, 0, 1,
		symbols.Symbols{
			symbols.NT_Fact,
			symbols.T_9,
			symbols.NT_Concatenation,
		},
		Rule0R1,
	},
	Rule0R2: {
		symbols.NT_Rule, 0, 2,
		symbols.Symbols{
			symbols.NT_Fact,
			symbols.T_9,
			symbols.NT_Concatenation,
		},
		Rule0R2,
	},
	Rule0R3: {
		symbols.NT_Rule, 0, 3,
		symbols.Symbols{
			symbols.NT_Fact,
			symbols.T_9,
			symbols.NT_Concatenation,
		},
		Rule0R3,
	},
	Statement0R0: {
		symbols.NT_Statement, 0, 0,
		symbols.Symbols{
			symbols.NT_Query,
			symbols.T_7,
		},
		Statement0R0,
	},
	Statement0R1: {
		symbols.NT_Statement, 0, 1,
		symbols.Symbols{
			symbols.NT_Query,
			symbols.T_7,
		},
		Statement0R1,
	},
	Statement0R2: {
		symbols.NT_Statement, 0, 2,
		symbols.Symbols{
			symbols.NT_Query,
			symbols.T_7,
		},
		Statement0R2,
	},
	Statement1R0: {
		symbols.NT_Statement, 1, 0,
		symbols.Symbols{
			symbols.NT_Fact,
			symbols.T_7,
		},
		Statement1R0,
	},
	Statement1R1: {
		symbols.NT_Statement, 1, 1,
		symbols.Symbols{
			symbols.NT_Fact,
			symbols.T_7,
		},
		Statement1R1,
	},
	Statement1R2: {
		symbols.NT_Statement, 1, 2,
		symbols.Symbols{
			symbols.NT_Fact,
			symbols.T_7,
		},
		Statement1R2,
	},
	Statement2R0: {
		symbols.NT_Statement, 2, 0,
		symbols.Symbols{
			symbols.NT_Rule,
			symbols.T_7,
		},
		Statement2R0,
	},
	Statement2R1: {
		symbols.NT_Statement, 2, 1,
		symbols.Symbols{
			symbols.NT_Rule,
			symbols.T_7,
		},
		Statement2R1,
	},
	Statement2R2: {
		symbols.NT_Statement, 2, 2,
		symbols.Symbols{
			symbols.NT_Rule,
			symbols.T_7,
		},
		Statement2R2,
	},
	StatementList0R0: {
		symbols.NT_StatementList, 0, 0,
		symbols.Symbols{
			symbols.NT_StatementList,
			symbols.NT_Statement,
		},
		StatementList0R0,
	},
	StatementList0R1: {
		symbols.NT_StatementList, 0, 1,
		symbols.Symbols{
			symbols.NT_StatementList,
			symbols.NT_Statement,
		},
		StatementList0R1,
	},
	StatementList0R2: {
		symbols.NT_StatementList, 0, 2,
		symbols.Symbols{
			symbols.NT_StatementList,
			symbols.NT_Statement,
		},
		StatementList0R2,
	},
	StatementList1R0: {
		symbols.NT_StatementList, 1, 0,
		symbols.Symbols{
			symbols.NT_Statement,
		},
		StatementList1R0,
	},
	StatementList1R1: {
		symbols.NT_StatementList, 1, 1,
		symbols.Symbols{
			symbols.NT_Statement,
		},
		StatementList1R1,
	},
}

var slotIndex = map[Index]Label{
	Index{symbols.NT_Arg, 0, 0}:            Arg0R0,
	Index{symbols.NT_Arg, 0, 1}:            Arg0R1,
	Index{symbols.NT_Arg, 1, 0}:            Arg1R0,
	Index{symbols.NT_Arg, 1, 1}:            Arg1R1,
	Index{symbols.NT_Arg, 2, 0}:            Arg2R0,
	Index{symbols.NT_Arg, 2, 1}:            Arg2R1,
	Index{symbols.NT_Arg, 3, 0}:            Arg3R0,
	Index{symbols.NT_Arg, 3, 1}:            Arg3R1,
	Index{symbols.NT_Arg, 4, 0}:            Arg4R0,
	Index{symbols.NT_Arg, 4, 1}:            Arg4R1,
	Index{symbols.NT_ArgList, 0, 0}:        ArgList0R0,
	Index{symbols.NT_ArgList, 0, 1}:        ArgList0R1,
	Index{symbols.NT_ArgList, 0, 2}:        ArgList0R2,
	Index{symbols.NT_ArgList, 0, 3}:        ArgList0R3,
	Index{symbols.NT_ArgList, 1, 0}:        ArgList1R0,
	Index{symbols.NT_ArgList, 1, 1}:        ArgList1R1,
	Index{symbols.NT_Concatenation, 0, 0}:  Concatenation0R0,
	Index{symbols.NT_Concatenation, 0, 1}:  Concatenation0R1,
	Index{symbols.NT_Concatenation, 0, 2}:  Concatenation0R2,
	Index{symbols.NT_Concatenation, 0, 3}:  Concatenation0R3,
	Index{symbols.NT_Concatenation, 1, 0}:  Concatenation1R0,
	Index{symbols.NT_Concatenation, 1, 1}:  Concatenation1R1,
	Index{symbols.NT_Concatenation, 1, 2}:  Concatenation1R2,
	Index{symbols.NT_Concatenation, 1, 3}:  Concatenation1R3,
	Index{symbols.NT_Concatenation, 2, 0}:  Concatenation2R0,
	Index{symbols.NT_Concatenation, 2, 1}:  Concatenation2R1,
	Index{symbols.NT_Concatenation, 3, 0}:  Concatenation3R0,
	Index{symbols.NT_Concatenation, 3, 1}:  Concatenation3R1,
	Index{symbols.NT_Cons, 0, 0}:           Cons0R0,
	Index{symbols.NT_Cons, 0, 1}:           Cons0R1,
	Index{symbols.NT_Cons, 0, 2}:           Cons0R2,
	Index{symbols.NT_Cons, 0, 3}:           Cons0R3,
	Index{symbols.NT_Fact, 0, 0}:           Fact0R0,
	Index{symbols.NT_Fact, 0, 1}:           Fact0R1,
	Index{symbols.NT_Fact, 1, 0}:           Fact1R0,
	Index{symbols.NT_Fact, 1, 1}:           Fact1R1,
	Index{symbols.NT_Fact, 2, 0}:           Fact2R0,
	Index{symbols.NT_Fact, 2, 1}:           Fact2R1,
	Index{symbols.NT_Fact, 2, 2}:           Fact2R2,
	Index{symbols.NT_Fact, 3, 0}:           Fact3R0,
	Index{symbols.NT_Fact, 3, 1}:           Fact3R1,
	Index{symbols.NT_Fact, 3, 2}:           Fact3R2,
	Index{symbols.NT_Fact, 4, 0}:           Fact4R0,
	Index{symbols.NT_Fact, 4, 1}:           Fact4R1,
	Index{symbols.NT_Fact, 4, 2}:           Fact4R2,
	Index{symbols.NT_Fact, 4, 3}:           Fact4R3,
	Index{symbols.NT_Fact, 4, 4}:           Fact4R4,
	Index{symbols.NT_Fact, 5, 0}:           Fact5R0,
	Index{symbols.NT_Fact, 5, 1}:           Fact5R1,
	Index{symbols.NT_Fact, 5, 2}:           Fact5R2,
	Index{symbols.NT_Fact, 5, 3}:           Fact5R3,
	Index{symbols.NT_Fact, 5, 4}:           Fact5R4,
	Index{symbols.NT_FactList, 0, 0}:       FactList0R0,
	Index{symbols.NT_FactList, 0, 1}:       FactList0R1,
	Index{symbols.NT_FactList, 0, 2}:       FactList0R2,
	Index{symbols.NT_FactList, 0, 3}:       FactList0R3,
	Index{symbols.NT_FactList, 1, 0}:       FactList1R0,
	Index{symbols.NT_FactList, 1, 1}:       FactList1R1,
	Index{symbols.NT_Factor, 0, 0}:         Factor0R0,
	Index{symbols.NT_Factor, 0, 1}:         Factor0R1,
	Index{symbols.NT_Factor, 1, 0}:         Factor1R0,
	Index{symbols.NT_Factor, 1, 1}:         Factor1R1,
	Index{symbols.NT_Factor, 2, 0}:         Factor2R0,
	Index{symbols.NT_Factor, 2, 1}:         Factor2R1,
	Index{symbols.NT_Factor, 2, 2}:         Factor2R2,
	Index{symbols.NT_Factor, 2, 3}:         Factor2R3,
	Index{symbols.NT_Infix, 0, 0}:          Infix0R0,
	Index{symbols.NT_Infix, 0, 1}:          Infix0R1,
	Index{symbols.NT_Infix, 0, 2}:          Infix0R2,
	Index{symbols.NT_Infix, 0, 3}:          Infix0R3,
	Index{symbols.NT_List, 0, 0}:           List0R0,
	Index{symbols.NT_List, 0, 1}:           List0R1,
	Index{symbols.NT_List, 1, 0}:           List1R0,
	Index{symbols.NT_List, 1, 1}:           List1R1,
	Index{symbols.NT_List, 1, 2}:           List1R2,
	Index{symbols.NT_List, 1, 3}:           List1R3,
	Index{symbols.NT_List, 2, 0}:           List2R0,
	Index{symbols.NT_List, 2, 1}:           List2R1,
	Index{symbols.NT_List, 2, 2}:           List2R2,
	Index{symbols.NT_List, 2, 3}:           List2R3,
	Index{symbols.NT_MathAssignment, 0, 0}: MathAssignment0R0,
	Index{symbols.NT_MathAssignment, 0, 1}: MathAssignment0R1,
	Index{symbols.NT_MathAssignment, 0, 2}: MathAssignment0R2,
	Index{symbols.NT_MathAssignment, 0, 3}: MathAssignment0R3,
	Index{symbols.NT_MathExpr, 0, 0}:       MathExpr0R0,
	Index{symbols.NT_MathExpr, 0, 1}:       MathExpr0R1,
	Index{symbols.NT_MathExpr, 0, 2}:       MathExpr0R2,
	Index{symbols.NT_MathExpr, 0, 3}:       MathExpr0R3,
	Index{symbols.NT_MathExpr, 1, 0}:       MathExpr1R0,
	Index{symbols.NT_MathExpr, 1, 1}:       MathExpr1R1,
	Index{symbols.NT_MathExpr, 1, 2}:       MathExpr1R2,
	Index{symbols.NT_MathExpr, 1, 3}:       MathExpr1R3,
	Index{symbols.NT_MathExpr, 2, 0}:       MathExpr2R0,
	Index{symbols.NT_MathExpr, 2, 1}:       MathExpr2R1,
	Index{symbols.NT_Mult, 0, 0}:           Mult0R0,
	Index{symbols.NT_Mult, 0, 1}:           Mult0R1,
	Index{symbols.NT_Mult, 0, 2}:           Mult0R2,
	Index{symbols.NT_Mult, 0, 3}:           Mult0R3,
	Index{symbols.NT_Mult, 1, 0}:           Mult1R0,
	Index{symbols.NT_Mult, 1, 1}:           Mult1R1,
	Index{symbols.NT_Mult, 1, 2}:           Mult1R2,
	Index{symbols.NT_Mult, 1, 3}:           Mult1R3,
	Index{symbols.NT_Mult, 2, 0}:           Mult2R0,
	Index{symbols.NT_Mult, 2, 1}:           Mult2R1,
	Index{symbols.NT_Query, 0, 0}:          Query0R0,
	Index{symbols.NT_Query, 0, 1}:          Query0R1,
	Index{symbols.NT_Query, 0, 2}:          Query0R2,
	Index{symbols.NT_Rule, 0, 0}:           Rule0R0,
	Index{symbols.NT_Rule, 0, 1}:           Rule0R1,
	Index{symbols.NT_Rule, 0, 2}:           Rule0R2,
	Index{symbols.NT_Rule, 0, 3}:           Rule0R3,
	Index{symbols.NT_Statement, 0, 0}:      Statement0R0,
	Index{symbols.NT_Statement, 0, 1}:      Statement0R1,
	Index{symbols.NT_Statement, 0, 2}:      Statement0R2,
	Index{symbols.NT_Statement, 1, 0}:      Statement1R0,
	Index{symbols.NT_Statement, 1, 1}:      Statement1R1,
	Index{symbols.NT_Statement, 1, 2}:      Statement1R2,
	Index{symbols.NT_Statement, 2, 0}:      Statement2R0,
	Index{symbols.NT_Statement, 2, 1}:      Statement2R1,
	Index{symbols.NT_Statement, 2, 2}:      Statement2R2,
	Index{symbols.NT_StatementList, 0, 0}:  StatementList0R0,
	Index{symbols.NT_StatementList, 0, 1}:  StatementList0R1,
	Index{symbols.NT_StatementList, 0, 2}:  StatementList0R2,
	Index{symbols.NT_StatementList, 1, 0}:  StatementList1R0,
	Index{symbols.NT_StatementList, 1, 1}:  StatementList1R1,
}

var alternates = map[symbols.NT][]Label{
	symbols.NT_StatementList:  []Label{StatementList0R0, StatementList1R0},
	symbols.NT_Statement:      []Label{Statement0R0, Statement1R0, Statement2R0},
	symbols.NT_Query:          []Label{Query0R0},
	symbols.NT_Rule:           []Label{Rule0R0},
	symbols.NT_Concatenation:  []Label{Concatenation0R0, Concatenation1R0, Concatenation2R0, Concatenation3R0},
	symbols.NT_Fact:           []Label{Fact0R0, Fact1R0, Fact2R0, Fact3R0, Fact4R0, Fact5R0},
	symbols.NT_Infix:          []Label{Infix0R0},
	symbols.NT_FactList:       []Label{FactList0R0, FactList1R0},
	symbols.NT_ArgList:        []Label{ArgList0R0, ArgList1R0},
	symbols.NT_Arg:            []Label{Arg0R0, Arg1R0, Arg2R0, Arg3R0, Arg4R0},
	symbols.NT_List:           []Label{List0R0, List1R0, List2R0},
	symbols.NT_Cons:           []Label{Cons0R0},
	symbols.NT_Factor:         []Label{Factor0R0, Factor1R0, Factor2R0},
	symbols.NT_Mult:           []Label{Mult0R0, Mult1R0, Mult2R0},
	symbols.NT_MathExpr:       []Label{MathExpr0R0, MathExpr1R0, MathExpr2R0},
	symbols.NT_MathAssignment: []Label{MathAssignment0R0},
}
