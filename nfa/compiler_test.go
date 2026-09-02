package nfa

import (
	"testing"

	"regex/ast"
)

func TestCompileBasicNodes(t *testing.T) {
	t.Run("literal", func(t *testing.T) {
		compiler := NewCompiler()

		machine, err := compiler.Compile(
			&ast.Literal{Value: 'a'},
		)
		if err != nil {
			t.Fatalf("unexpected compile error: %v", err)
		}
		if len(machine.Start.Transitions) != 1 {
			t.Fatal("expected one transition")
		}

		transition := machine.Start.Transitions[0]

		if transition.Type != TransitionLiteral ||
			transition.Value != 'a' ||
			transition.To != machine.Accept {
			t.Fatal("unexpected literal NFA")
		}
	})

	t.Run("dot", func(t *testing.T) {
		compiler := NewCompiler()

		machine, err := compiler.Compile(&ast.Dot{})
		if err != nil {
			t.Fatalf("unexpected compile error: %v", err)
		}

		if len(machine.Start.Transitions) != 1 {
			t.Fatal("expected one transition")
		}

		transition := machine.Start.Transitions[0]

		if transition.Type != TransitionAny ||
			transition.To != machine.Accept {
			t.Fatal("unexpected dot NFA")
		}
	})
}

func TestCompileConcat(t *testing.T) {
	compiler := NewCompiler()

	node := &ast.Concat{
		Left:  &ast.Literal{Value: 'a'},
		Right: &ast.Literal{Value: 'b'},
	}

	machine, err := compiler.Compile(node)
	if err != nil {
		t.Fatalf("unexpected compile error: %v", err)
	}

	if len(machine.Start.Transitions) != 1 {
		t.Fatal("expected literal a transition")
	}

	first := machine.Start.Transitions[0]
	if first.Type != TransitionLiteral ||
		first.Value != 'a' {
		t.Fatal("unexpected first literal transition")
	}

	if len(first.To.Transitions) != 1 {
		t.Fatal("expected epsilon transition after literal a")
	}

	epsilon := first.To.Transitions[0]
	if epsilon.Type != TransitionEpsilon {
		t.Fatal("expected epsilon transition")
	}

	secondStart := epsilon.To
	if len(secondStart.Transitions) != 1 {
		t.Fatal("expected literal b transition")
	}

	second := secondStart.Transitions[0]
	if second.Type != TransitionLiteral ||
		second.Value != 'b' ||
		second.To != machine.Accept {
		t.Fatal("unexpected concatenated NFA")
	}
}

func TestCompileAlternation(t *testing.T) {
	compiler := NewCompiler()

	node := &ast.Alternation{
		Left:  &ast.Literal{Value: 'a'},
		Right: &ast.Literal{Value: 'b'},
	}

	machine, err := compiler.Compile(node)
	if err != nil {
		t.Fatalf("unexpected compile error: %v", err)
	}

	if len(machine.Start.Transitions) != 2 {
		t.Fatal("expected two epsilon branches from start")
	}

	for _, transition := range machine.Start.Transitions {
		if transition.Type != TransitionEpsilon {
			t.Fatal("expected epsilon transition from an alternation start")
		}
	}
}
