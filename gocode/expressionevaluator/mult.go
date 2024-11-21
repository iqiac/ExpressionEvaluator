package expressionevaluator

type Mult struct {
	BinaryOperation
}

func NewMult(left IExpression, right IExpression) *Mult {
	return &Mult{BinaryOperation{left, right}}
}

func (m Mult) Evaluate() int {
	return m.left.Evaluate() * m.right.Evaluate()
}
