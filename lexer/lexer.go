package lexer

// Package lexer implements tokenization (lexing) of regex input.
// It converts an input string into a stream of tokens for parsing.

type Lexer struct { // Defines lexer structure, storing input regex and current position.
	input []rune // Store input as slice of runes (unicode support!!)
	pos   int
}

// New creates and returns a Lexer for the provided input string.
func New(input string) *Lexer {
	return &Lexer{
		input: []rune(input), // Starts the lexer at first character.
		pos:   0,
	}
}

// NextToken returns the next Token from the input stream. When the
// end of input is reached it returns a Token with Type TokenEOF.
func (l *Lexer) NextToken() Token {
	if l.pos >= len(l.input) {
		return Token{ // Return EOF if no chars left.
			Type: TokenEOF,
		}
	}

	ch := l.input[l.pos]
	l.pos++

	if ch == '\\' {
		return l.readEscape()
	}

	switch ch { // Match regex chars and convert them into corresponding types.
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

// readEscape handles an escape sequence starting after a backslash.
// If the backslash is the final character or the following rune is not
// escapable, it treats the backslash as a literal backslash.
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

// isEscapable reports whether a rune may follow a backslash to form
// a valid escape sequence in this lexer.
func isEscapable(ch rune) bool {
	switch ch {
	case '.', '*', '+', '?', '|', '(', ')', '\\':
		return true
	default:
		return false
	}
}
