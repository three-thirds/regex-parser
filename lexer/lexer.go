package lexer

type Lexer struct {
	input []rune
	pos   int
}

func New(input string) *Lexer {
	return &Lexer{
		input: []rune(input),
		pos:   0,
	}
}

func (l *Lexer) NextToken() Token {
	if l.pos >= len(l.input) {
		return Token{
			Type: TokenEOF,
		}
	}

	ch := l.input[l.pos]
	l.pos++

	if ch == '\\' {
		return l.readEscape()
	}

	switch ch {
	case '.':
		return Token{Type: TokenDot}
	case '*':
		return Token{Type: TokenStar}
	case '+':
		return Token{Type: TokenPlus}
	case '?':
		return Token{Type: TokenQuestion}
	case '|':
		return Token{Type: TokenPipe}
	case '(':
		return Token{Type: TokenLParen}
	case ')':
		return Token{Type: TokenRParen}
	default:
		return Token{
			Type:  TokenLiteral,
			Value: ch,
		}
	}
}

func (l *Lexer) readEscape() Token {
	if l.pos >= len(l.input) {
		return Token{
			Type:  TokenLiteral,
			Value: '\\',
		}
	}

	next := l.input[l.pos]

	if !isEscapable(next) {
		return Token{
			Type:  TokenLiteral,
			Value: '\\',
		}
	}

	l.pos++

	return Token{
		Type:  TokenLiteral,
		Value: next,
	}
}

func isEscapable(ch rune) bool {
	switch ch {
	case '.', '*', '+', '?', '|', '(', ')', '\\':
		return true
	default:
		return false
	}
}
