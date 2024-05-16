package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	parser "github.com/GabrielModog/go-evaluate-expressions/shunting_yard"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	
	fmt.Print("= ")

	scanner.Scan()
	expressionInput := scanner.Text()

	expression := strings.TrimSpace(expressionInput)

	tokens, _ := parser.InfixToPostfix(expression)
	result, _ := parser.EvaluatePostfix(tokens)

	fmt.Println(result)
}