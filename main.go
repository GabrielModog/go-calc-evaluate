package main

import (
	"fmt"
	"strings"

	"github.com/GabrielModog/go-calc-evaluate/internal"
)

type Num interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

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