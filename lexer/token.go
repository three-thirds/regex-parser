package lexer

import "fmt"

type TokenType int

const (
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

type Token struct {
	Type  TokenType
	Value rune
}

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
