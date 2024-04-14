package internal

import (
	"testing"
)

func TestStackPushIntegers(t *testing.T) {
	expectedResultFirstEl := 34
	expectedResultSecondEl := 36
	expectedResultLastEl := 44

	stack := &Stack[int]{}

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
		t.Errorf("Result expected: %d got %d", expectedResultLastEl, lastEl)
	}
}

func TestStackPushFloats(t *testing.T) {
	expectedResultFirstEl := float32(34.55)
	expectedResultSecondEl := float32(36.55)
	expectedResultLastEl := float32(44.55)

	stack := &Stack[float32]{}

	stack.Push(34.55)
	stack.Push(36.55)
	stack.Push(44.55)
	
	firstEl := stack.Data[0]
	secondEl := stack.Data[1]
	lastEl := stack.Data[2]

	if firstEl != expectedResultFirstEl {
		t.Errorf("Result expected: %f got %f", expectedResultFirstEl, firstEl)
	}

	if secondEl != expectedResultSecondEl {
		t.Errorf("Result expected: %f got %f", expectedResultSecondEl, secondEl)
	}

	if lastEl != expectedResultLastEl {
		t.Errorf("Result expected: %f got %f", expectedResultLastEl, lastEl)
	}
}

func TestStackPushStrings(t *testing.T) {
	expectedResultFirstEl := "34"
	expectedResultSecondEl := "36"
	expectedResultLastEl := "44"

	stack := &Stack[string]{}

	stack.Push("34")
	stack.Push("36")
	stack.Push("44")
	
	firstEl := stack.Data[0]
	secondEl := stack.Data[1]
	lastEl := stack.Data[2]

	if firstEl != expectedResultFirstEl {
		t.Errorf("Result expected: %s got %s", expectedResultFirstEl, firstEl)
	}

	if secondEl != expectedResultSecondEl {
		t.Errorf("Result expected: %s got %s", expectedResultSecondEl, secondEl)
	}

	if lastEl != expectedResultLastEl {
		t.Errorf("Result expected: %s got %s", expectedResultLastEl, lastEl)
	}
}