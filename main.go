package main

import (
	"fmt"
	"strings"

	parser "github.com/GabrielModog/go-calc-evaluate/shunting_yard"
)

func main() {
	expression := ""

	fmt.Print("= ")
	fmt.Scanln(&expression)

	if strings.TrimSpace(expression) == "" {
		panic("insufficient expressions")
	}

	tokens, _ := parser.InfixToPostfix(expression)
	result, _ := parser.EvaluatePostfix(tokens)

	fmt.Println(result)
}