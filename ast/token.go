package ast

import "fmt"

type Token struct {
	TokenType
	Literal string
	Typ     Type
	Line    int
	Column  int
}

type Type string

const (
	Integer Type = "int"
	Bool    Type = "bool"
	Float   Type = "float"
	String  Type = "string"
	Any     Type = "any"
)

func NewToken(typ TokenType, lit string, line, column int) Token {
	return Token{TokenType: typ, Literal: lit, Line: line, Column: column}
}

// Precedence Order Table
//
// ()
// * /
// + -
// > >= < <=
// == !=
// =
// (a + b) * c > 2 == False -> ((a + b) * c) > 2) == False

func NewOperator(typ TokenType, lit string, line, column, left, right int) Token {
	return Token{
		TokenType: typ,
		Literal:   lit,
		Line:      line,
		Column:    column,
	}
}

func (r Token) String() string {
	return fmt.Sprintf("<%s %s %s>", r.TokenType, r.Literal, r.Typ)
}

type TokenType int

const (
	Let TokenType = iota
	Arrow
	Abstract
	Constraint
	Assert
	Extends
	Is
	Not
	Plus
	Minus
	Mul
	Div
	Comma
	Semicolon
	LeftBrace
	RightBrace
	LeftParen
	RightParen
	Assign
	Gt
	Lt
	Gte
	Lte
	Equal
	NotEqual
	Number
	Ident
	Eof
)

func (r TokenType) String() string {
	switch r {
	case Let:
		return "LET"
	case Abstract:
		return "ABSTRACT"
	case Constraint:
		return "CONSTRAINT"
	case Assert:
		return "ASSERT"
	case Extends:
		return "EXTENDS"
	case Is:
		return "IS"
	case Not:
		return "NOT"
	case Plus:
		return "PLUS"
	case Minus:
		return "MINUS"
	case Mul:
		return "MUL"
	case Div:
		return "SLASH"
	case Comma:
		return "COMMA"
	case Semicolon:
		return "SEMICOLON"
	case LeftBrace:
		return "LEFT_BRACE"
	case RightBrace:
		return "RIGHT_BRACE"
	case LeftParen:
		return "LEFT_PAREN"
	case RightParen:
		return "RIGHT_PAREN"
	case Assign:
		return "ASSIGN"
	case Gt:
		return "GREATER"
	case Lt:
		return "LESS"
	case Gte:
		return "GREATER_EQUAL"
	case Lte:
		return "LESS_EQUAL"
	case Equal:
		return "EQUAL"
	case NotEqual:
		return "NOT_EQUAL"
	case Number:
		return "NUM"
	case Ident:
		return "ID"
	case Eof:
		return "EOF"
	}
	return "UNDEFINED"
}
