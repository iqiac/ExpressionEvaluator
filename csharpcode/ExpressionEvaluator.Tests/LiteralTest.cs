namespace ExpressionEvaluator.Tests;

public class LiteralTests
{
    [Test]
    public void Evaluate5()
    {
        Literal literal = new Literal(5);
        Assert.That(literal.Evaluate(), Is.EqualTo(5));
    }

    [Test]
    public void Evaluate42()
    {
        Literal literal = new Literal(42);
        Assert.That(literal.Evaluate(), Is.EqualTo(42));
    }
}
