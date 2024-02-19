## Prerequisites
`customs` is a domain-specific language (DSL) for defining constraints of APIs

The specs of `customs` language is defined in this [blueprint](blueprint.md)

## About

This is a compiler for `customs` language written in Golang. It compiles the `customs` language into a `YAML` schema.

## Stages
- ``18/02/2024`` first release
- ``19/02/2024`` major refactoring

### Define Language Specs
- [x] Define the language
### Lexer
- [x] Define tokens
  - [x] Support `float`, `string`, `integer`, `boolean`
- [x] Implement scanner
### Parser
- [x] Implement sane error handler


- [x] Implement expression parser 
  - [x] Parse unary expression
  - [x] Parse binary expression
  - [x] Parse token expression
  - [x] Parse boolean expression
  - [x] Support parenthesis
  - [ ] Support string expression


- [x] Implement statement parser
  - [x] Implement constraint statement
  - [x] Implement let statement
  - [x] Implement assert statement
### Semantic Analysis
- [x] Variable type inference


- [x] Implement semantic analysis
  - [x] Check for duplicate identifier
  - [x] Check for undeclared identifier
  - [x] Check for type mismatch
  - [ ] Check assert expression cross constraints
### Code Generation
- [x] Implement ast to targeted `yaml` file
## Todo
### Minor changes
- Support `float`, `string` and `boolean`
- Support `parenthesis`
- Implement error handler, quick return error when error is found
- Restrict only `assert` statements are allowed to be nested in other `assert` statements
- Refactor the codebase

### Major changes
- Supports inheritance (abstract and extends)
- Implement WASM for this compiler
## Contact
For any concern please contact me at this [email](mailto:lwuminhtris@gmail.com)

