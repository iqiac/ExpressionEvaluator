using System.Linq.Expressions;

namespace ExpressionEvaluator;

public abstract class BinaryOperation : IExpression
{
    protected readonly IExpression _left;
    protected readonly IExpression _right;

    public BinaryOperation(IExpression left, IExpression right)
    {
        _left = left;
        _right = right;
    }

    public abstract int Evaluate();
}
