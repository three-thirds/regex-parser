package ast

type Node interface { // Define the node interface that all AST nodes must implement.
	node()
}

type Literal struct { // Defines literal AST node representing a literal character.
	Value rune
}

func (*Literal) node() {} // Implements the Node interface for the literal.
// Defines all stuff like Dot, Concat, Alternation, Star, Plus, Question, etc.
// In the lines given below by storing them then implementing the node interface.
type Dot struct{}

func (*Dot) node() {}

type Concat struct {
	Left  Node
	Right Node
}

func (*Concat) node() {}

type Alternation struct {
	Left  Node
	Right Node
}

func (*Alternation) node() {}

type Star struct {
	Expr Node
}

func (*Star) node() {}

type Plus struct {
	Expr Node
}

func (*Plus) node() {}

type Question struct {
	Expr Node
}

func (*Question) node() {}
