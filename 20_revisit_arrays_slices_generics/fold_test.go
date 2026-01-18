package revisitarraysslicesgenerics

import (
	"strings"
	"testing"
)

func TestReduce(t *testing.T) {
	t.Run("addition", func(t *testing.T) {
		list := []int{1, 2, 3, 4}
		want := 10

		acc := func(a, b int) int { return (a + b) }

		got := Reduce(0, acc, list)

		if got != want {
			t.Errorf("got %+v want %+v", got, want)
		}
	})

	t.Run("multiplication of all elements", func(t *testing.T) {
		multiply := func(x, y int) int {
			return x * y
		}
		AssertEqual(t, Reduce(1, multiply, []int{1, 2, 3}), 6)
	})

	t.Run("concatenate string", func(t *testing.T) {
		concatenate := func(a, b string) string {
			return (a + b)
		}

		AssertEqual(t, Reduce("", concatenate, []string{"a", "b", "c"}), "abc")
	})
}

func TestFind(t *testing.T) {
	t.Run("find first even number", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		firstEvenNumber, found := Find(numbers, func(x int) bool {
			return x%2 == 0
		})

		AssertTrue(t, found)
		AssertEqual(t, firstEvenNumber, 2)
	})
	type Person struct {
		Name string
	}

	t.Run("Find the best programmer", func(t *testing.T) {
		people := []Person{
			Person{Name: "Kent Beck"},
			Person{Name: "Martin Fowler"},
			Person{Name: "Chris James"},
		}

		king, found := Find(people, func(p Person) bool {
			return strings.Contains(p.Name, "Chris")
		})

		AssertTrue(t, found)
		AssertEqual(t, king, Person{Name: "Chris James"})
	})
}

func AssertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func AssertTrue(t *testing.T, got bool) {
	t.Helper()
	if !got {
		t.Errorf("got %v want true", got)
	}
}
