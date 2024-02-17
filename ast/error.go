package ast

import "fmt"

type SemanticAnalysisError string

func getPosition(token Token) string {
	return fmt.Sprintf("%v:%v", token.Line, token.Column)
}

func TypeMismatchError(operator Token, left, right Type) SemanticAnalysisError {
	err := getPosition(operator) + fmt.Sprintf(" Type mismatch: Expected same type, got %v and %v", left, right)
	return SemanticAnalysisError(err)
}
