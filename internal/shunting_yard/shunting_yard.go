package internal

import (
	"errors"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	internal "github.com/GabrielModog/go-calc-evaluate/internal/stack"
)

func GetPrecedence(operator string) int {
	precedences := map[string]int{
		"+": 1,
		"-": 1,
		"*": 2,
		"/": 2,
		"^": 3,
	}
	return precedences[operator]
}

func Tokenize(expression string) []string {
	pattern := `\d+|\+|-|\*|/|\(|\)|\^`
	token_regex, err := regexp.Compile(pattern)

	if err != nil {
		return []string{"ERROR"}
	}

	return token_regex.FindAllString(expression, -1)
}

func IsNumber(str string) bool {
	match, err := regexp.MatchString("^-?[0-9]+(?:\\.[0-9]*)?$", str)
	if err != nil {
		return false
	}
  return match
}

func InfixToPostfix(expression string) (string, error) {
	output := &internal.Stack[string]{}
	operators := &internal.Stack[string]{}

	tokens := Tokenize(expression)

	for _, token := range tokens {
		if IsNumber(token) {
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

					if GetPrecedence(topValue) >= GetPrecedence(token) && topValue != "(" {
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

	fmt.Println(output.Data)

	var postfixExpression string

	for _, val := range output.Data {
		postfixExpression += val + " "
	}

	postfixExpression = strings.TrimSpace(postfixExpression)

	return postfixExpression, nil
}

func EvaluatePostfix(expression string) (float64, error) {
	stack := &internal.Stack[float64]{}
	tokens := strings.Split(expression, " ")

	for _, token := range tokens {
		if IsNumber(token) {
			n, err := strconv.ParseFloat(token, 64)
			if err != nil {
				return 0, errors.New("expression doesn't has valid numeric character")
			}
			stack.Push(n)
		} else {
			b,_ := stack.Pop()
			a,_ := stack.Pop()

			switch token {
			case "+":
				stack.Push(a + b)
			case "-":
				stack.Push(a - b)
			case "*":
				stack.Push(a * b)
			case "/":
				stack.Push(a / b)
			case "^":
				stack.Push(math.Pow(a, b))
			}
		}
	}

	result, _ := stack.Pop()
	
	return result, nil
}