package nfa

// Package nfa contains a Compiler that converts AST nodes into NFA
// fragments used during regex compilation.

import (
	"fmt"

	"regex/ast"
)

// Compiler converts AST nodes into NFA fragments used during
// regex compilation.
type Compiler struct {
	builder *Builder
}

// fragment represents a small NFA with a single entry (`start`)
// and a single accept (`accept`) state used while composing
// larger NFAs from AST nodes.
type fragment struct {
	start  *State
	accept *State
}

func NewCompiler() *Compiler {
	return &Compiler{
		builder: NewBuilder(),
	}
}

// NewCompiler constructs and returns a ready-to-use Compiler.

func Compile(node ast.Node) (*NFA, error) {
	compiler := NewCompiler()
	return compiler.Compile(node)
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

	case *ast.Star:
		return c.compileStar(n)

	case *ast.Plus:
		return c.compilePlus(n)

	case *ast.Question:
		return c.compileQuestion(n)

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

func (c *Compiler) compileStar(node *ast.Star) (fragment, error) {
	inner, err := c.compileNode(node.Expr)
	if err != nil {
		return fragment{}, err
	}

	start := c.builder.NewState()
	accept := c.builder.NewState()

	AddEpsilon(start, inner.start)
	AddEpsilon(start, accept)

	AddEpsilon(inner.accept, inner.start)
	AddEpsilon(inner.accept, accept)

	return fragment{
		start:  start,
		accept: accept,
	}, nil
}

func (c *Compiler) compilePlus(node *ast.Plus) (fragment, error) {
	inner, err := c.compileNode(node.Expr)
	if err != nil {
		return fragment{}, err
	}

	start := c.builder.NewState()
	accept := c.builder.NewState()

	AddEpsilon(inner.accept, inner.start)
	AddEpsilon(inner.accept, accept)

	return fragment{
		start:  start,
		accept: accept,
	}, nil
}

// compilePlus creates a fragment that matches one or more repetitions
// of the inner expression (the `+` quantifier). It ensures at least
// one instance of the inner fragment can be matched, then loops back
// to allow additional repetitions.

func (c *Compiler) compileQuestion(node *ast.Question) (fragment, error) {
	inner, err := c.compileNode(node.Expr)
	if err != nil {
		return fragment{}, err
	}

	start := c.builder.NewState()
	accept := c.builder.NewState()

	AddEpsilon(start, inner.start)
	AddEpsilon(start, accept)

	AddEpsilon(inner.accept, accept)

	return fragment{
		start:  start,
		accept: accept,
	}, nil
}

// compileQuestion creates a fragment that matches zero or one
// occurrences of the inner expression (the `?` quantifier). It
// provides an epsilon path that bypasses the inner fragment.
