package scanner

import (
	"fmt"
	"testing"
)

func TestLexer_Scan(t *testing.T) {
	input := `x = 10.5`
	lexer := NewLexer(input)
	err := lexer.Scan()
	if err != nil {
		fmt.Printf("Error = %v\n", err)
	}
	fmt.Printf("%v\n", lexer.Tokens)
}

func TestLexer_Scan2(t *testing.T) {
	input := `x = "Hello World!"`
	lexer := NewLexer(input)
	err := lexer.Scan()
	if err != nil {
		fmt.Printf("Error = %v\n", err)
	}
	fmt.Printf("%v\n", lexer.Tokens)
}
