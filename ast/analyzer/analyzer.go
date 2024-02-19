package analyzer

import v2 "customs/ast"

type Analyzer struct {
	Stmt  []v2.Stmt
	stack map[string]v2.Token
	local map[string]v2.Token
}

func NewAnalyzer(stmt []v2.Stmt) Analyzer {
	return Analyzer{Stmt: stmt, stack: make(map[string]v2.Token), local: make(map[string]v2.Token)}
}

func (r *Analyzer) Analyze() {
	for _, stmt := range r.Stmt {
		stmt.Accept(r)
	}
}

func (r *Analyzer) VisitConstraintStmt(stmt v2.ConstraintStmt) {
	//TODO implement me
	panic("implement me")
}

func (r *Analyzer) VisitAssertStmt(stmt v2.AssertStmt) {
	//TODO implement me
	panic("implement me")
}

func (r *Analyzer) VisitAssignStmt(stmt v2.AssignStmt) {
	if _, ok := r.local[stmt.Id.Literal]; ok {
		panic("Variable already declared")
	}
	stmt.Id.LiteralType = stmt.Expr.Accept(r)
	r.local[stmt.Id.Literal] = stmt.Id
}

func (r *Analyzer) VisitBinaryExpr(expr v2.BinaryExpr) (typ v2.LiteralType) {
	switch expr.Op.TokenType {
	case v2.Plus, v2.Minus, v2.Multiply, v2.Divide:
		left, right := expr.Left.Accept(r), expr.Right.Accept(r)
		if left == v2.Integer && left == right {
			typ = v2.Integer
			return
		}

		if left == v2.Float && left == right {
			typ = v2.Float
			return
		}

		// type conversion
		if left == v2.Integer && right == v2.Float {
			typ = v2.Float
			return
		}

		if left == v2.Float && right == v2.Integer {
			typ = v2.Float
			return
		}

		panic("Type mismatch")
	case v2.Equal, v2.NotEqual, v2.LessThan, v2.LessThanOrEqual, v2.GreaterThan, v2.GreaterThanOrEqual:
		left, right := expr.Left.Accept(r), expr.Right.Accept(r)
		if left == right {
			typ = v2.Boolean
			return
		} else {
			panic("Type mismatch")
		}
	case v2.And, v2.Or:
		left, right := expr.Left.Accept(r), expr.Right.Accept(r)
		if left == v2.Boolean && left == right {
			typ = v2.Boolean
			return
		} else {
			panic("Type mismatch")
		}
	}
	return
}

func (r *Analyzer) VisitUnaryExpr(expr v2.UnaryExpr) (typ v2.LiteralType) {
	switch expr.Op.TokenType {
	case v2.Plus, v2.Minus:
		typ = expr.Expr.Accept(r)
		if typ == v2.Integer || typ == v2.Float {
			return
		}
		panic("Type mismatch")
	case v2.Not:
		typ = expr.Expr.Accept(r)
		if typ == v2.Boolean {
			return
		}
		panic("Type mismatch")
	}
	return
}

func (r *Analyzer) VisitToken(token v2.Token) (typ v2.LiteralType) {
	if token.TokenType == v2.Ident {
		if v, ok := r.local[token.Literal]; ok {
			typ = v.LiteralType
			return
		}
		panic("Variable not declared")
	}
	typ = token.LiteralType
	return
}
