package nfa

// Package nfa contains a Compiler that converts AST nodes into NFA
// fragments used during regex compilation.

import (
	"fmt"

	"regex/ast"
)

type Compiler struct {
	builder *Builder
}

type fragment struct {
	start  *State
	accept *State
}

func NewCompiler() *Compiler {
	return &Compiler{
		builder: NewBuilder(),
	}
}

// Compile converts the given AST node into an NFA. It returns the
// constructed NFA or an error if the AST contains unsupported nodes.
func (c *Compiler) Compile(node ast.Node) (*NFA, error) {
	frag, err := c.compileNode(node)
	if err != nil {
		return nil, err
	}

	return &NFA{
		Start:  frag.start,
		Accept: frag.accept,
	}, nil
}

// compileNode dispatches compilation based on the concrete AST node type.
func (c *Compiler) compileNode(node ast.Node) (fragment, error) {
	switch n := node.(type) {
	case *ast.Literal:
		return c.compileLiteral(n), nil

	case *ast.Dot:
		return c.compileDot(), nil

	case *ast.Concat:
		return c.compileConcat(n)

	case *ast.Alternation:
		return c.compileAlternation(n)

	default:
		return fragment{}, fmt.Errorf("unsupported AST node: %T", node)
	}
}

// compileLiteral creates a fragment representing a single literal rune.
func (c *Compiler) compileLiteral(node *ast.Literal) fragment {
	start := c.builder.NewState()
	accept := c.builder.NewState()

	AddLiteral(start, accept, node.Value)

	return fragment{
		start:  start,
		accept: accept,
	}
}

// compileDot creates a fragment that matches any single rune (dot).
func (c *Compiler) compileDot() fragment {
	start := c.builder.NewState()
	accept := c.builder.NewState()

	AddAny(start, accept)

	return fragment{
		start:  start,
		accept: accept,
	}
}

// compileConcat combines two fragments by connecting the accept state
// of the left fragment to the start state of the right fragment with
// an epsilon transition, producing a new fragment that represents
// the concatenation of the two subexpressions.
func (c *Compiler) compileConcat(node *ast.Concat) (fragment, error) {
	left, err := c.compileNode(node.Left)
	if err != nil {
		return fragment{}, err
	}

	right, err := c.compileNode(node.Right)
	if err != nil {
		return fragment{}, err
	}

	// Use the builder helper to add an epsilon transition between
	// the two sub-fragments.
	AddEpsilon(left.accept, right.start)

	return fragment{
		start:  left.start,
		accept: right.accept,
	}, nil
}

func (c *Compiler) compileAlternation(node *ast.Alternation) (fragment, error) {
	left, err := c.compileNode(node.Left)
	if err != nil {
		return fragment{}, err
	}

	right, err := c.compileNode(node.Right)
	if err != nil {
		return fragment{}, err
	}

	start := c.builder.NewState()
	accept := c.builder.NewState()

	AddEpsilon(start, left.start)
	AddEpsilon(start, right.start)

	AddEpsilon(left.accept, accept)
	AddEpsilon(right.accept, accept)

	return fragment{
		start:  start,
		accept: accept,
	}, nil
}
