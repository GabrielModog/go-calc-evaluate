package internal

import (
	"testing"
)

func TestStackPushAndSeeFirstElement(t *testing.T) {
	expectedResultFirstEl := 34
	expectedResultSecondEl := 36
	expectedResultLastEl := 44

	stack := &Stack{}

	stack.Push(34)
	stack.Push(36)
	stack.Push(44)
	
	firstEl := stack.Data[0]
	secondEl := stack.Data[1]
	lastEl := stack.Data[2]

	if firstEl != expectedResultFirstEl {
		t.Errorf("Result expected: %d got %d", expectedResultFirstEl, firstEl)
	}

	if secondEl != expectedResultSecondEl {
		t.Errorf("Result expected: %d got %d", expectedResultSecondEl, secondEl)
	}

	if lastEl != expectedResultLastEl {
		t.Errorf("Result of <4 + 3> expected: %d got %d", expectedResultLastEl, lastEl)
	}
}