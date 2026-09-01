package lexer

type Lexer struct { // Defines lexer structure, storing input regex and current position.
	input []rune // Store input as slice of runes (unicode support!!)
	pos   int
}

func New(input string) *Lexer { // Creates and initializes a new Lexer from the given input string.
	return &Lexer{
		input: []rune(input), // Starts the lexer at first character.
		pos:   0,
	}
}

// Return the next token from the input.
func (l *Lexer) NextToken() Token { // Check whether lexer reached end of input.
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

func (l *Lexer) readEscape() Token { // Check whether backslash is final character in input.
	if l.pos >= len(l.input) {
		return Token{ // Treat trailing backslash as a literal backslash.
			Type:  TokenLiteral,
			Value: '\\',
		}
	}

	next := l.input[l.pos]

	if !isEscapable(next) { // Checks whether a character can be escaped by the lexer.
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

func isEscapable(ch rune) bool { // Determines all characters which can be escaped by the lexer.
	switch ch {
	case '.', '*', '+', '?', '|', '(', ')', '\\':
		return true
	default:
		return false
	}
}
