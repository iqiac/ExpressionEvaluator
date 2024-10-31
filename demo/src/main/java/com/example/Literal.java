package com.example;

public class Literal implements Expression {
  private final int value;

  public Literal(int value) {
    this.value = value;
  }

  @Override
  public int evaluate() {
    return value;
  }
}
