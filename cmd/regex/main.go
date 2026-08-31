package main

import (
	"fmt"
	"os"

	"regex/lexer"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: regex <pattern>")
		os.Exit(1)
	}

	pattern := os.Args[1]

	l := lexer.New(pattern)

	for {
		token := l.NextToken()

		fmt.Println(token)

		if token.Type == lexer.TokenEOF {
			break
		}
	}
}
