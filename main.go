package main

import (
	"customs/ast"
	"fmt"
)

func main() {
	input := `10 + 2;
	`
	lexer := ast.Lexer{Text: input}
	tokens, err := lexer.Scan()
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Printf("tokens => %v\n", tokens)
	//parser := ast.NewParser(tokens)
	//stmt, err := parser.ParseExpr()
	//fmt.Printf("stmt => %v\n", stmt)
}
