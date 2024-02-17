package engine

import (
	"customs/ast"
	analyzer2 "customs/ast/analyzer"
	"testing"
)

func TestGenerator_TestGenerateYaml(t *testing.T) {
	input := `
	let x = 10 - 2 + 3;
	constraint RegisterApi {
		let y = 10;
		assert Token (t) => {
			t > y * x;
		};
		assert Usage (u) => {
			u >= y + x;
		};
	}
	`

	scanner := ast.Lexer{Text: input}
	tokens, _ := scanner.Scan()

	parser := ast.NewParser(tokens)
	stmts := parser.Parse()

	analyzer := analyzer2.NewAnalyzer(stmts)
	analyzer.Check()

	resolver := NewResolver(analyzer.Stmts)
	resolver.Compute()

	g := NewGenerator(resolver, resolver.Stmts)
	g.GenerateYaml()

	//fmt.Printf("%v\n", g.Stmts)
	//fmt.Printf("%v\n", g.Token)
}
