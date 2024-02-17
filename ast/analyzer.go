package ast

import (
	"fmt"
	"slices"
)

type Analyzer struct {
	stmts []Stmt // global statements
	stack map[string]Token
	local map[string]Token
}

func (r *Analyzer) GetOperatorType(op Token) Type {
	intOps := []TokenType{Plus, Minus, Mul, Div}
	boolOps := []TokenType{Equal, NotEqual, Lt, Lte, Gt, Gte}

	if slices.Contains(intOps, op.TokenType) {
		return Integer
	}

	if slices.Contains(boolOps, op.TokenType) {
		return Bool
	}

	return Any
}

func (r *Analyzer) VisitBinaryExp(expr BinaryExpr) (typ Type) {

	left, right := expr.Left.Accept(r), expr.Right.Accept(r)

	switch r.GetOperatorType(expr.Operator) {
	case Integer:
		if left == Integer && left == right {
			typ = Integer
		} else {
			panic(TypeMismatchError(expr.Operator, left, right))
		}

	case Bool:
		if left == right {
			typ = Bool
		} else {
			panic(TypeMismatchError(expr.Operator, left, right))
		}

	default:
		panic(fmt.Sprintf("Unexpected operator: %v", expr.Operator))
	}

	return
}

func (r *Analyzer) VisitUnaryExp(expr UnaryExpr) Type {
	return expr.Right.Accept(r)
}

func (r *Analyzer) VisitToken(token Token) Type {
	switch typ := token.TokenType; typ {
	case Number:
		return Integer
	case Ident:
		// First search in local scope
		if v, ok := r.local[token.Literal]; ok {
			return v.Typ
		}

		// Then search in parent scope
		if v, ok := r.stack[token.Literal]; ok {
			return v.Typ
		}

		panic(fmt.Sprintf("Undefined variable: %s", token.Literal))
	default:
		panic(fmt.Sprintf("%v:%v Unexpected token: %v", token.Line, token.Column, token))
	}
}

func NewAnalyzer(stmts []Stmt) Analyzer {
	return Analyzer{stmts: stmts, stack: make(map[string]Token), local: make(map[string]Token)}
}

func (r *Analyzer) Check() {
	for _, stmt := range r.stmts {
		stmt.Accept(r)
	}
}

func (r *Analyzer) VisitLetStmt(stmt LetStmt) {
	// Check if variable is already defined
	if v, ok := r.stack[stmt.Ident.Literal]; ok {
		panic(fmt.Sprintf("Name %s already defined at line %v col %v", stmt.Ident.Literal, v.Line, v.Column))
	}

	// Type inference
	stmt.Ident.Typ = stmt.Exp.Accept(r)

	r.local[stmt.Ident.Literal] = stmt.Ident
}

func (r *Analyzer) VisitAssertStmt(stmt AssertStmt) {
	// Check existed assert name
	if v, ok := r.stack[stmt.Ident.Literal]; ok {
		panic(fmt.Sprintf("Name %s already defined at line %v col %v", stmt.Ident.Literal, v.Line, v.Column))
	}

	// Check existed alias name
	if v, ok := r.local[stmt.Alias.Literal]; ok {
		panic(fmt.Sprintf("Name %s already defined at line %v col %v", stmt.Alias.Literal, v.Line, v.Column))
	}

	// Type inference, if expr
	stmt.Ident.Typ, stmt.Alias.Typ = Integer, Integer
	r.stack[stmt.Ident.Literal] = stmt.Ident
	r.stack[stmt.Alias.Literal] = stmt.Alias

	for _, exp := range stmt.Exps {
		exp.Accept(r)
	}

	env := NewAnalyzerWithStack(stmt.NestedAsserts, MergeStacks(r.local, r.stack))
	env.Check()
}

func (r *Analyzer) VisitConstraintStmt(stmt ConstraintStmt) {
	// Check if constraint is already defined
	if v, ok := r.stack[stmt.Ident.Literal]; ok {
		panic(fmt.Sprintf("Name %s already defined at line %v col %v", stmt.Ident.Literal, v.Line, v.Column))
	}

	stmt.Ident.Typ = "constraint"
	r.stack[stmt.Ident.Literal] = stmt.Ident

	r.VisitBlock(stmt.Block)
}

func NewAnalyzerWithStack(stmts []Stmt, stack map[string]Token) Analyzer {
	return Analyzer{stmts: stmts, stack: stack, local: make(map[string]Token)}
}

func (r *Analyzer) VisitBlock(stmt Block) {
	env := NewAnalyzerWithStack(stmt, MergeStacks(r.local, r.stack))
	env.Check()
}

func MergeStacks(local, stack map[string]Token) (st map[string]Token) {
	st = make(map[string]Token)

	for k, v := range local {
		st[k] = v
	}

	for k, v := range stack {
		// shadowing variables
		if _, ok := st[k]; !ok {
			st[k] = v
		}
	}

	return
}
