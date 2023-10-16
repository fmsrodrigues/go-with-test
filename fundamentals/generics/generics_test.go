package generics_test

import (
	g "fundamentals/generics"
	"testing"
)

func TestAssertFunctions(t *testing.T) {
	t.Run("Asserting on integers", func(t *testing.T) {
		AssertEqual(t, 1, 1)
		AssertNotEqual(t, 1, 2)
	})

	t.Run("Asserting on strings", func(t *testing.T) {
		AssertEqual(t, "Hello", "Hello")
		AssertNotEqual(t, "Hello", "World")
	})
}

func TestStack(t *testing.T) {
	t.Run("Integer stack", func(t *testing.T) {
		myStackOfInts := new(g.Stack[int])

		AssertTrue(t, myStackOfInts.IsEmpty())

		myStackOfInts.Push(123)
		AssertFalse(t, myStackOfInts.IsEmpty())

		myStackOfInts.Push(456)
		value, _ := myStackOfInts.Pop()
		AssertEqual(t, value, 456)
		value, _ = myStackOfInts.Pop()
		AssertEqual(t, value, 123)
		AssertTrue(t, myStackOfInts.IsEmpty())
	})

	t.Run("String stack", func(t *testing.T) {
		myStackOfStrings := new(g.Stack[string])

		AssertTrue(t, myStackOfStrings.IsEmpty())

		myStackOfStrings.Push("Hello")
		AssertFalse(t, myStackOfStrings.IsEmpty())

		myStackOfStrings.Push("World")
		value, _ := myStackOfStrings.Pop()
		AssertEqual(t, value, "World")
		value, _ = myStackOfStrings.Pop()
		AssertEqual(t, value, "Hello")
		AssertTrue(t, myStackOfStrings.IsEmpty())
	})
}
