package expressionevaluator

import (
	"expressionevaluator/expressionevaluator"
	"testing"
)

func TestMult1(t *testing.T) {
	add := expressionevaluator.NewMult(expressionevaluator.NewLiteral(2), expressionevaluator.NewLiteral(3))

	actual := add.Evaluate()

	expected := 6
	if actual != expected {
		t.Errorf("expected %d but got %d", expected, actual)
	}
}

func TestMult2(t *testing.T) {
	add := expressionevaluator.NewMult(
		expressionevaluator.NewMult(
			expressionevaluator.NewLiteral(3),
			expressionevaluator.NewLiteral(2),
		),
		expressionevaluator.NewLiteral(4),
	)

	actual := add.Evaluate()

	expected := 24
	if actual != expected {
		t.Errorf("expected %d but got %d", expected, actual)
	}
}
func TestMult3(t *testing.T) {
	add := expressionevaluator.NewMult(
		expressionevaluator.NewLiteral(4),
		expressionevaluator.NewMult(
			expressionevaluator.NewLiteral(2),
			expressionevaluator.NewLiteral(3),
		),
	)

	actual := add.Evaluate()

	expected := 24
	if actual != expected {
		t.Errorf("expected %d but got %d", expected, actual)
	}
}
func TestMult4(t *testing.T) {
	add := expressionevaluator.NewMult(
		expressionevaluator.NewMult(
			expressionevaluator.NewLiteral(2),
			expressionevaluator.NewLiteral(3),
		),
		expressionevaluator.NewMult(
			expressionevaluator.NewLiteral(3),
			expressionevaluator.NewLiteral(2),
		),
	)

	actual := add.Evaluate()

	expected := 36
	if actual != expected {
		t.Errorf("expected %d but got %d", expected, actual)
	}
}
