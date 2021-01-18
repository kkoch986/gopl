
# GOPL

```
package  "github.com/kkoch986/gopl"
```


## Statements

At the highest level, a source file is a `StatementList`.

```
StatementList 
  : StatementList Statement
  | Statement
  ;

Statement 
  : Query "."
  | Fact "."
  | Rule "."
  ;

Query : "?-" Concatenation ;

Rule : Fact ":-" Concatenation ;
```

## Concatenation

A concatenation is a series of predicates joined by a comma. 
This indicates the `AND` operation is being applied to each.

```
Concatenation 
  : Concatenation "," Fact
  | Concatenation "," MathAssignment
  | Fact
  | MathAssignment
  ;
```

## Facts

Facts are the most ground things...TODO flesh out details

for example

    a()
    fun()
    square(2)
    "Complex Fact Name"()
    rect(10,50)

```
Fact 
  : Infix 
  | List
  | atom"()"
  | string_lit"()"
  | atom"(" ArgList ")"
  | string_lit"(" ArgList ")"
  ;

Infix : Arg infix_operator Arg ;

FactList
  : FactList "," Fact
  | Fact
  ;

ArgList 
  : ArgList "," Arg
  | Arg
  ;

Arg
  : string_lit
  | num_lit
  | atom
  | var
  | Fact 
  ;
```
# TODO: add support for `is <math expr>`

## Lists

TODO: outline how lists are just syntactic sugar and get converted to `|/0` or `|/2` facts

```
List 
  : "[]"
  | "[" Cons "]"
  | "[" ArgList "]"
  ;

Cons : ArgList "|" ArgList ;
```


## Primitives

These are the most basic primitive types. TODO: more details here

```
atom : lowcase {letter|number|'_'} ;
var : (upcase|'_') {letter|number|'_'} ;
string_lit : '"' {not "\\\"" | '\\' any "\\\"nrt"} '"' ;
num_lit : ['-'] number {number} ['.' {number}] ;
```

## Operators

```
infix_operator : '=';
```

## Math Expressions

```
Factor
  : num_lit
  | var
  | "(" MathExpr ")" ;

Mult
  : Factor "*" Factor
  | Factor "/" Factor
  | Factor ;

MathExpr
  : Mult "+" Mult
  | Mult "-" Mult
  | Mult ;

MathAssignment
  : var "is" MathExpr ;
