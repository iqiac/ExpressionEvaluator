package com.example;

import static org.junit.Assert.assertEquals;

import org.junit.Test;

public class LiteralTest {
  @Test
  public void evaluate5() {
    Literal literal = new Literal(5);
    assertEquals(5, literal.evaluate());
  }

  @Test
  public void evaluate42() {
    Literal literal = new Literal(42);
    assertEquals(42, literal.evaluate());
  }
}
