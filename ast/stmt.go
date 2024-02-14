package ast

import "fmt"

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

func (r Block) Accept(visitor StmtVisitor) {
	visitor.VisitBlock(r)
}

type LetStmt struct {
	Ident Token
	Exp   Expr
}

func (r LetStmt) String() string {
	return fmt.Sprintf("LetStmt(%s, %s)", r.Ident, r.Exp)

}

func (r LetStmt) Accept(visitor StmtVisitor) {
	visitor.VisitLetStmt(r)
}

type AssertStmt struct {
	Ident Token
	Alias Token
	Exps  []Expr
}

func (r AssertStmt) Accept(visitor StmtVisitor) {
	visitor.VisitAssertStmt(r)
}

type ConstraintStmt struct {
	Abstract bool
	Ident    Token
	Block    Block
	Extends  Token
}

func (r ConstraintStmt) Accept(visitor StmtVisitor) {
	visitor.VisitConstraintStmt(r)
}
