package main

import (
	"fmt"
	"os"

	"regex/ast"
	"regex/lexer" // Importing custom lexer package!
	"regex/parser"
)

func main() {
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
	fmt.Println("  regex tokens <pattern>")
	fmt.Println("  regex parse <pattern>")
}
