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

func (r *Parser) GetPrecedence() int {
	switch r.This().TokenType {
	case LeftParen, RightParen:
		return 0
	case Mul, Div:
		return 1
	case Plus, Minus:
		return 2
	default:
		return -1
	}
}

func (r *Parser) ParseExpr(head, tail int) (expr Expr) {
	// Case of single token
	if head == tail {
		return r.Token[head]
	}

	tmpHead, tmpTail := head, tail

	// Get most associated left and lowest precedence operator
	lowest, pos := GetPrecedence(r.Token[tail].TokenType), tail
	for head <= tail {
		if IsOperator(r.Token[head].TokenType) {
			if tmp := GetPrecedence(r.Token[head].TokenType); tmp >= lowest {
				lowest = tmp
				pos = head
			}
		}
		head += 1
	}

	isBinary := true
	// Define if binary or unary expression
	if typ := r.Token[pos].TokenType; typ == Plus || typ == Minus {
		if Prev(r, pos).TokenType == Number {
			// Binary expression
		} else {
			isBinary = false
		}
	}

	// Parse binary
	if isBinary {
		expr = BinaryExpr{
			Left:     r.ParseExpr(tmpHead, pos-1),
			Operator: r.Token[pos],
			Right:    r.ParseExpr(pos+1, tmpTail),
		}
		return
	}

	// Parse unary
	expr = UnaryExpr{
		Operator: r.Token[pos],
		Right:    r.ParseExpr(pos+1, tmpTail),
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
		return v.Operator.Literal + "(" + PrefixTraversal(v.Left) + " " + PrefixTraversal(v.Right) + ")"
	}
	return ""
}
