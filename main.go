package main

import (
	"errors"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type Stack struct {
	data []interface{}
}

func (s *Stack) Push(val interface{}) {
	fmt.Println("[push]:", val)
	s.data = append(s.data, val)
}

func (s *Stack) Pop() (interface{}, bool) {
	if len(s.data) == 0 {
		return nil, false
	}

	lastIndex := len(s.data) - 1
	lastVal := s.data[lastIndex]
	s.data = s.data[:lastIndex]

	fmt.Println("[pop]: ", lastVal)

	return lastVal, true
}

func (s *Stack) Peek() (interface{}, bool) {
	if len(s.data) == 0 {
		return nil, false
	}
	return s.data[len(s.data)-1], true
}

func (s *Stack) isEmpty() bool {
	return len(s.data) == 0
}

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
	output := &Stack{}
	operators := &Stack{}
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

				if val.(string) == "(" {
					break
				}

				output.Push(val)
			}
		} else {
			if len(operators.data) > 0 {
				for {
					topValue, ok := operators.Peek()

					if !ok {
						break
					}

					if GetPrecedence(topValue.(string)) >= GetPrecedence(token) && topValue.(string) != "(" {
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

	for len(operators.data) > 0 {
		lastVal, _ := operators.Pop()
		output.Push(lastVal)
	}

	fmt.Println(output.data)

	var postfixExpression string

	for _, val := range output.data {
		postfixExpression += val.(string) + " "
	}

	postfixExpression = strings.TrimSpace(postfixExpression)

	return postfixExpression, nil
}

func EvaluatePostfix(expression string) (float64, error) {
	stack := &Stack{}
	tokens := strings.Split(expression, " ")

	for _, token := range tokens {
		if IsNumber(token) {
			fmt.Println("In Stack:", token)
			n, err := strconv.ParseFloat(token, 64)
			if err != nil {
				return 0, errors.New("expression doesn't has valid numeric character")
			}
			stack.Push(n)
		} else {
			b,_ := stack.Pop()
			a,_ := stack.Pop()

			fmt.Println(stack.data)
			fmt.Println(b, a)

			switch token {
			case "+":
				b := b.(float64)
				a := a.(float64)
				stack.Push(a + b)
			case "-":
				b := b.(float64)
				a := a.(float64)
				stack.Push(a - b)
			case "*":
				b := b.(float64)
				a := a.(float64)
				stack.Push(a * b)
			case "/":
				b := b.(float64)
				a := a.(float64)
				stack.Push(a / b)
			case "^":
				b := b.(float64)
				a := a.(float64)
				stack.Push(math.Pow(a, b))
			}
		}
	}
	result, _ := stack.Pop()
	return result.(float64), nil
}

func main() {
	tokens, _ := InfixToPostfix("3 + 4 ^ 2")
	result, _ := EvaluatePostfix(tokens)

	fmt.Println(result)
}