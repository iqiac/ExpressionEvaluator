namespace ExpressionEvaluator;

public class Mult : BinaryOperation
{
    public Mult(IExpression left, IExpression right) : base(left, right) { }

    public override int Evaluate()
    {
        return _left.Evaluate() * _right.Evaluate();
    }
}
