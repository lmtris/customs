package analyzer

import (
	parser2 "customs/ast/parser"
	"customs/ast/scanner"
	"fmt"
	"testing"
)

func TestAnalyzer_VisitAssignStmt(t *testing.T) {
	input := `let x = 2;
	let y = (1 + x) * 3;`
	lexer := scanner.NewLexer(input)
	err := lexer.Scan()
	if err != nil {
		t.Errorf("Error = %v\n", err)
	}
	parser := parser2.NewParser(lexer.Tokens)
	stmts, _ := parser.Parse()
	for _, stmt := range stmts {
		fmt.Printf("%s\n", stmt)
	}
	if err != nil {
		t.Errorf("Error = %v\n", err)
	}
	analyzer2 := NewAnalyzer(stmts)
	analyzer2.Analyze()
}
