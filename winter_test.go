package winter

import "testing"

type T struct {
  value string
}

func newT() *T {
  return &T{"test"}
}

func TestDipi(t *testing.T) {
  container := NewContainer()

  container.add(newT())
  container.add("jose")
  container.add(10)

  count := 0

  function := func (number int) {
    if number != 10 {
      t.Errorf("Number should be 10 but was %d", number)
    }
    count++
  }

  function1 := func (z *T) {
    if z.value != "test" {
      t.Errorf("T.value should be test but was %s", z.value)
    }
    count++
  }

  function2 := func (name string) {
    if (name != "jose") {
      t.Errorf("Name should be jose but was %s", name)
    }
    count++
  }

  function3 := func (number int, z *T, name string) {
    if number != 10 {
      t.Errorf("Number should be 10 but was %d", number)
    }

    if z.value != "test" {
      t.Errorf("T.value should be test but was %s", z.value)
    }

    if (name != "jose") {
      t.Errorf("Name should be jose but was %s", name)
    }
    count++
  }

  container.execute(function)
  container.execute(function1)
  container.execute(function2)
  container.execute(function3)

  if count != 4 {
    t.Errorf("Count should be 4 but was %d", count)
  }
}
