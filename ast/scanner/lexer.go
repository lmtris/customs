package scanner

import v2 "customs/ast"

type Lexer struct {
	Text    string
	Tokens  []v2.Token
	current int
}

func NewLexer(text string) *Lexer {
	return &Lexer{Text: text, Tokens: make([]v2.Token, 0), current: 0}
}

func (r *Lexer) IsEof() bool {
	return r.current >= len(r.Text)
}

func (r *Lexer) This() rune {
	if r.IsEof() {
		return rune(0)
	}
	return rune(r.Text[r.current])
}

func (r *Lexer) Peek() rune {
	if r.IsEof() {
		return rune(0)
	}
	return rune(r.Text[r.current+1])
}

func (r *Lexer) Advance() {
	r.current += 1
}

func (r *Lexer) IsDigit() bool {
	return r.This() >= '0' && r.This() <= '9'
}

func (r *Lexer) IsLetter() bool {
	return (r.This() >= 'a' && r.This() <= 'z') || (r.This() >= 'A' && r.This() <= 'Z') || r.This() == '_'
}

func (r *Lexer) IsString() bool {
	return r.This() == '"'
}

func (r *Lexer) Scan() (err error) {
	var line, column = 1, 1
	for !r.IsEof() {
		switch r.This() {
		case '\n':
			line = line + 1
			column = 1
			r.Advance()
			continue
		case ' ':
			r.Advance()
			continue
		case '+':
			r.Tokens = append(r.Tokens, v2.NewToken(v2.Plus, "+", v2.Any, line, column))
		case '-':
			r.Tokens = append(r.Tokens, v2.NewToken(v2.Minus, "-", v2.Any, line, column))
		case '*':
			r.Tokens = append(r.Tokens, v2.NewToken(v2.Multiply, "*", v2.Any, line, column))
		case '/':
			r.Tokens = append(r.Tokens, v2.NewToken(v2.Divide, "/", v2.Any, line, column))
		case '(':
			r.Tokens = append(r.Tokens, v2.NewToken(v2.LeftParen, "(", v2.Any, line, column))
		case ')':
			r.Tokens = append(r.Tokens, v2.NewToken(v2.RightParen, ")", v2.Any, line, column))
		case '{':
			r.Tokens = append(r.Tokens, v2.NewToken(v2.LeftBrace, "{", v2.Any, line, column))
		case '}':
			r.Tokens = append(r.Tokens, v2.NewToken(v2.RightBrace, "}", v2.Any, line, column))
		case ';':
			r.Tokens = append(r.Tokens, v2.NewToken(v2.Semicolon, ";", v2.Any, line, column))
		case '=':
			switch r.Peek() {
			case '=':
				r.Tokens = append(r.Tokens, v2.NewToken(v2.Equal, "==", v2.Any, line, column))
				r.Advance()
			case '>':
				r.Tokens = append(r.Tokens, v2.NewToken(v2.Arrow, "=>", v2.Any, line, column))
				r.Advance()
			case ' ':
				r.Tokens = append(r.Tokens, v2.NewToken(v2.Assign, "=", v2.Any, line, column))
				r.Advance()
			default:
				err = v2.InvalidTokenErr(line, column)
				return
			}
		case '>':
			switch r.Peek() {
			case '=':
				r.Tokens = append(r.Tokens, v2.NewToken(v2.GreaterThanOrEqual, ">=", v2.Any, line, column))
				r.Advance()
			case ' ':
				r.Tokens = append(r.Tokens, v2.NewToken(v2.GreaterThan, ">", v2.Any, line, column))
				r.Advance()
			default:
				err = v2.InvalidTokenErr(line, column)
				return
			}
		case '<':
			switch r.Peek() {
			case '=':
				r.Tokens = append(r.Tokens, v2.NewToken(v2.LessThanOrEqual, "<=", v2.Any, line, column))
				r.Advance()
			case ' ':
				r.Tokens = append(r.Tokens, v2.NewToken(v2.LessThan, "<", v2.Any, line, column))
				r.Advance()
			default:
				err = v2.InvalidTokenErr(line, column)
				return
			}
		default:
			if r.IsDigit() {
				start := r.current
				dots := false
				for (r.IsDigit() || r.This() == '.') && !r.IsEof() {
					if r.This() == '.' {
						dots = true
					}
					r.Advance()
				}
				if dots {
					r.Tokens = append(r.Tokens, v2.NewToken(v2.Value, r.Text[start:r.current], v2.Float, line, column))
				} else {
					r.Tokens = append(r.Tokens, v2.NewToken(v2.Value, r.Text[start:r.current], v2.Integer, line, column))
				}
				continue
			}
			if r.IsLetter() {
				start := r.current
				for r.IsLetter() && !r.IsEof() {
					r.Advance()
				}
				switch txt := r.Text[start:r.current]; txt {
				case "let":
					r.Tokens = append(r.Tokens, v2.NewToken(v2.Let, txt, v2.Any, line, column))
				case "assert":
					r.Tokens = append(r.Tokens, v2.NewToken(v2.Assert, txt, v2.Any, line, column))
				case "constraint":
					r.Tokens = append(r.Tokens, v2.NewToken(v2.Constraint, txt, v2.Any, line, column))
				case "abstract":
					r.Tokens = append(r.Tokens, v2.NewToken(v2.Abstract, txt, v2.Any, line, column))
				case "is":
					r.Tokens = append(r.Tokens, v2.NewToken(v2.Is, txt, v2.Any, line, column))
				case "extends":
					r.Tokens = append(r.Tokens, v2.NewToken(v2.Extends, txt, v2.Any, line, column))
				case "as":
					r.Tokens = append(r.Tokens, v2.NewToken(v2.As, txt, v2.Any, line, column))
				case "not":
					r.Tokens = append(r.Tokens, v2.NewToken(v2.Not, txt, v2.Any, line, column))
				default:
					r.Tokens = append(r.Tokens, v2.NewToken(v2.Ident, txt, v2.Any, line, column))
				}
				continue
			}
			if r.IsString() {
				start := r.current
				r.Advance()
				for !r.IsString() && !r.IsEof() {
					r.Advance()
				}
				r.Tokens = append(r.Tokens, v2.NewToken(v2.Value, r.Text[start:r.current+1], v2.String, line, column))
				continue
			}
		}
		column = column + 1
		r.Advance()
	}
	r.Tokens = append(r.Tokens, v2.NewToken(v2.Eof, "Eof", v2.Any, line, column))
	return
}
