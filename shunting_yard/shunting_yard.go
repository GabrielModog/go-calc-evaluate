package shunting_yard

import (
	"errors"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	stack "github.com/GabrielModog/go-evaluate-expressions/internal/stack"
)

func getPrecedence(operator string) int {
	precedences := map[string]int{
		"+": 1,
		"-": 1,
		"*": 2,
		"/": 2,
		"^": 3,
	}
	return precedences[operator]
}

func tokenize(expression string) ([]string, error) {
	pattern := `\d+|\+|-|\*|/|\(|\)|\^`
	token_regex, err := regexp.Compile(pattern)

	if err != nil {
		return nil, fmt.Errorf("[error]: %w", err)
	}

	return token_regex.FindAllString(expression, -1), nil
}

func isNumber(str string) bool {
	match, err := regexp.MatchString("^-?[0-9]+(?:\\.[0-9]*)?$", str)
	if err != nil {
		return false
	}
  return match
}

func InfixToPostfix(expression string) (string, error) {
	output := &stack.Stack[string]{}
	operators := &stack.Stack[string]{}

	tokens, err := tokenize(expression)

	if err != nil {
		return "", err
	}

	for _, token := range tokens {
		if isNumber(token) {
			output.Push(token)
		} else if token == "(" {
			operators.Push(token)
		} else if token == ")" {
			for {
				val, ok := operators.Pop()

				if !ok {
					return "", errors.New("unmatched parenthesis")
				}

				if val == "(" {
					break
				}

				output.Push(val)
			}
		} else {
			if len(operators.Data) > 0 {
				for {
					topValue, ok := operators.Peek()

					if !ok {
						break
					}

					if getPrecedence(topValue) >= getPrecedence(token) && topValue != "(" {
						value, _ := operators.Pop()
						output.Push(value)
					}	else {
						break
					}
				}
			}

			operators.Push(token)
		}
	}

	for len(operators.Data) > 0 {
		lastVal, _ := operators.Pop()
		output.Push(lastVal)
	}

	var postfixExpression string

	for _, val := range output.Data {
		postfixExpression += val + " "
	}

	postfixExpression = strings.TrimSpace(postfixExpression)

	return postfixExpression, nil
}

func EvaluatePostfix(expression string) (float64, error) {
	opStack := &stack.Stack[float64]{}
	tokens := strings.Split(expression, " ")

	for _, token := range tokens {
		if isNumber(token) {
			n, err := strconv.ParseFloat(token, 64)
			if err != nil {
				return 0, errors.New("expression doesn't has valid numeric character")
			}
			opStack.Push(n)
		} else {
			b,_ := opStack.Pop()
			a,_ := opStack.Pop()

			switch token {
			case "+":
				opStack.Push(a + b)
			case "-":
				opStack.Push(a - b)
			case "*":
				opStack.Push(a * b)
			case "/":
				opStack.Push(a / b)
			case "^":
				opStack.Push(math.Pow(a, b))
			}
		}
	}

	result, _ := opStack.Pop()
	
	return result, nil
}