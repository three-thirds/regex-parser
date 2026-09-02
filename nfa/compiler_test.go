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
