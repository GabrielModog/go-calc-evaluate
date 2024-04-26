package shunting_yard

import (
	"testing"
)

func TestSimpleArithmeticShuntingYard(t *testing.T) {
	infixExpression := "4 + 3"
	expectedResult := float64(7)

	tokens, _ := InfixToPostfix(infixExpression)
	result, _ := EvaluatePostfix(tokens)

	if result != expectedResult {
		t.Errorf("Result of <4 + 3> expected: %f got %f", expectedResult, result)
	}
}

func TestParenthesisOrderShuntingYard(t *testing.T) {
	infixExpression := "2 + 1 + (4 + 8)"
	expectedResult := float64(15)

	tokens, _ := InfixToPostfix(infixExpression)
	result, _ := EvaluatePostfix(tokens)

	if result != expectedResult {
		t.Errorf("Result of <2 + 1 + (4 + 8)> expected: %f got %f", expectedResult, result)
	}
}