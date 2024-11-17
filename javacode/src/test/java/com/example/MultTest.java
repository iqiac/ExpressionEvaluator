package com.example;

import static org.junit.Assert.assertEquals;
import org.junit.Test;

public class MultTest {
  @Test
  public void testMult1() {
    Mult mult = new Mult(new Literal(2), new Literal(3));
    assertEquals(6, mult.evaluate());
  }

  @Test
  public void testMult2() {
    Mult mult = new Mult(new Mult(new Literal(3), new Literal(2)), new Literal(4));
    assertEquals(24, mult .evaluate());
  }

  @Test
  public void testMult3() {
    Mult mult = new Mult(new Literal(4), new Mult(new Literal(2), new Literal(3)));
    assertEquals(24, mult .evaluate());
  }

  @Test
  public void testMult4() {
    Mult mult = new Mult(new Mult(new Literal(2), new Literal(3)), new Mult(new Literal(3), new Literal(2)));
    assertEquals(36, mult .evaluate());
  }
}
