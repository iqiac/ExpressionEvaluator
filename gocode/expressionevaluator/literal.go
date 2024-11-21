package expressionevaluator

type Literal struct {
	Value int
}

func (l *Literal) Evaluate() int {
	return l.Value
}

func NewLiteral(value int) *Literal {
	return &Literal{Value: value}
}
