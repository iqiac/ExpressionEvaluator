package com.example;

public class Add extends BinaryOperation {
  public Add(Expression left, Expression right) {
    super(left, right);
  }

  @Override
  public int evaluate() {
    return left.evaluate() + right.evaluate();
  }
}
