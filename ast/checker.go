package ast

import (
	"fmt"
	"slices"
)

type Checker struct {
	stmts   []Stmt // global statements
	stack   map[string]Token
	checker *Checker
}

func (r *Checker) VisitBinaryExp(expr BinaryExpr) (typ Type) {
	intOperators := []TokenType{Plus, Minus, Mul, Div}
	boolOperators := []TokenType{Equal, NotEqual, Lt, Lte, Gt, Gte}

	if slices.Contains(intOperators, expr.Operator.TokenType) {
		left := expr.Left.Accept(r)
		right := expr.Right.Accept(r)

		if left != Integer || right != Integer {
			panic(fmt.Sprintf("Type mismatch: Expected integer type, got %v and %v", left, right))
		}

		return Integer

	}

	if slices.Contains(boolOperators, expr.Operator.TokenType) {
		left := expr.Left.Accept(r)
		right := expr.Right.Accept(r)

		if left != right {
			panic(fmt.Sprintf("Type mismatch: Expected same type, got %v and %v", left, right))
		}

		return Bool
	}
	return
}

func (r *Checker) VisitUnaryExp(expr UnaryExpr) Type {
	return expr.Right.Accept(r)
}

func (r *Checker) VisitToken(token Token) Type {
	switch typ := token.TokenType; typ {
	case Number:
		return Integer
	case Ident:
		if v, ok := r.stack[token.Literal]; ok {
			return v.Typ
		}
		panic(fmt.Sprintf("Undefined variable: %s", token.Literal))
	default:
		panic(fmt.Sprintf("%v:%v Unexpected token: %v", token.Line, token.Column, token))
	}
}

func NewChecker(stmts []Stmt) Checker {
	return Checker{stmts: stmts, stack: make(map[string]Token)}
}

func (r *Checker) Check() {
	for _, stmt := range r.stmts {
		stmt.Accept(r)
	}
}

func (r *Checker) VisitLetStmt(stmt LetStmt) {
	// Check if variable is already defined
	if v, ok := r.stack[stmt.Ident.Literal]; ok {
		panic(fmt.Sprintf("Name %s already defined at line %v col %v", stmt.Ident.Literal, v.Line, v.Column))
	}

	// Type inference
	stmt.Ident.Typ = stmt.Exp.Accept(r)

	r.stack[stmt.Ident.Literal] = stmt.Ident
}

func (r *Checker) VisitAssertStmt(stmt AssertStmt) {

}

func (r *Checker) VisitConstraintStmt(stmt ConstraintStmt) {
	// Check if constraint is already defined
	if v, ok := r.stack[stmt.Ident.Literal]; ok {
		panic(fmt.Sprintf("Name %s already defined at line %v col %v", stmt.Ident.Literal, v.Line, v.Column))
	}

	r.stack[stmt.Ident.Literal] = stmt.Ident
}

func (r *Checker) VisitBlock(stmt Block) {

}
