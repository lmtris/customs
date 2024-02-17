package engine

import (
	"customs/ast"
	analyzer2 "customs/ast/analyzer"
	"fmt"
	"testing"
)

func TestResolver_TestCompute(t *testing.T) {
	input := `
	let x = 10 - 2 + 3;
	constraint RegisterApi {
		let y = 10;
		assert token (t) => {
			t > y;
		}
	}
	`

	scanner := ast.Lexer{Text: input}
	tokens, _ := scanner.Scan()

	parser := ast.NewParser(tokens)
	stmts := parser.Parse()

	analyzer := analyzer2.NewAnalyzer(stmts)
	analyzer.Check()

	g := NewResolver(analyzer.Stmts)
	g.Compute()
	fmt.Printf("%v\n", g.Stmts)
	fmt.Printf("%v\n", g.Token)
}
