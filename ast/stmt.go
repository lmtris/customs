package ast

import (
	"fmt"
)

type StmtVisitor interface {
	VisitLetStmt(LetStmt)
	VisitAssertStmt(AssertStmt)
	VisitConstraintStmt(ConstraintStmt)
	VisitBlock(Block)
}

type Stmt interface {
	Accept(StmtVisitor)
}

type Block []Stmt

func (r Block) String() string {
	str := "BlockStmt {\n"
	for _, stmt := range r {
		str += fmt.Sprintf("  %v\n", stmt)
	}
	str += "}"
	return str
}

func (r Block) Accept(visitor StmtVisitor) {
	visitor.VisitBlock(r)
}

type LetStmt struct {
	Ident Token
	Exp   Expr
}

func (r LetStmt) String() string {
	return fmt.Sprintf("LetStmt {Ident=%s Exp=%s}", r.Ident, r.Exp)
}

func (r LetStmt) Accept(visitor StmtVisitor) {
	visitor.VisitLetStmt(r)
}

type AssertStmt struct {
	Ident         Token
	Alias         Token
	Exps          []Expr
	NestedAsserts []Stmt
}

func (r AssertStmt) String() string {
	return fmt.Sprintf("AssertStmt {Ident=%s Alias=%s Exps=%s NestedAsserts=%v}", r.Ident, r.Alias, r.Exps, r.NestedAsserts)
}

func (r AssertStmt) Accept(visitor StmtVisitor) {
	visitor.VisitAssertStmt(r)
}

type ConstraintStmt struct {
	Abstract bool
	Ident    Token
	Block    Block
	Extends  *Token
}

func (r ConstraintStmt) String() string {
	return fmt.Sprintf("ConstraintStmt {Abstract=%v Ident=%v %v Extends=%v}", r.Abstract, r.Ident, r.Block, r.Extends)
}

func (r ConstraintStmt) Accept(visitor StmtVisitor) {
	visitor.VisitConstraintStmt(r)
}
