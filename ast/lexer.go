package ast

type Lexer struct {
	Text string
}

type UndefinedToken error

func (r Lexer) Scan() (tokens []Token, err error) {
	chars := []rune(r.Text)
	line, column := 1, 1
	index := 0
	for !r.isEof(index) {
		//fmt.Printf("%v => %s\n", current, string(chars[current]))
		word := chars[index]

		/// Break line token
		if word == '\n' {
			line += 1
			column = 1
			index += 1
			continue
		}

		/// Scan tokens
		switch word {
		case '+':
			tokens = append(tokens, NewToken(Plus, "+", line, column))
		case '-':
			tokens = append(tokens, NewToken(Minus, "-", line, column))
		case '*':
			tokens = append(tokens, NewToken(Mul, "*", line, column))
		case '/':
			tokens = append(tokens, NewToken(Div, "/", line, column))
		case ',':
			tokens = append(tokens, NewToken(Comma, ",", line, column))
		case ';':
			tokens = append(tokens, NewToken(Semicolon, ";", line, column))
		case '{':
			tokens = append(tokens, NewToken(LeftBrace, "{", line, column))
		case '}':
			tokens = append(tokens, NewToken(RightBrace, "}", line, column))
		case '(':
			tokens = append(tokens, NewToken(LeftParen, "(", line, column))
		case ')':
			tokens = append(tokens, NewToken(RightParen, ")", line, column))
		case '=':
			if r.peek(index+1) == '=' {
				tokens = append(tokens, NewToken(Equal, "==", line, column))
				index += 1
				break
			}
			tokens = append(tokens, NewToken(Assign, "=", line, column))
		case '>':
			if r.peek(index+1) == '=' {
				tokens = append(tokens, NewToken(Gte, ">=", line, column))
				index += 1
				break
			}
			tokens = append(tokens, NewToken(Gt, ">", line, column))
		case '<':
			if r.peek(index+1) == '=' {
				tokens = append(tokens, NewToken(Lte, "<=", line, column))
				index += 1
				break
			}
			tokens = append(tokens, NewToken(Lt, "<", line, column))
		default:
			/// Num
			if r.isDigit(word) {
				num := []rune{word}
				for !r.isEof(index+1) && r.isDigit(r.peek(index+1)) {
					num = append(num, r.peek(index+1))
					index += 1
				}
				tokens = append(tokens, NewToken(Number, string(num), line, column))
			}
			/// Letter
			if r.isLetter(word) {
				letter := []rune{word}
				for !r.isEof(index+1) && r.isLetter(r.peek(index+1)) {
					letter = append(letter, r.peek(index+1))
					index += 1
				}
				switch string(letter) {
				case "let":
					tokens = append(tokens, NewToken(Let, "let", line, column))
				case "abstract":
					tokens = append(tokens, NewToken(Abstract, "abstract", line, column))
				case "constraint":
					tokens = append(tokens, NewToken(Constraint, "constraint", line, column))
				case "assert":
					tokens = append(tokens, NewToken(Assert, "assert", line, column))
				case "extends":
					tokens = append(tokens, NewToken(Extends, "extends", line, column))
				case "is":
					tokens = append(tokens, NewToken(Is, "is", line, column))
				case "not":
					tokens = append(tokens, NewToken(Not, "not", line, column))
				default:
					tokens = append(tokens, NewToken(Ident, string(letter), line, column))
				}
			}

		}

		index += 1
		column += 1
	}
	tokens = append(tokens, NewToken(Eof, "Eof", line, column))
	//fmt.Printf("%v\n", tokens)
	return
}

func (r Lexer) isDigit(ch rune) bool {
	return ch >= '0' && ch <= '9'
}

func (r Lexer) peek(index int) rune {
	if r.isEof(index) {
		return '\x00'
	}
	return rune(r.Text[index])
}

func (r Lexer) isEof(index int) bool {
	return index >= len(r.Text)
}

func (r Lexer) isLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || ch == '_'
}
