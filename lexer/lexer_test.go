package lexer

import "testing"

func TestLexer(t *testing.T) { // Define TestLexer function
	tests := []struct { // Define all tests in a table structure
		name     string
		input    string
		expected []Token
	}{
		{ // test ordinary chars as literals.
			name:  "literals",
			input: "abc",
			expected: []Token{
				{Type: TokenLiteral, Value: 'a'},
				{Type: TokenLiteral, Value: 'b'},
				{Type: TokenLiteral, Value: 'c'},
				{Type: TokenEOF},
			},
		},
		{ // test recognition of operators
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
		{ // Tests grouping and alternation
			name:  "grouping and alternation",
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
		{ // Test complex patterns
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
		{ // Test an escaping star
			name:  "escaped star",
			input: `a\*b`,
			expected: []Token{
				{Type: TokenLiteral, Value: 'a'},
				{Type: TokenLiteral, Value: '*'},
				{Type: TokenLiteral, Value: 'b'},
				{Type: TokenEOF},
			},
		},
		{ // test escaping pipe
			name:  "escaped pipe",
			input: `a\|b`,
			expected: []Token{
				{Type: TokenLiteral, Value: 'a'},
				{Type: TokenLiteral, Value: '|'},
				{Type: TokenLiteral, Value: 'b'},
				{Type: TokenEOF},
			},
		},
		{ // tests escaped parentheses
			name:  "escaped parentheses",
			input: `\(abc\)`,
			expected: []Token{
				{Type: TokenLiteral, Value: '('},
				{Type: TokenLiteral, Value: 'a'},
				{Type: TokenLiteral, Value: 'b'},
				{Type: TokenLiteral, Value: 'c'},
				{Type: TokenLiteral, Value: ')'},
				{Type: TokenEOF},
			},
		},
		{ // Tests an escaped backslash
			name:  "escaped backslash",
			input: `\\`,
			expected: []Token{
				{Type: TokenLiteral, Value: '\\'},
				{Type: TokenEOF},
			},
		},
		{ // test an unsupported escape
			name:  "unsupported escape",
			input: `\d`,
			expected: []Token{
				{Type: TokenLiteral, Value: '\\'},
				{Type: TokenLiteral, Value: 'd'},
				{Type: TokenEOF},
			},
		},
		{ // tests for trailing backslash
			name:  "trailing backslash",
			input: `abc\`,
			expected: []Token{
				{Type: TokenLiteral, Value: 'a'},
				{Type: TokenLiteral, Value: 'b'},
				{Type: TokenLiteral, Value: 'c'},
				{Type: TokenLiteral, Value: '\\'},
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
