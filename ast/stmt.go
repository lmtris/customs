package ast

import (
	"fmt"
)

type StmtVisitor interface {
	VisitAssignStmt(AssignStmt)
	VisitConstraintStmt(ConstraintStmt)
	VisitAssertStmt(AssertStmt)
}

type Stmt interface {
	Accept(StmtVisitor)
}

type AssignStmt struct {
	Id   Token
	Expr Expr
}

func (r AssignStmt) String() string {
	return fmt.Sprintf("AssignStmt{%s %s}", r.Id, PrefixTraversal(r.Expr))
}

func (r AssignStmt) Accept(v StmtVisitor) {
	v.VisitAssignStmt(r)
}

type ConstraintStmt struct {
	IsAbstract       bool
	Id               Token
	ParentConstraint Token
	LetStmts         []AssignStmt
	AssertStmts      []AssertStmt
}

func (r ConstraintStmt) String() string {
	return fmt.Sprintf("ConstraintStmt{%s %v %v}", r.Id, r.LetStmts, r.AssertStmts)
}

func (r ConstraintStmt) Accept(v StmtVisitor) {
	v.VisitConstraintStmt(r)
}

type AssertStmt struct {
	Id    Token
	Alias Token
	Exprs []Expr
	Stmts []AssertStmt
}

func (r AssertStmt) String() string {
	return fmt.Sprintf("AssertStmt{%s %s %v}", r.Id, r.Alias, r.Exprs)
}

func (r AssertStmt) Accept(v StmtVisitor) {
	v.VisitAssertStmt(r)
}

func PrefixTraversal(expr Expr) string {
	switch v := expr.(type) {
	case Token:
		return v.Literal
	case UnaryExpr:
		return v.Op.Literal + "(" + PrefixTraversal(v.Expr) + ")"
	case BinaryExpr:
		return "(" + v.Op.Literal + " " + PrefixTraversal(v.Left) + " " + PrefixTraversal(v.Right) + ")"
	}
	return ""
}
