namespace ExpressionEvaluator.Tests;

public class MultTests
{
    [Test]
    public void testMult1()
    {
        Mult mult = new Mult(new Literal(2), new Literal(3));
        Assert.That(mult.Evaluate(), Is.EqualTo(6));
    }

    [Test]
    public void testMult2()
    {
        Mult mult = new Mult(new Mult(new Literal(3), new Literal(2)), new Literal(4));
        Assert.That(mult.Evaluate(), Is.EqualTo(24));
    }

    [Test]
    public void testMult3()
    {
        Mult mult = new Mult(new Literal(4), new Mult(new Literal(2), new Literal(3)));
        Assert.That(mult.Evaluate(), Is.EqualTo(24));
    }

    [Test]
    public void testMult4()
    {
        Mult mult = new Mult(new Mult(new Literal(2), new Literal(3)), new Mult(new Literal(3), new Literal(2)));
        Assert.That(mult.Evaluate(), Is.EqualTo(36));
    }
}
