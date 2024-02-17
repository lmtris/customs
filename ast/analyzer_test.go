package ast

import (
	"testing"
)

func TestChecker_VisitLetStmt(t *testing.T) {
	input := `let x = 100;
	let y = 20;
	let z = x > y;
	let k = 10 * 100 + 30;
	`
	tokens, _ := Lexer{Text: input}.Scan()
	parser := NewParser(tokens)
	stmts := parser.Parse()
	checker := NewAnalyzer(stmts)
	checker.Check()
}

func TestChecker_VisitConstraintStmt(t *testing.T) {
	input := `
	let x = 100 > 10;
	constraint Root {
		let y = 20;
		let z = y * 3 + 4 - 5;
		assert token (t) => {
			t > z; 
		}
	}
	`
	tokens, _ := Lexer{Text: input}.Scan()
	parser := NewParser(tokens)
	stmts := parser.Parse()
	checker := NewAnalyzer(stmts)
	checker.Check()
}

func TestChecker_VisitConstraintStmt2(t *testing.T) {
	input := `
	let x = 100 > 10;
	constraint Root {
		let y = 20;
		let z = y * 3 + 4 - 5;
		assert token (z) => {
			assert name => {
				name > z;
				name < 100 < x;
			}
		}
	}
	`
	tokens, _ := Lexer{Text: input}.Scan()
	parser := NewParser(tokens)
	stmts := parser.Parse()
	checker := NewAnalyzer(stmts)
	checker.Check()
}
