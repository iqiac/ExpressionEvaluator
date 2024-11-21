package expressionevaluator

import (
	"expressionevaluator/expressionevaluator"
	"testing"
)

func TestLiteral_Evaluate5(t *testing.T) {
	literal := expressionevaluator.NewLiteral(5)

	actual := literal.Evaluate()

	expected := 5
	if actual != expected {
		t.Errorf("expected %d but got %d", expected, actual)
	}
}
func TestLiteral_Evaluate42(t *testing.T) {
	literal := expressionevaluator.NewLiteral(42)

	actual := literal.Evaluate()

	expected := 42
	if actual != expected {
		t.Errorf("expected %d but got %d", expected, actual)
	}
}
