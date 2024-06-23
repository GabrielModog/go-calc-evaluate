package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	parser "github.com/GabrielModog/go-evaluate-expressions/shunting_yard"
)

func usage() {
	fmt.Println("\nUsage:\n- calcexpr <pathname_to_file.txt>\n- or execute program CLI.")
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		usage()

		scanner := bufio.NewScanner(os.Stdin)

		fmt.Print("= ")

		scanner.Scan()
		expressionInput := scanner.Text()

		expression := strings.TrimSpace(expressionInput)

		tokens, _ := parser.InfixToPostfix(expression)
		result, _ := parser.EvaluatePostfix(tokens)

		fmt.Println(result)
		os.Exit(1)
	}

	path := args[0]
	checkExtension := strings.LastIndex(path, ".txt")

	if checkExtension == -1 {
		usage()
	}

	data, err := os.ReadFile(path)
	check(err)

	expressions := strings.Split(string(data), "\n")

	for i := range expressions {
		exp := strings.TrimSpace(expressions[i])
		tokens, _ := parser.InfixToPostfix(exp)
		evaluated, _ := parser.EvaluatePostfix(tokens)

		result := fmt.Sprintf("- %s = %.2f", exp, evaluated)

		fmt.Println(result)
	}

	os.Exit(1)
}
