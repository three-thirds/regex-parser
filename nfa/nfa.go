package nfa

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

func AddEplipson(from, to *State) {
	from.Transitions = append(
		from.Transitions,
		Transition{
			Type: TransitionEpsilon,
			To:   to,
		},
	)
}

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

func AddAny(from, to *State) {
	from.Transitions = append(
		from.Transitions,
		Transition{
			Type: TransitionAny,
			To:   to,
		},
	)
}
