package com.example;

public abstract class BinaryOperation implements Expression {
  protected final Expression left;
  protected final Expression right;

  public BinaryOperation(Expression left, Expression right) {
    this.left = left;
    this.right = right;
  }

  @Override
  public abstract int evaluate();
}
