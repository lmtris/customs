package engine

import (
	"customs/ast"
	"strconv"
)

type Resolver struct {
	Stmts []ast.Stmt
	Token map[string]ast.Token
}

func NewResolver(stmts []ast.Stmt) *Resolver {
	return &Resolver{Stmts: stmts, Token: make(map[string]ast.Token)}
}

func (r *Resolver) Compute() {
	for i := range r.Stmts {
		r.ComputeStmt(&r.Stmts[i])
	}
}

func (r *Resolver) ComputeStmt(stmt *ast.Stmt) {
	if letStmt, ok := (*stmt).(ast.LetStmt); ok {
		r.ComputeLetStmt(&letStmt)
		*stmt = letStmt
	}

	if constraintStmt, ok := (*stmt).(ast.ConstraintStmt); ok {
		for i := range constraintStmt.Block {
			r.ComputeStmt(&constraintStmt.Block[i])
		}
	}
}

func (r *Resolver) ComputeLetStmt(stmt *ast.LetStmt) {
	v, typ := r.ComputeExpr(stmt.Exp)
	stmt.Ident.Val = v
	stmt.Ident.Typ = typ
	r.Token[stmt.Ident.Literal] = stmt.Ident
}

func (r *Resolver) ComputeExpr(expr ast.Expr) (interface{}, ast.Type) {
	switch expr := expr.(type) {
	case ast.BinaryExpr:
		return r.ComputeBinaryExpr(expr)
	case ast.UnaryExpr:
		return r.ComputeUnaryExpr(expr)
	case ast.Token:
		tmp, typ := r.ComputeToken(expr)
		expr.Val = tmp
		return tmp, typ
	}
	return nil, ast.Any
}

func (r *Resolver) ComputeBinaryExpr(expr ast.BinaryExpr) (interface{}, ast.Type) {
	switch typ := expr.Operator.TokenType; typ {
	case ast.Plus:
		left, t := r.ComputeExpr(expr.Left)
		right, k := r.ComputeExpr(expr.Right)
		if t == ast.Integer && k == ast.Integer {
			return left.(int) + right.(int), ast.Integer
		}
		if t == ast.Float && k == ast.Float {
			return left.(float64) + right.(float64), ast.Float
		}
	case ast.Minus:
		left, t := r.ComputeExpr(expr.Left)
		right, k := r.ComputeExpr(expr.Right)
		if t == ast.Integer && k == ast.Integer {
			return left.(int) - right.(int), ast.Integer
		}
		if t == ast.Float && k == ast.Float {
			return left.(float64) - right.(float64), ast.Float
		}
	case ast.Mul:
		left, t := r.ComputeExpr(expr.Left)
		right, k := r.ComputeExpr(expr.Right)
		if t == ast.Integer && k == ast.Integer {
			return left.(int) * right.(int), ast.Integer
		}
		if t == ast.Float && k == ast.Float {
			return left.(float64) * right.(float64), ast.Float
		}
	case ast.Div:
		left, t := r.ComputeExpr(expr.Left)
		right, k := r.ComputeExpr(expr.Right)
		if t == ast.Integer && k == ast.Integer {
			return left.(int) / right.(int), ast.Integer
		}
		if t == ast.Float && k == ast.Float {
			return left.(float64) / right.(float64), ast.Float
		}
	case ast.Equal:
		left, t := r.ComputeExpr(expr.Left)
		right, k := r.ComputeExpr(expr.Right)
		if t == k {
			return left == right, ast.Bool
		}
	case ast.NotEqual:
		left, t := r.ComputeExpr(expr.Left)
		right, k := r.ComputeExpr(expr.Right)
		if t == k {
			return left != right, ast.Bool
		}
	case ast.Lt:
		left, t := r.ComputeExpr(expr.Left)
		right, k := r.ComputeExpr(expr.Right)
		if t == k {
			return left.(int) < right.(int), ast.Bool
		}
	case ast.Gt:
		left, t := r.ComputeExpr(expr.Left)
		right, k := r.ComputeExpr(expr.Right)
		if t == k {
			return left.(int) > right.(int), ast.Bool
		}
	default:
		panic("Unresolved binary operator")
	}
	r.ComputeExpr(expr.Left)
	r.ComputeExpr(expr.Right)
	return nil, ast.Any
}

func (r *Resolver) ComputeUnaryExpr(expr ast.UnaryExpr) (interface{}, ast.Type) {
	switch typ := expr.Operator.TokenType; typ {
	case ast.Plus:
		return r.ComputeExpr(expr.Right)
	case ast.Minus:
		v, exprType := r.ComputeExpr(expr.Right)
		switch exprType {
		case ast.Integer:
			return -v.(int), ast.Integer
		case ast.Float:
			return -v.(float64), ast.Float
		}
	default:
		panic("Unresolved unary operator")
	}
	return nil, ast.Any
}

func (r *Resolver) ComputeToken(token ast.Token) (interface{}, ast.Type) {
	if token.TokenType == ast.Ident {
		if v, ok := r.Token[token.Literal]; ok {
			return v.Val, v.Typ
		} else {
			return nil, ast.Any
		}
	}
	switch token.Typ {
	case ast.Integer:
		v, _ := strconv.Atoi(token.Literal)
		return v, ast.Integer
	case ast.Bool:
		return token.Literal == "true", ast.Bool
	default:
		panic("Unresolved token type")
	}
}
