// Package symbols is generated by gogll. Do not edit.
package symbols

type Symbol interface {
	isSymbol()
	IsNonTerminal() bool
	String() string
}

func (NT) isSymbol() {}
func (T) isSymbol()  {}

// NT is the type of non-terminals symbols
type NT int

const (
	NT_Arg NT = iota
	NT_ArgList
	NT_Concatenation
	NT_Cons
	NT_Fact
	NT_FactList
	NT_Infix
	NT_List
	NT_Query
	NT_Rule
	NT_Statement
	NT_StatementList
)

// T is the type of terminals symbols
type T int

const (
	T_0  T = iota // (
	T_1           // ()
	T_2           // )
	T_3           // ,
	T_4           // .
	T_5           // :-
	T_6           // ?-
	T_7           // [
	T_8           // []
	T_9           // ]
	T_10          // atom
	T_11          // infix_operator
	T_12          // num_lit
	T_13          // string_lit
	T_14          // var
	T_15          // |
)

type Symbols []Symbol

func (ss Symbols) Strings() []string {
	strs := make([]string, len(ss))
	for i, s := range ss {
		strs[i] = s.String()
	}
	return strs
}

func (NT) IsNonTerminal() bool {
	return true
}

func (T) IsNonTerminal() bool {
	return false
}

func (nt NT) String() string {
	return ntToString[nt]
}

func (t T) String() string {
	return tToString[t]
}

var ntToString = []string{
	"Arg",           /* NT_Arg */
	"ArgList",       /* NT_ArgList */
	"Concatenation", /* NT_Concatenation */
	"Cons",          /* NT_Cons */
	"Fact",          /* NT_Fact */
	"FactList",      /* NT_FactList */
	"Infix",         /* NT_Infix */
	"List",          /* NT_List */
	"Query",         /* NT_Query */
	"Rule",          /* NT_Rule */
	"Statement",     /* NT_Statement */
	"StatementList", /* NT_StatementList */
}

var tToString = []string{
	"(",              /* T_0 */
	"()",             /* T_1 */
	")",              /* T_2 */
	",",              /* T_3 */
	".",              /* T_4 */
	":-",             /* T_5 */
	"?-",             /* T_6 */
	"[",              /* T_7 */
	"[]",             /* T_8 */
	"]",              /* T_9 */
	"atom",           /* T_10 */
	"infix_operator", /* T_11 */
	"num_lit",        /* T_12 */
	"string_lit",     /* T_13 */
	"var",            /* T_14 */
	"|",              /* T_15 */
}

var stringNT = map[string]NT{
	"Arg":           NT_Arg,
	"ArgList":       NT_ArgList,
	"Concatenation": NT_Concatenation,
	"Cons":          NT_Cons,
	"Fact":          NT_Fact,
	"FactList":      NT_FactList,
	"Infix":         NT_Infix,
	"List":          NT_List,
	"Query":         NT_Query,
	"Rule":          NT_Rule,
	"Statement":     NT_Statement,
	"StatementList": NT_StatementList,
}
