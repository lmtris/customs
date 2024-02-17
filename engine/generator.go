package engine

import (
	"customs/ast"
	"gopkg.in/yaml.v2"
	"os"
)

type Yaml struct {
	Data map[string]map[string][]map[string]int `yaml:",inline"`
}

type Generator struct {
	Resolver *Resolver
	Stmts    []ast.Stmt
}

func NewGenerator(resolver *Resolver, stmts []ast.Stmt) Generator {
	return Generator{Resolver: resolver, Stmts: stmts}
}

func (r *Generator) GenerateYaml() {
	y := Yaml{Data: make(map[string]map[string][]map[string]int)}

	for _, stmt := range r.Stmts {
		switch stmt := stmt.(type) {
		case ast.ConstraintStmt:
			y.Data[stmt.Ident.Literal] = make(map[string][]map[string]int)
			for _, blockStmt := range stmt.Block {
				if assertStmt, ok := blockStmt.(ast.AssertStmt); ok {
					for _, exp := range assertStmt.Exps {
						if binaryExp, ok := exp.(ast.BinaryExpr); ok {
							right, typ := r.Resolver.ComputeExpr(binaryExp.Right)
							if typ == ast.Integer {
								y.Data[stmt.Ident.Literal][assertStmt.Ident.Literal] =
									append(
										y.Data[stmt.Ident.Literal][assertStmt.Ident.Literal],
										map[string]int{binaryExp.Operator.Literal: right.(int)})
							}
						}
					}
				}
			}
		}
	}

	d, err := yaml.Marshal(&y)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("output.yaml", d, 0644)
	if err != nil {
		panic(err)
	}
}
