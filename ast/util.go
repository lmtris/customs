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
	case Gt, Gte, Lt, Lte:
		return 3
	default:
		return -1
	}
}

func IsOperator(typ TokenType) bool {
	operators := []TokenType{Plus, Minus, Mul, Div, Gt, Gte, Lt, Lte, Equal, NotEqual}
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

func GetType(r *Parser, pos int) TokenType {
	return r.Token[pos].TokenType
}
