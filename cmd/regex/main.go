package main

import (
	"fmt"
	"os"

	"regex/lexer" // Importing custom lexer package!
)

func main() { // Define program's entry point
	if len(os.Args) != 2 {
		fmt.Println("usage: regex <pattern>")
		os.Exit(1) // Exit with non zero if any error comes.
	}

	pattern := os.Args[1]

	l := lexer.New(pattern) // Create a new lexer for any given lexer pattern.

	for {
		token := l.NextToken() // Read next token from Lexer

		fmt.Println(token) // Print token using String()

		if token.Type == lexer.TokenEOF { // break if EOF
			break
		}
	}
}
