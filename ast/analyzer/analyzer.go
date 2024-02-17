package analyzer

import (
	"customs/ast"
	"fmt"
	"slices"
)

type Analyzer struct {
	Stmts []ast.Stmt // global statements
	stack map[string]ast.Token
	local map[string]ast.Token
}

func (r *Analyzer) GetOperatorType(op ast.Token) ast.Type {
	intOps := []ast.TokenType{ast.Plus, ast.Minus, ast.Mul, ast.Div}
	boolOps := []ast.TokenType{ast.Equal, ast.NotEqual, ast.Lt, ast.Lte, ast.Gt, ast.Gte}

	if slices.Contains(intOps, op.TokenType) {
		return ast.Integer
	}

	if slices.Contains(boolOps, op.TokenType) {
		return ast.Bool
	}

	return ast.Any
}

func (r *Analyzer) VisitBinaryExp(expr ast.BinaryExpr) (typ ast.Type) {

	left, right := expr.Left.Accept(r), expr.Right.Accept(r)

	switch r.GetOperatorType(expr.Operator) {
	case ast.Integer:
		if left == ast.Integer && left == right {
			typ = ast.Integer
		} else {
			panic(ast.TypeMismatchError(expr.Operator, left, right))
		}

	case ast.Bool:
		if left == right {
			typ = ast.Bool
		} else {
			panic(ast.TypeMismatchError(expr.Operator, left, right))
		}

	default:
		panic(fmt.Sprintf("Unexpected operator: %v", expr.Operator))
	}

	return
}

func (r *Analyzer) VisitUnaryExp(expr ast.UnaryExpr) ast.Type {
	return expr.Right.Accept(r)
}

func (r *Analyzer) VisitToken(token ast.Token) ast.Type {
	switch typ := token.TokenType; typ {
	case ast.Number:
		return ast.Integer
	case ast.Ident:
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

func NewAnalyzer(stmts []ast.Stmt) Analyzer {
	return Analyzer{Stmts: stmts, stack: make(map[string]ast.Token), local: make(map[string]ast.Token)}
}

func (r *Analyzer) Check() {
	for _, stmt := range r.Stmts {
		stmt.Accept(r)
	}
}

func (r *Analyzer) VisitLetStmt(stmt ast.LetStmt) {
	// Check if variable is already defined
	if v, ok := r.local[stmt.Ident.Literal]; ok {
		panic(fmt.Sprintf("Name %s already defined at line %v col %v", stmt.Ident.Literal, v.Line, v.Column))
	}

	// Type inference
	stmt.Ident.Typ = stmt.Exp.Accept(r)

	r.local[stmt.Ident.Literal] = stmt.Ident
}

func (r *Analyzer) VisitAssertStmt(stmt ast.AssertStmt) {
	// Check existed assert name
	if v, ok := r.stack[stmt.Ident.Literal]; ok {
		panic(fmt.Sprintf("Name %s already defined at line %v col %v", stmt.Ident.Literal, v.Line, v.Column))
	}

	// Check existed alias name
	if v, ok := r.local[stmt.Alias.Literal]; ok {
		panic(fmt.Sprintf("Name %s already defined at line %v col %v", stmt.Alias.Literal, v.Line, v.Column))
	}

	// Type inference, if expr
	stmt.Ident.Typ, stmt.Alias.Typ = ast.Integer, ast.Integer
	r.stack[stmt.Ident.Literal] = stmt.Ident
	r.stack[stmt.Alias.Literal] = stmt.Alias

	for _, exp := range stmt.Exps {
		exp.Accept(r)
	}

	env := NewAnalyzerWithStack(stmt.NestedAsserts, MergeStacks(r.local, r.stack))
	env.Check()
}

func (r *Analyzer) VisitConstraintStmt(stmt ast.ConstraintStmt) {
	// Check if constraint is already defined
	if v, ok := r.stack[stmt.Ident.Literal]; ok {
		panic(fmt.Sprintf("Name %s already defined at line %v col %v", stmt.Ident.Literal, v.Line, v.Column))
	}

	stmt.Ident.Typ = "constraint"
	r.stack[stmt.Ident.Literal] = stmt.Ident

	r.VisitBlock(stmt.Block)
}

func NewAnalyzerWithStack(stmts []ast.Stmt, stack map[string]ast.Token) Analyzer {
	return Analyzer{Stmts: stmts, stack: stack, local: make(map[string]ast.Token)}
}

func (r *Analyzer) VisitBlock(stmt ast.Block) {
	env := NewAnalyzerWithStack(stmt, MergeStacks(r.local, r.stack))
	env.Check()
}

func MergeStacks(local, stack map[string]ast.Token) (st map[string]ast.Token) {
	st = make(map[string]ast.Token)

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
