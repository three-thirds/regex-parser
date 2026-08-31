package lexer

import "testing"

func TestLexer(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []Token
	}{
		{
			name:  "literals",
			input: "abc",
			expected: []Token{
				{Type: TokenLiteral, Value: 'a'},
				{Type: TokenLiteral, Value: 'b'},
				{Type: TokenLiteral, Value: 'c'},
				{Type: TokenEOF},
			},
		},
		{
			name:  "operators",
			input: ".*+?",
			expected: []Token{
				{Type: TokenDot},
				{Type: TokenStar},
				{Type: TokenPlus},
				{Type: TokenQuestion},
				{Type: TokenEOF},
			},
		},
		{
			name:  "grouping and alteration",
			input: "(a|b)",
			expected: []Token{
				{Type: TokenLParen},
				{Type: TokenLiteral, Value: 'a'},
				{Type: TokenPipe},
				{Type: TokenLiteral, Value: 'b'},
				{Type: TokenRParen},
				{Type: TokenEOF},
			},
		},
		{
			name:  "complex pattern",
			input: "ab(c|d)*e+",
			expected: []Token{
				{Type: TokenLiteral, Value: 'a'},
				{Type: TokenLiteral, Value: 'b'},
				{Type: TokenLParen},
				{Type: TokenLiteral, Value: 'c'},
				{Type: TokenPipe},
				{Type: TokenLiteral, Value: 'd'},
				{Type: TokenRParen},
				{Type: TokenStar},
				{Type: TokenLiteral, Value: 'e'},
				{Type: TokenPlus},
				{Type: TokenEOF},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := New(tt.input)

			for i, expected := range tt.expected {
				actual := l.NextToken()

				if actual != expected {
					t.Fatalf(
						"token %d: expected %v, got %v",
						i,
						expected,
						actual,
					)
				}
			}
		})
	}
}
