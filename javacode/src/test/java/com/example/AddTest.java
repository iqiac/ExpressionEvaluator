package com.example;

import static org.junit.Assert.assertEquals;
import org.junit.Test;

public class AddTest {
  @Test
  public void testAdd1() {
    Add add = new Add(new Literal(5), new Literal(42));
    assertEquals(47, add.evaluate());
  }

  @Test
  public void testAdd2() {
    Add add = new Add(new Add(new Literal(5), new Literal(6)), new Literal(42));
    assertEquals(53, add.evaluate());
  }

  @Test
  public void testAdd3() {
    Add add = new Add(new Literal(42), new Add(new Literal(5), new Literal(6)));
    assertEquals(53, add.evaluate());
  }

  @Test
  public void testAdd4() {
    Add add = new Add(new Add(new Literal(5), new Literal(6)), new Add(new Literal(5), new Literal(6)));
    assertEquals(22, add.evaluate());
  }
}
