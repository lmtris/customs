package ast

import (
	"fmt"
	"testing"
)

func TestParser_ParseExprSimple(t *testing.T) {
	input := `10 + 2 - 4;
	`
	lexer := Lexer{Text: input}
	tokens, err := lexer.Scan()
	fmt.Printf("tokens => %v\n", tokens)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	parser := NewParser(tokens)
	expr := parser.ParseExpr(0, len(tokens)-3)
	fmt.Printf("stmt => %v\n", PrefixTraversal(expr))
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if expr == nil {
		t.Errorf("stmt is nil")
	}
}

func TestParser_ParseExprIntermediate(t *testing.T) {
	input := `10 + 2 * 4 - a;
	`
	lexer := Lexer{Text: input}
	tokens, err := lexer.Scan()
	fmt.Printf("tokens => %v\n", tokens)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	parser := NewParser(tokens)
	expr := parser.ParseExpr(0, len(tokens)-3)
	fmt.Printf("stmt => %v\n", PrefixTraversal(expr))
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if expr == nil {
		t.Errorf("stmt is nil")
	}
}
