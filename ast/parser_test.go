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
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	parser := NewParser(tokens)
	expr := parser.ParseExpr(0, len(tokens)-3)
	fmt.Printf("input = %v", input)
	fmt.Printf("stmt => %v\n", PrefixTraversal(expr))
}

func TestParser_ParseExprIntermediate(t *testing.T) {
	input := `10 + 2 * 4 - a;
	`
	lexer := Lexer{Text: input}
	tokens, err := lexer.Scan()
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	parser := NewParser(tokens)
	expr := parser.ParseExpr(0, len(tokens)-3)
	fmt.Printf("input = %v", input)
	fmt.Printf("stmt => %v\n", PrefixTraversal(expr))
}

func TestParser_ParseExprAdvance(t *testing.T) {
	input := `2 - 4 + 10 * 8;
	`
	lexer := Lexer{Text: input}
	tokens, err := lexer.Scan()
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	parser := NewParser(tokens)
	expr := parser.ParseExpr(0, len(tokens)-3)
	fmt.Printf("input = %v", input)
	fmt.Printf("stmt => %v\n", PrefixTraversal(expr))
}
