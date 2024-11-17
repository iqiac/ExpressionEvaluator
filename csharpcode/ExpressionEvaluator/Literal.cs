namespace ExpressionEvaluator;

public class Literal : IExpression
{
    private readonly int _value;

    public Literal(int value) => _value = value;

    public int Evaluate() => _value;
}
