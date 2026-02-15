package main

import (
	"bufio"
	"fmt"
	"os"
	"calculator/internal/parser"
	"strings"
	"calculator/internal/core"
)

type Token struct {
	Number   float64
	Operator string
	IsNum    bool
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("\nEnter expression: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error:", core.ErrInvalidInput)
			continue
		}

		expr := strings.TrimSpace(input)
		expr = strings.ReplaceAll(expr, " ", "")


		if strings.ToLower(expr) == "q" || strings.ToLower(expr) == "exit" {
			fmt.Println("Goodbye!")
			return
		

		} else if expr == "" {
			fmt.Println("Error:", core.ErrEmptyExpression)
			continue
		}


		tokens, err := parser.Tokenize(expr)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		result, err := parser.Evaluate(tokens)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		fmt.Printf("Result: %.2f\n", result)
	}

}
