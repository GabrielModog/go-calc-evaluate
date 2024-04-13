package main

import (
	"fmt"
	"strings"

	internal "github.com/GabrielModog/go-calc-evaluate/internal/shunting_yard"
)

func main() {
	expression := ""

	fmt.Print("= ")
	fmt.Scanln(&expression)

	if strings.TrimSpace(expression) == "" {
		panic("insufficient expressions")
	}

	tokens, _ := internal.InfixToPostfix(expression)
	result, _ := internal.EvaluatePostfix(tokens)

	fmt.Println(result)
}