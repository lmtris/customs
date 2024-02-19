package parser

import (
	"customs/ast/scanner"
	"fmt"
	"testing"
)

func TestParser_Parse(t *testing.T) {
	input := `let x = 1 + 2 * (3 - 4);`
	lexer := scanner.NewLexer(input)
	err := lexer.Scan()
	if err != nil {
		t.Errorf("Error = %v\n", err)
	}
	parser := NewParser(lexer.Tokens)
	stmts, _ := parser.Parse()
	for _, stmt := range stmts {
		fmt.Printf("%s\n", stmt)
	}
	//fmt.Printf("%s\n", PrefixTraversal(expr))
	if err != nil {
		t.Errorf("Error = %v\n", err)
	}
}
