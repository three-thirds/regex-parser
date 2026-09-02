package ast

// Print returns a formatted, human-readable representation
// of the AST rooted at the given node.

import (
	"fmt"
	"strings"
)

func Print(node Node) string {
	var builder strings.Builder

	printNode(&builder, node, 0)

	return builder.String()
}

// printNode writes a textual representation of node into builder
// indenting child nodes according to depth.
func printNode(builder *strings.Builder, node Node, depth int) {
	indent := strings.Repeat("  ", depth)

	switch n := node.(type) {
	case *Literal:
		fmt.Fprintf(builder, "%sLiteral(%c)\n", indent, n.Value)

	case *Dot:
		fmt.Fprintf(builder, "%sDot\n", indent)

	case *Concat:
		fmt.Fprintf(builder, "%sConcat\n", indent)
		printNode(builder, n.Left, depth+1)
		printNode(builder, n.Right, depth+1)
	case *Alternation:
		fmt.Fprintf(builder, "%sAlternation\n", indent)
		printNode(builder, n.Left, depth+1)
		printNode(builder, n.Right, depth+1)
	case *Star:
		fmt.Fprintf(builder, "%sStar\n", indent)
		printNode(builder, n.Expr, depth+1)
	case *Plus:
		fmt.Fprintf(builder, "%sPlus\n", indent)
		printNode(builder, n.Expr, depth+1)
	case *Question:
		fmt.Fprintf(builder, "%sQuestion\n", indent)
		printNode(builder, n.Expr, depth+1)
	default:
		fmt.Fprintf(builder, "%sUnknown\n", indent)
	}
}
