package lexer

import "fmt" // Importing fmt for formatting literal token values.

type TokenType int // Defines TokenType as an integer based type for identifying token kinds.

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

type Token struct { // Defines a token, containing it's type and character value.
	Type  TokenType
	Value rune
}

func (t Token) String() string { // Converts a token into human readable string. (really needed ts for debugging DO NOT DELETE IT!)
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
