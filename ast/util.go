package ast

import (
	"slices"
)

func GetPrecedence(typ TokenType) int {
	switch typ {
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

func IsOperator(typ TokenType) bool {
	operators := []TokenType{Plus, Minus, Mul, Div}
	if slices.Contains(operators, typ) {
		return true
	}
	return false
}

func Prev(r *Parser, pos int) Token {
	if pos == 0 {
		return NewToken(Eof, "BEGIN", -1, -1)
	}
	return r.Token[pos-1]
}
