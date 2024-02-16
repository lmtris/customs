package ast

import (
	"testing"
)

func TestChecker_VisitLetStmt(t *testing.T) {
	input := `let x = 100;
	let y = 20;
	let z = x > y;
	let z = 100;
	`
	tokens, _ := Lexer{Text: input}.Scan()
	parser := NewParser(tokens)
	stmts := parser.Parse()
	checker := NewChecker(stmts)
	checker.Check()
	for _, token := range checker.stack {
		t.Logf("%v", token)
	}
}
