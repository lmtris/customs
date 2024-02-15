package ast

import (
	"slices"
)

type Parser struct {
	Token   []Token
	current int
	head    int
	tail    int
}

func NewParser(tokens []Token) Parser {
	return Parser{Token: tokens, current: 0, head: 0, tail: len(tokens)}
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

// RemoveParenthesis removes all parenthesis from Parenthesis+ Token
func RemoveParenthesis(tokens []Token) (num Token, ok bool) {
	typs := []TokenType{LeftParen, RightParen, Number}

	// Case check valid token types
	for token := range tokens {
		if !slices.Contains(typs, tokens[token].TokenType) {
			return Token{}, false
		}
	}

	// Check only one number
	count := 0
	for token := range tokens {
		if tokens[token].TokenType == Number {
			num = tokens[token]
			count += 1
		}
	}

	if count != 1 {
		return Token{}, false
	}

	return num, true
}

// ParseExpr parses an expression from a list of tokens
//
// head: the index of the first token
//
// tail: the index of the last token
//
// level: the level of the expression, incremented at each inner parenthesis
func (r *Parser) ParseExpr(left, right int) (expr Expr) {
	// Case of single token
	if left == right {
		return r.Token[left]
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
		if Prev(r, pos).TokenType == Number {
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
