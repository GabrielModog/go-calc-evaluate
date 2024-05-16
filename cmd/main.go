package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	parser "github.com/GabrielModog/go-calc-evaluate/shunting_yard"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	
	fmt.Print("= ")

	scanner.Scan()
	expressionInput := scanner.Text()

	expression := strings.TrimSpace(expressionInput)

	fmt.Println("expression:",expression)

	fmt.Println("input:", expressionInput)

	tokens, _ := parser.InfixToPostfix(expression)
	result, _ := parser.EvaluatePostfix(tokens)

	fmt.Println(result)
}