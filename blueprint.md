## Example
### `example.cus`
```text
abstract constraint Root {
    let threshold = 2 * (30 - 10 / 2);
    assert token (t) => {
        t > 0 and t < 100
    };
    assert usage => usage is not empty;
    assert extra_info {
        assert name (n) => n is not empty;
    };
}

constraint RegisterApi extends Root;
```
### `example.yaml`
```yaml
RegisterApi:
  Token:
    - Gt: 0
    - Lt: 100
  Usage:
    - NotEmpty: true
  ExtraInfo:
    Name:
      - NotEmpty: true
```
## Context Free Grammar
```ebnf
Program -> Constraint+

Constraint -> AbstractConstraint | ConcreteConstraint

AbstractConstraint -> 'abstract' 'constraint' Identifier '{' BlockStmt+ '}'

BlockStmt -> LetStmt | AssertStmt | NestedAssertStmt

LetStmt -> 'let' Identifier '=' Expression ';'

AssertStmt -> 'assert' Identifier '(' Identifier ')' '=>' ComplexExpression ';'
          | 'assert' Identifier '=>' SimpleExpression ';'
          
NestedAssertStmt -> 'assert' Identifier '{' AssertStmt+ | NestedAssertStmt+ '}'

ComplexExpression -> '{' Expression+ '}'
                | Expression
                
SimpleExpression -> Identifier ComparisonOperator Identifier
                | Identifier LogicalOperator Identifier
                | Identifier 'is not' 'empty'
                
ComparisonOperator -> '>' | '<'

LogicalOperator -> 'and' | 'or'

ConcreteConstraint -> 'constraint' Identifier 'extends' Identifier '{' BlockStmt* '}'

Expression -> Identifier | Number | Expression LogicalOperator Expression
          | '(' Expression ')' | Expression ComparisonOperator Expression
          
Identifier -> [a-zA-Z][a-zA-Z0-9]*

Number -> [0-9]+

```
## Notation
Only `concrete constraint` will be rendered in the generated code.

File format `*.cus`

## Warning
### W001 `implicit request definition`
This warning is shown when the constraint is not having any request definition. 
This is a warning because it is not a good practice to have a constraint without a request definition. 
It is recommended to have a request definition for the constraint.
## Error
### E001 `invalid constraint`