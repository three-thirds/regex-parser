package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"

	"regex/ast"
	"regex/lexer"
	"regex/parser"
	"regex/tui"
	"regex/tui/engine"
)

func main() {
	// If no args or explicit "tui" command, start the interactive TUI
	if len(os.Args) == 1 || (len(os.Args) == 2 && os.Args[1] == "tui") {
		runTUI()
		return
	}

	if len(os.Args) < 3 {
		printUsage()
		os.Exit(1)
	}

	command := os.Args[1]
	pattern := os.Args[2]

	switch command {
	case "tokens":
		runTokens(pattern)

	case "parse":
		runParse(pattern)

	default:
		fmt.Printf("unknown command: %s\n", command)
		printUsage()
		os.Exit(1)
	}
}

func runTUI() {
	// Willgo will pass the real bridge here once matcher is hooked up
	eng := engine.NewMockEngine() // or engine.NewBridgeEngine()

	app := tui.NewApp(eng)
	p := tea.NewProgram(app)

	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "TUI Error: %v\n", err)
		os.Exit(1)
	}
}

func runTokens(pattern string) {
	l := lexer.New(pattern)

	for {
		token := l.NextToken()
		fmt.Println(token)
		if token.Type == lexer.TokenEOF {
			break
		}
	}
}

func runParse(pattern string) {
	p := parser.New(pattern)
	node, err := p.Parse()
	if err != nil {
		fmt.Printf("Parse error: %v\n", err)
		os.Exit(1)
	}

	fmt.Print(ast.Print(node))
}

func printUsage() {
	fmt.Println("usage:")
	fmt.Println("  regex [tui]            Launch interactive TUI")
	fmt.Println("  regex tokens <pattern> Print lexer tokens")
	fmt.Println("  regex parse <pattern>  Print AST representation")
}
