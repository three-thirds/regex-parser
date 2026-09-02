package engine

import (
	"strings"
	"time"
)

type MockEngine struct{}

func NewMockEngine() Engine {
	return &MockEngine{}
}

func (m *MockEngine) Evaluate(pattern string, testString string) AnalysisResult {
	start := time.Now()

	// Simulate error when unclosed parenthesis exists
	if strings.Contains(pattern, "(") && !strings.Contains(pattern, ")") {
		return AnalysisResult{
			Pattern:    pattern,
			TestString: testString,
			SyntaxError: &SyntaxError{
				Message: "Unclosed group parenthesis",
				Offset:  len(pattern) - 1,
			},
			ExecutionTime: time.Since(start),
		}
	}

	// Mock matches if "test" or email pattern is used
	var matches []Match
	if pattern != "" && testString != "" {
		idx := strings.Index(testString, "hello")
		if idx != -1 {
			matches = append(matches, Match{
				Index: 1,
				Start: idx,
				End:   idx + 5,
				Value: "hello",
				Groups: []GroupMatch{
					{Index: 1, Name: "greeting", Start: idx, End: idx + 5, Value: "hello"},
				},
			})
		}
	}

	return AnalysisResult{
		Pattern:       pattern,
		TestString:    testString,
		Matches:       matches,
		ASTFormatted:  "Concat(\n  Literal('h'),\n  Literal('e'),\n  Literal('l'),\n  Literal('l'),\n  Literal('o')\n)",
		ExecutionTime: time.Since(start),
		SyntaxError:   nil,
	}
}
