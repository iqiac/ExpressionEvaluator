namespace ExpressionEvaluator;

public class Add : BinaryOperation
{
    public Add(IExpression left, IExpression right) : base(left, right) { }

    public override int Evaluate()
    {
        return _left.Evaluate() + _right.Evaluate();
    }
}
