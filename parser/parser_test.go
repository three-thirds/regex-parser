package parser

import (
	"reflect"
	"testing"

	"regex/ast"
)

func TestParserPrimary(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected ast.Node
	}{
		{
			name:  "literal",
			input: "a",
			expected: &ast.Literal{
				Value: 'a',
			},
		},
		{
			name:     "dot",
			input:    ".",
			expected: &ast.Dot{},
		},
		{
			name:  "grouped literal",
			input: "(a)",
			expected: &ast.Literal{
				Value: 'a',
			},
		},
		{
			name:  "nested groups",
			input: "((a))",
			expected: &ast.Literal{
				Value: 'a',
			},
		},
		{
			name:  "escaped metacharacter",
			input: `\*`,
			expected: &ast.Literal{
				Value: '*',
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := New(tt.input)

			actual, err := p.Parse()
			if err != nil {
				t.Fatalf("unexpected parser error: %v", err)
			}

			if !reflect.DeepEqual(actual, tt.expected) {
				t.Fatalf("expected %#v, got %#v", tt.expected, actual)
			}
		})
	}
}

func TestParserErrors(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{name: "empty input", input: ""},
		{name: "empty group", input: "()"},
		{name: "missing closing parenthesis", input: "(a"},
		{name: "unexpected closing parenthesis", input: "a)"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := New(tt.input)

			_, err := p.Parse()
			if err == nil {
				t.Fatal("expected parse error, got nil")
			}
		})
	}
}

func TestParsedQuantifiers(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected ast.Node
	}{
		{
			name:  "star",
			input: "a*",
			expected: &ast.Star{
				Expr: &ast.Literal{Value: 'a'},
			},
		},
		{
			name:  "plus",
			input: "a+",
			expected: &ast.Plus{
				Expr: &ast.Literal{Value: 'a'},
			},
		},
		{
			name:  "question",
			input: "a?",
			expected: &ast.Question{
				Expr: &ast.Literal{Value: 'a'},
			},
		},
		{
			name:  "grouped star",
			input: "(a)*",
			expected: &ast.Star{
				Expr: &ast.Literal{Value: 'a'},
			},
		},
		{
			name:  "dot star",
			input: ".*",
			expected: &ast.Star{
				Expr: &ast.Dot{},
			},
		},
		{
			name:     "escaped star stays literal",
			input:    `\*`,
			expected: &ast.Literal{Value: '*'},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := New(tt.input)

			actual, err := p.Parse()
			if err != nil {
				t.Fatalf("unexpected parse error: %v", err)
			}

			if !reflect.DeepEqual(actual, tt.expected) {
				t.Fatalf("expected %#v, got %#v", tt.expected, actual)
			}
		})
	}
}
