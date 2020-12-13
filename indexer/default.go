package indexer

import (
	"fmt"

	"github.com/kkoch986/gopl/ast"
)

type Default struct {
	bySig   map[string][]ast.Statement
	nextVar int
}

func NewDefault() *Default {
	return &Default{
		bySig:   make(map[string][]ast.Statement),
		nextVar: 0,
	}
}

// TODO: prevent duplicates of the same facts from being indexed
func (d *Default) IndexStatement(s ast.Statement) {
	switch s.GetType() {
	case ast.T_Fact:
		d.indexFact(s.(*ast.Fact))
	case ast.T_Rule:
		d.indexRule(s.(*ast.Rule))
	}
}

func (d *Default) indexFact(f *ast.Fact) {
	d.bySig[f.Signature().String()] = append(d.bySig[f.Signature().String()], f)
}

func (d *Default) indexRule(r *ast.Rule) {
	d.bySig[r.Signature().String()] = append(d.bySig[r.Signature().String()], r)
}

func (d *Default) StatementsForSignature(s *ast.Signature) []ast.Statement {
	return d.bySig[s.String()]
}

func (d *Default) nextAnonymousVariable() string {
	d.nextVar = d.nextVar + 1
	return fmt.Sprintf("_h%d", d.nextVar)
}
