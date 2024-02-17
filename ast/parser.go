package ast

import (
	"fmt"
	"slices"
)

type Parser struct {
	Token   []Token
	current int
}

func NewParser(tokens []Token) Parser {
	return Parser{Token: tokens, current: 0}
}

func (r *Parser) This() Token {
	return r.Token[r.current]
}

func (r *Parser) Prev() Token {
	return r.Token[r.current-1]
}

func (r *Parser) Peek() Token {
	return r.Token[r.current+1]
}

func (r *Parser) Advance() {
	r.current += 1
}

func (r *Parser) Back() {
	r.current -= 1
}

func (r *Parser) IsAtEnd() bool {
	return r.Token[r.current].TokenType == Eof
}

func (r *Parser) IsKeyword() bool {
	keywords := []TokenType{Let, Assert, Constraint, Abstract, Extends, Is, Assert}
	if slices.Contains(keywords, r.This().TokenType) {
		return true
	}
	return false
}

func (r *Parser) IsOperator() bool {
	operators := []TokenType{Plus, Minus, Mul, Div}
	if slices.Contains(operators, r.This().TokenType) {
		return true
	}
	return false
}

func (r *Parser) MatchAndConsume(tokenType TokenType) (Token, bool) {
	if r.This().TokenType == tokenType {
		token := r.This()
		r.Advance()
		return token, true
	}
	return Token{}, false
}

func (r *Parser) Parse() (stmts []Stmt) {
	for !r.IsAtEnd() {
		switch r.This().TokenType {
		case Let:
			stmts = append(stmts, r.ParseLetStmt())
		case Assert:
			stmts = append(stmts, r.ParseAssertStmt())
		case Abstract, Constraint:
			stmts = append(stmts, r.ParseConstraintStmt())
		default:
			r.Advance()
		}
	}
	return
}

func (r *Parser) ParseLetStmt() (stmt LetStmt) {
	_, _ = r.MatchAndConsume(Let)
	ident, _ := r.MatchAndConsume(Ident)
	_, _ = r.MatchAndConsume(Equal)

	// Find the end of expression
	right := r.current
	for GetType(r, right) != Semicolon && GetType(r, right) != Eof {
		right += 1
	}

	exp := r.ParseExpr(r.current+1, right-1)
	r.current = right
	return LetStmt{Ident: ident, Exp: exp}
}

func (r *Parser) ParseAssertStmt() (stmt AssertStmt) {
	_, _ = r.MatchAndConsume(Assert)
	stmt.Ident, _ = r.MatchAndConsume(Ident)
	stmt.Alias = stmt.Ident // default alias
	_, aliased := r.MatchAndConsume(LeftParen)

	// If having alias
	if aliased {
		stmt.Alias, _ = r.MatchAndConsume(Ident)
		_, _ = r.MatchAndConsume(RightParen)
	}

	_, _ = r.MatchAndConsume(Arrow)
	_, _ = r.MatchAndConsume(LeftBrace)

	var exps []Expr
	var nested []Stmt
	for r.This().TokenType != RightBrace {

		if r.This().TokenType == Assert {
			nested = append(nested, r.ParseAssertStmt())
			continue
		}

		left, right := r.current, r.current
		for r.This().TokenType != Semicolon {
			r.Advance()
		}
		right = r.current
		expr := r.ParseExpr(left, right-1)
		exps = append(exps, expr)
		r.Advance()
	}

	stmt.Exps = exps
	stmt.NestedAsserts = nested

	return
}

func (r *Parser) ParseConstraintStmt() (stmt ConstraintStmt) {
	_, stmt.Abstract = r.MatchAndConsume(Abstract)
	_, _ = r.MatchAndConsume(Constraint)
	stmt.Ident, _ = r.MatchAndConsume(Ident)
	stmt.Extends = nil

	// Unexpected token
	if r.This().TokenType != LeftBrace {
		return ConstraintStmt{}
	}

	// Find the end of block
	var block []Token
	for balance := 1; balance != 0 && r.This().TokenType != Eof; {
		if r.This().TokenType == LeftBrace {
			balance += 1
		}

		if r.This().TokenType == RightBrace {
			balance -= 1
		}

		block = append(block, r.This())

		r.Advance()
	}

	block = block[1 : len(block)-1]
	block = append(block, NewToken(Eof, "", 0, 0))
	blockParser := NewParser(block)
	stmt.Block = blockParser.Parse()
	return
}

// ParseExpr parses an expression from a list of tokens
//
// left: the index of the first token
//
// right: the index of the last token
func (r *Parser) ParseExpr(left, right int) (expr Expr) {
	// Case of single token
	if token := r.Token[left]; left == right {
		switch token.TokenType {
		case Ident, Number:
			return token
		default:
			panic(fmt.Sprintf("Unexpected token: %v", token))
		}
	}

	tmpLeft, tmpRight := left, right

	// Get most associated left and lowest precedence operator
	lowest, pos := GetPrecedence(GetType(r, left)), left
	for left <= right {
		typ := GetType(r, left)
		// Normal operators
		if IsOperator(typ) {
			if tmp := GetPrecedence(typ); tmp >= lowest {
				lowest = tmp
				pos = left
			}
		}

		left += 1
	}

	isBinary := true
	// Define if binary or unary expression
	if typ := GetType(r, pos); typ == Plus || typ == Minus {
		if Prev(r, pos).TokenType == Number || Prev(r, pos).TokenType == Ident {
			// Binary expression
		} else {
			isBinary = false
		}
	}

	// Parse binary
	if isBinary {
		expr = BinaryExpr{
			Left:     r.ParseExpr(tmpLeft, pos-1),
			Operator: r.Token[pos],
			Right:    r.ParseExpr(pos+1, tmpRight),
		}
		return
	}

	// Parse unary
	expr = UnaryExpr{
		Operator: r.Token[pos],
		Right:    r.ParseExpr(pos+1, tmpRight),
	}

	return
}

func PrefixTraversal(expr Expr) string {
	switch v := expr.(type) {
	case Token:
		return v.Literal
	case UnaryExpr:
		return v.Operator.Literal + "(" + PrefixTraversal(v.Right) + ")"
	case BinaryExpr:
		return "(" + v.Operator.Literal + " " + PrefixTraversal(v.Left) + " " + PrefixTraversal(v.Right) + ")"
	}
	return ""
}
