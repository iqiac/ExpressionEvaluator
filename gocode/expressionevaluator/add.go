package expressionevaluator

type Add struct {
	BinaryOperation
}

func NewAdd(left IExpression, right IExpression) *Add {
	return &Add{BinaryOperation{left, right}}
}

func (a Add) Evaluate() int {
	return a.left.Evaluate() + a.right.Evaluate()
}
