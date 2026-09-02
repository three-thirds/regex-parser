package nfa

// Package nfa provides simple non-deterministic finite automaton
// (NFA) construction primitives used by the regex compiler.

type TransitionType int

const (
	TransitionEpsilon TransitionType = iota
	TransitionLiteral
	TransitionAny
)

type State struct {
	ID          int
	Transitions []Transition
}

type Transition struct {
	Type  TransitionType
	Value rune
	To    *State
}

type NFA struct {
	Start  *State
	Accept *State
}

type Builder struct {
	nextID int
}

func NewBuilder() *Builder {
	return &Builder{}
}

func (b *Builder) NewState() *State {
	state := &State{
		ID: b.nextID,
	}

	b.nextID++

	return state
}

// AddEplipson adds an epsilon (empty) transition from `from` to `to`.
// NOTE: name kept as-is to avoid changing external call sites.
func AddEplipson(from, to *State) {
	from.Transitions = append(
		from.Transitions,
		Transition{
			Type: TransitionEpsilon,
			To:   to,
		},
	)
}

// AddLiteral adds a transition that consumes the given rune value.
func AddLiteral(from, to *State, value rune) {
	from.Transitions = append(
		from.Transitions,
		Transition{
			Type:  TransitionLiteral,
			Value: value,
			To:    to,
		},
	)
}

// AddAny adds a transition that matches any single rune.
func AddAny(from, to *State) {
	from.Transitions = append(
		from.Transitions,
		Transition{
			Type: TransitionAny,
			To:   to,
		},
	)
}
