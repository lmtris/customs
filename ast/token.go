package ast

import (
	"fmt"
)

type TokenType int

type LiteralType int

type Token struct {
	Literal     string
	LiteralType LiteralType
	TokenType   TokenType
	DebugInfo   DebugInfo
	Precedence  float64
}

func NewToken(tokenType TokenType, literal string, literalType LiteralType, line, column int) Token {
	return Token{
		Literal:     literal,
		LiteralType: literalType,
		TokenType:   tokenType,
		DebugInfo:   DebugInfo{Line: line, Column: column},
	}
}

func (r Token) String() string {
	return fmt.Sprintf("%s{%s %s %v}", r.TokenType, r.Literal, r.LiteralType, r.DebugInfo)
}

type DebugInfo struct {
	Line   int
	Column int
}

func (r *DebugInfo) String() string {
	return fmt.Sprintf("%d:%d", r.Line, r.Column)
}

const (
	Integer LiteralType = iota
	Float
	String
	Boolean
	Any // Undefined
)

func (r LiteralType) String() string {
	switch r {
	case Integer:
		return "Integer"
	case Float:
		return "Float"
	case String:
		return "String"
	case Boolean:
		return "Boolean"
	case Any:
		return "Any"
	}
	return "Undefined"
}

const (
	Let TokenType = iota
	Assert
	Constraint
	Abstract
	Extends
	Equal
	NotEqual
	GreaterThan
	GreaterThanOrEqual
	LessThan
	LessThanOrEqual
	Plus
	Minus
	Multiply
	Divide
	And
	Or
	LeftParen
	RightParen
	LeftBrace
	RightBrace
	As
	Is
	Not
	Assign
	Arrow
	Semicolon
	Value
	Ident
	Eof
)

func (r TokenType) String() string {
	switch r {
	case Let:
		return "Let"
	case Assert:
		return "Assert"
	case Constraint:
		return "Constraint"
	case Abstract:
		return "Abstract"
	case Extends:
		return "Extends"
	case Equal:
		return "Equal"
	case NotEqual:
		return "NotEqual"
	case GreaterThan:
		return "GreaterThan"
	case GreaterThanOrEqual:
		return "GreaterThanOrEqual"
	case LessThan:
		return "LessThan"
	case LessThanOrEqual:
		return "LessThanOrEqual"
	case Plus:
		return "Plus"
	case Minus:
		return "Minus"
	case Multiply:
		return "Multiply"
	case Divide:
		return "Divide"
	case And:
		return "And"
	case Or:
		return "Or"
	case LeftParen:
		return "LeftParen"
	case RightParen:
		return "RightParen"
	case LeftBrace:
		return "LeftBrace"
	case RightBrace:
		return "RightBrace"
	case As:
		return "As"
	case Is:
		return "Is"
	case Not:
		return "Not"
	case Assign:
		return "Assign"
	case Arrow:
		return "Arrow"
	case Semicolon:
		return "Semicolon"
	case Value:
		return "Value"
	case Ident:
		return "Ident"
	case Eof:
		return "Eof"
	}
	return "Undefined"
}
