## Prerequisites
`customs` is a domain-specific language (DSL) for defining constraints of APIs

The specs of `customs` language is defined in this [blueprint](blueprint.md)

## This Project

This is a compiler for `customs` language written in Golang. It compiles the `customs` language into a `YAML` schema.

## Checklist
### Define Language Specs
- [x] Define the language
### Lexer
- [x] Implement lexer
### Parser
- [x] Implement numeric expression parser
- [ ] Support parenthesis
- [ ] Support boolean expression
- [ ] Support string expression
- [ ] Implement statement parser
### Static Analysis
- [ ] Implement static checker
### Code Generation
- [ ] Implement ast to targeted `yaml` file

## About
For any concern please contact me at this [email](mailto:lwuminhtris@gmail.com)

