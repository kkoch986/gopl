package indexer

import (
	"github.com/kkoch986/gopl/ast"
)

type Indexer interface {
	IndexStatement(ast.Statement)
	StatementsForSignature(*ast.Signature) []ast.Statement
}
