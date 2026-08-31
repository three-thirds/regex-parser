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
