// Package engine defines contract for TUI and regex, to be implemented by @Willgob
package engine

import "time"

// GroupMatch represents a capturing group match
// ($1, $2, or named)
type GroupMatch struct {
	Index int
	Name  string
	Start int
	End   int
	Value string
}

// Match represents a single full match in
// the test string
type Match struct {
	Index  int
	Start  int
	End    int
	Value  string
	Groups []GroupMatch
}

// AnalysisResult is what the TUI receives after regex evaluation
type AnalysisResult struct {
	Pattern       string
	TestString    string
	Matches       []Match
	ASTFormatted  string        // Output from ast.Print
	ExecutionTime time.Duration // Engine run time
	SyntaxError   *SyntaxError  // Nil if pattern is valid
}

type SyntaxError struct {
	Message string
	Offset  int // Character index where the syntax broke
}

// Engine defines what the TUI needs from Dev's engine
type Engine interface {
	Evaluate(pattern string, testString string) AnalysisResult
}
