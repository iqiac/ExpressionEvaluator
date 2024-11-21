package expressionevaluator

import (
	"expressionevaluator/expressionevaluator"
	"testing"
)

func TestAdd1(t *testing.T) {
	add := expressionevaluator.NewAdd(expressionevaluator.NewLiteral(5), expressionevaluator.NewLiteral(42))

	actual := add.Evaluate()

	expected := 47
	if actual != expected {
		t.Errorf("expected %d but got %d", expected, actual)
	}
}

func TestAdd2(t *testing.T) {
	add := expressionevaluator.NewAdd(
		expressionevaluator.NewAdd(
			expressionevaluator.NewLiteral(5),
			expressionevaluator.NewLiteral(6),
		),
		expressionevaluator.NewLiteral(42),
	)

	actual := add.Evaluate()

	expected := 53
	if actual != expected {
		t.Errorf("expected %d but got %d", expected, actual)
	}
}
func TestAdd3(t *testing.T) {
	add := expressionevaluator.NewAdd(
		expressionevaluator.NewLiteral(42),
		expressionevaluator.NewAdd(
			expressionevaluator.NewLiteral(5),
			expressionevaluator.NewLiteral(6),
		),
	)

	actual := add.Evaluate()

	expected := 53
	if actual != expected {
		t.Errorf("expected %d but got %d", expected, actual)
	}
}
func TestAdd4(t *testing.T) {
	add := expressionevaluator.NewAdd(
		expressionevaluator.NewAdd(
			expressionevaluator.NewLiteral(5),
			expressionevaluator.NewLiteral(6),
		),
		expressionevaluator.NewAdd(
			expressionevaluator.NewLiteral(5),
			expressionevaluator.NewLiteral(6),
		),
	)

	actual := add.Evaluate()

	expected := 22
	if actual != expected {
		t.Errorf("expected %d but got %d", expected, actual)
	}
}
