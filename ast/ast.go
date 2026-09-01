package ast

type Node interface {
	node()
}

type Literal struct {
	Value rune
}

func (*Literal) node() {}

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
