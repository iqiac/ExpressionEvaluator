package com.example;

public class Mult extends BinaryOperation {
  public Mult(Expression left, Expression right) {
    super(left, right);
  }

  @Override
  public int evaluate() {
    return left.evaluate() * right.evaluate();
  }
}
