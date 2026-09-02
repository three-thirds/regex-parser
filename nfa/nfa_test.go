package nfa

import "testing"

func TestBuilder(t *testing.T) {
	builder := NewBuilder()

	start := builder.NewState()
	accept := builder.NewState()

	AddLiteral(start, accept, 'a')

	if start.ID != 0 || accept.ID != 1 {
		t.Fatal("unexpected state IDs")
	}

	if len(start.Transitions) != 1 {
		t.Fatal("expected one transition")
	}

	transition := start.Transitions[0]

	if transition.Type != TransitionLiteral ||
		transition.Value != 'a' ||
		transition.To != accept {
		t.Fatal("unexpected Literal transition.")
	}
}
