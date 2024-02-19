package parser

import (
	"customs/ast"
	"slices"
)

type Parser struct {
	Token   []ast.Token
	current int
}

func NewParser(tokens []ast.Token) Parser {
	return Parser{Token: tokens, current: 0}
}

func (r *Parser) This() ast.Token {
	return r.Token[r.current]
}

func (r *Parser) Prev() ast.Token {
	return r.Token[r.current-1]
}

func (r *Parser) Peek() ast.Token {
	return r.Token[r.current+1]
}

func (r *Parser) Advance() {
	r.current += 1
}

func (r *Parser) IsValue() bool {
	return r.This().TokenType == ast.Value
}

func (r *Parser) IsAtEnd() bool {
	return r.TokenType() == ast.Eof
}

func (r *Parser) IsKeyword() bool {
	keywords := []ast.TokenType{ast.Let, ast.Assert, ast.Constraint, ast.Abstract, ast.Extends, ast.Is, ast.Assert}
	if slices.Contains(keywords, r.This().TokenType) {
		return true
	}
	return false
}

func (r *Parser) IsOperator() bool {
	operators := []ast.TokenType{ast.Plus, ast.Minus, ast.Multiply, ast.Divide, ast.And, ast.Or,
		ast.GreaterThan, ast.GreaterThanOrEqual, ast.LessThan, ast.LessThanOrEqual, ast.Equal, ast.NotEqual}
	if slices.Contains(operators, r.This().TokenType) {
		return true
	}
	return false
}

func (r *Parser) TokenType() ast.TokenType {
	return r.This().TokenType
}

func (r *Parser) MatchAndConsume(typ ast.TokenType) (ast.Token, bool) {
	if r.TokenType() == typ {
		token := r.This()
		r.Advance()
		return token, true
	}
	return ast.Token{}, false
}

func (r *Parser) Parse() (stmts []ast.Stmt, err error) {
	for !r.IsAtEnd() {
		switch r.TokenType() {
		case ast.Let:
			assign, ok := r.ParseAssignStmt()
			if !ok {
				err = ast.InvalidTokenErr(r.This().DebugInfo.Line, r.This().DebugInfo.Column)
			}
			stmts = append(stmts, assign)
		case ast.Abstract:
			r.Advance()
			constraint, ok := r.ParseConstraintStmt(true)
			if !ok {
				err = ast.InvalidTokenErr(r.This().DebugInfo.Line, r.This().DebugInfo.Column)
			}
			stmts = append(stmts, constraint)
		case ast.Constraint:
			constraint, ok := r.ParseConstraintStmt(false)
			if !ok {
				err = ast.InvalidTokenErr(r.This().DebugInfo.Line, r.This().DebugInfo.Column)
			}
			stmts = append(stmts, constraint)
		case ast.Assert:
			assert, ok := r.ParseAssertStmt()
			if !ok {
				err = ast.InvalidTokenErr(r.This().DebugInfo.Line, r.This().DebugInfo.Column)
			}
			stmts = append(stmts, assert)
		default:
			panic("Invalid token")
		}
	}
	return
}

func (r *Parser) ParseConstraintStmt(prefixAbstract bool) (stmt ast.ConstraintStmt, ok bool) {
	stmt.IsAbstract = prefixAbstract
	_, ok = r.MatchAndConsume(ast.Constraint)
	stmt.Id, ok = r.MatchAndConsume(ast.Ident)
	if !ok {
		panic("Expected identifier")
	}
	_, ok = r.MatchAndConsume(ast.Extends)
	if ok {
		stmt.ParentConstraint = r.This()
		r.Advance()
	}
	_, ok = r.MatchAndConsume(ast.LeftBrace)
	if !ok {
		panic("Expected left brace")
	}

	// Stmts
	for r.TokenType() == ast.Let {
		assign, _ := r.ParseAssignStmt()
		stmt.LetStmts = append(stmt.LetStmts, assign)
	}
	for r.TokenType() == ast.Assert {
		assert, _ := r.ParseAssertStmt()
		stmt.AssertStmts = append(stmt.AssertStmts, assert)
	}

	_, ok = r.MatchAndConsume(ast.RightBrace)
	if !ok {
		panic("Expected right brace")
	}
	_, ok = r.MatchAndConsume(ast.Semicolon)
	if !ok {
		panic("Expected semicolon")
	}
	return
}

func (r *Parser) ParseAssertStmt() (stmt ast.AssertStmt, ok bool) {
	_, ok = r.MatchAndConsume(ast.Assert)
	stmt.Id, ok = r.MatchAndConsume(ast.Ident)
	if !ok {
		panic("Expected identifier")
	}
	_, ok = r.MatchAndConsume(ast.As)
	if ok {
		stmt.Alias, ok = r.MatchAndConsume(ast.Ident)
		if !ok {
			panic("Expected identifier")
		}
	}
	_, ok = r.MatchAndConsume(ast.Arrow)
	if !ok {
		panic("Expected arrow")
	}
	stmt.Exprs = append(stmt.Exprs, r.ParseExpr())
	for r.TokenType() == ast.And || r.TokenType() == ast.Or {
		token := r.This()
		r.Advance()
		stmt.Exprs = append(stmt.Exprs, token)
		stmt.Exprs = append(stmt.Exprs, r.ParseExpr())
	}
	_, ok = r.MatchAndConsume(ast.Semicolon)
	if !ok {
		panic("Expected semicolon")
	}
	return

}

func (r *Parser) ParseAssignStmt() (stmt ast.AssignStmt, ok bool) {
	_, ok = r.MatchAndConsume(ast.Let)
	stmt.Id, ok = r.MatchAndConsume(ast.Ident)
	if !ok {
		panic("Expected identifier")
	}
	_, ok = r.MatchAndConsume(ast.Assign)
	if !ok {
		panic("Expected assignment")
	}
	stmt.Expr = r.ParseExpr()
	_, ok = r.MatchAndConsume(ast.Semicolon)
	if !ok {
		panic("Expected semicolon")
	}
	return
}

func (r *Parser) ParseExpr() ast.Expr {
	return r.ParseComparison()
}

func (r *Parser) ParseComparison() ast.Expr {
	left := r.ParsePlusMinus()
	for r.TokenType() == ast.GreaterThan || r.TokenType() == ast.GreaterThanOrEqual ||
		r.TokenType() == ast.LessThan || r.TokenType() == ast.LessThanOrEqual ||
		r.TokenType() == ast.Equal || r.TokenType() == ast.NotEqual {
		token := r.This()
		r.Advance()
		right := r.ParsePlusMinus()
		left = ast.BinaryExpr{Left: left, Op: token, Right: right}
	}
	return left

}

func (r *Parser) ParsePlusMinus() ast.Expr {
	left := r.ParseMultiplyDivide()
	for r.TokenType() == ast.Plus || r.TokenType() == ast.Minus {
		token := r.This()
		r.Advance()
		right := r.ParseMultiplyDivide()
		left = ast.BinaryExpr{Left: left, Op: token, Right: right}
	}
	return left
}

func (r *Parser) ParseMultiplyDivide() ast.Expr {
	left := r.ParseLogical()
	for r.TokenType() == ast.Multiply || r.TokenType() == ast.Divide {
		token := r.This()
		r.Advance()
		right := r.ParseLogical()
		left = ast.BinaryExpr{Left: left, Op: token, Right: right}
	}
	return left
}

func (r *Parser) ParseLogical() ast.Expr {
	left := r.ParseParentheses()
	for r.TokenType() == ast.And || r.TokenType() == ast.Or {
		token := r.This()
		r.Advance()
		right := r.ParseParentheses()
		left = ast.BinaryExpr{Left: left, Op: token, Right: right}
	}
	return left
}

func (r *Parser) ParseParentheses() ast.Expr {
	if r.TokenType() == ast.LeftParen {
		r.Advance()
		expr := r.ParseComparison()
		r.Advance()
		return expr
	}
	return r.ParseUnary()
}

func (r *Parser) ParseUnary() ast.Expr {
	if r.TokenType() == ast.Not || r.TokenType() == ast.Minus {
		token := r.This()
		r.Advance()
		expr := r.ParseValue()
		return ast.UnaryExpr{Op: token, Expr: expr}
	}
	return r.ParseValue()
}

func (r *Parser) ParseValue() ast.Expr {
	token := r.This()
	r.Advance()
	return token
}
