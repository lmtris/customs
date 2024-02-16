package ast

import (
	"fmt"
	"testing"
)

//func TestParser_ParseExprSimple(t *testing.T) {
//	input := `10 + 2 - 4;
//	`
//	lexer := Lexer{Text: input}
//	tokens, err := lexer.Scan()
//	if err != nil {
//		t.Errorf("Error: %v", err)
//	}
//	parser := NewParser(tokens)
//	expr := parser.ParseExpr(0, len(tokens)-3)
//	fmt.Printf("input = %v", input)
//	fmt.Printf("stmts => %v\n", PrefixTraversal(expr))
//}
//
//func TestParser_ParseExprIntermediate(t *testing.T) {
//	input := `10 + 2 * 4 - a;
//	`
//	lexer := Lexer{Text: input}
//	tokens, err := lexer.Scan()
//	if err != nil {
//		t.Errorf("Error: %v", err)
//	}
//	parser := NewParser(tokens)
//	expr := parser.ParseExpr(0, len(tokens)-3)
//	fmt.Printf("input = %v", input)
//	fmt.Printf("stmts => %v\n", PrefixTraversal(expr))
//}
//
//func TestParser_ParseExprAdvance(t *testing.T) {
//	input := `2 - 4 + 10 * 8;
//	`
//	lexer := Lexer{Text: input}
//	tokens, err := lexer.Scan()
//	if err != nil {
//		t.Errorf("Error: %v", err)
//	}
//	parser := NewParser(tokens)
//	expr := parser.ParseExpr(0, len(tokens)-3)
//	fmt.Printf("input = %v", input)
//	fmt.Printf("stmts => %v\n", PrefixTraversal(expr))
//}

func TestParser_ParseLetStmt(t *testing.T) {
	input := `let a = 10 + 3 * b;
	`
	lexer := Lexer{Text: input}
	tokens, err := lexer.Scan()
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	parser := NewParser(tokens)
	stmt := parser.Parse()
	fmt.Printf("input = %v", input)
	fmt.Printf("stmts => %v\n", stmt)
}

// a > b is valid, a > b > c is invalid
func TestParser_ParseConstraintStmt(t *testing.T) {
	input := `constraint Root {
		let threshold = 10 + 2 * 3;
		assert token (t) => {
			t > threshold;
			t < 100;
		};
	}
	`
	//input := `constraint Root {
	//	let threshold = 10 + 2 * 3;
	//}
	//`
	lexer := Lexer{Text: input}
	tokens, err := lexer.Scan()
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	parser := NewParser(tokens)
	stmt := parser.Parse()
	//fmt.Printf("input = %v", input)
	for _, s := range stmt {
		fmt.Printf("stmts => %v\n", s)
	}
	//fmt.Printf("stmts => %v\n", stmts)
}

func TestParser_ParseConstraintStmt2(t *testing.T) {
	input := `constraint Root {
		let threshold = 10 + 2 * 3;
		assert token (t) => {
			t > threshold;
			t < 100;
		};
		assert usage => {
			usage == 100;
		};
		constraint Child {
			let foo = 10 + 2 * 3;
		};
	}
	`
	lexer := Lexer{Text: input}
	tokens, err := lexer.Scan()
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	parser := NewParser(tokens)
	stmt := parser.Parse()
	for _, s := range stmt {
		fmt.Printf("stmts => %v\n", s)
	}
	//fmt.Printf("stmts => %v\n", stmts)
}
