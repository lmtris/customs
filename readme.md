## Prerequisites
`customs` is a domain-specific language (DSL) for defining constraints of APIs

The specs of `customs` language is defined in this [blueprint](blueprint.md)

## About

This is a compiler for `customs` language written in Golang. It compiles the `customs` language into a `YAML` schema.

## Stages
### Define Language Specs
- [x] Define the language
### Lexer
- [x] Implement scanner
### Parser
- [x] Implement numeric expression parser
- [ ] Support parenthesis
- [x] Support boolean expression
- [ ] Support string expression
- [x] Implement statement parser
### Semantic Analysis
- [x] Variable type inference
- [x] Implement semantic analysis
  - [x] Check for duplicate identifier
  - [x] Check for undeclared identifier
  - [x] Check for type mismatch
### Code Generation
- [ ] Implement ast to targeted `yaml` file

## Contact
For any concern please contact me at this [email](mailto:lwuminhtris@gmail.com)

