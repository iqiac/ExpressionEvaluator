namespace ExpressionEvaluator.Tests;

public class AddTests
{
    [Test]
    public void TestAdd1()
    {
        Add add = new Add(new Literal(5), new Literal(42));
        Assert.That(add.Evaluate(), Is.EqualTo(47));
    }

    [Test]
    public void TestAdd2()
    {
        Add add = new Add(new Add(new Literal(5), new Literal(6)), new Literal(42));
        Assert.That(add.Evaluate(), Is.EqualTo(53));
    }
    [Test]
    public void TestAdd3()
    {
        Add add = new Add(new Literal(42), new Add(new Literal(5), new Literal(6)));
        Assert.That(add.Evaluate(), Is.EqualTo(53));
    }

    [Test]
    public void TestAdd4()
    {
        Add add = new Add(new Add(new Literal(5), new Literal(6)), new Add(new Literal(5), new Literal(6)));
        Assert.That(add.Evaluate(), Is.EqualTo(22));
    }
}
