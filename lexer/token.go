package lexer

// Package-level comments: Token types and Token values used by the lexer.

import "fmt" // Importing fmt for formatting literal token values.

type TokenType int // TokenType enumerates the different kinds of tokens emitted by the lexer.

const ( // Declaring all token types. (supported by lexer)
	TokenLiteral TokenType = iota
	TokenDot
	TokenStar
	TokenPlus
	TokenQuestion
	TokenPipe
	TokenLParen
	TokenRParen
	TokenEOF
)

// Token represents a single token produced by the lexer.
type Token struct {
	Type  TokenType
	Value rune
}

// String returns a human-readable representation of t. Useful for debugging.
func (t Token) String() string {
	switch t.Type {
	case TokenLiteral:
		return fmt.Sprintf("LITERAL(%c)", t.Value)
	case TokenDot:
		return "DOT"
	case TokenStar:
		return "STAR"
	case TokenPlus:
		return "PLUS"
	case TokenQuestion:
		return "QUESTION"
	case TokenPipe:
		return "PIPE"
	case TokenLParen:
		return "LPAREN"
	case TokenRParen:
		return "RPAREN"
	case TokenEOF:
		return "EOF"
	default:
		return "UNKNOWN"
	}
}
