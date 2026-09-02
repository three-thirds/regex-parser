package ast

import "testing"

func TestPrint(t *testing.T) {
	node := &Concat{
		Left: &Literal{
			Value: 'a',
		},
		Right: &Star{
			Expr: &Alternation{
				Left: &Literal{
					Value: 'b',
				},
				Right: &Literal{
					Value: 'c',
				},
			},
		},
	}

	expected := "Concat\n" +
		"  Literal(a)\n" +
		"  Star\n" +
		"    Alternation\n" +
		"      Literal(b)\n" +
		"      Literal(c)\n"

	actual := Print(node)

	if actual != expected {
		t.Fatalf(
			"expected:\n%q\ngot:\n%q",
			expected,
			actual,
		)
	}
}
