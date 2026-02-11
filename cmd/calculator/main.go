package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

	fmt.Print("Enter expression: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error:", core.ErrInvalidInput)
		return
	}

	expr := strings.TrimSpace(input)
	expr = strings.ReplaceAll(expr, " ", "")

	if expr == "" {
		fmt.Println("Error:", core.ErrEmptyExpression)
		return
	}

	tokens, err := tokenize(expr)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	result, err := evaluate(tokens)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Result: %.2f\n", result)
}

func tokenize(expr string) ([]Token, error) {
	var tokens []Token
	var numStr strings.Builder

	for i := 0; i < len(expr); i++ {
		c := expr[i]

		if (c >= '0' && c <= '9') || c == '.' {
			numStr.WriteByte(c)
		} else if strings.ContainsRune("+-*/", rune(c)) {
			if numStr.Len() == 0 {
				return nil, core.ErrInvalidExpression
			}

			num, err := strconv.ParseFloat(numStr.String(), 64)
			if err != nil {
				return nil, core.ErrInvalidInput
			}

			tokens = append(tokens, Token{Number: num, IsNum: true})
			numStr.Reset()

			tokens = append(tokens, Token{Operator: string(c), IsNum: false})
		} else {
			return nil, core.ErrInvalidInput
		}
	}

	if numStr.Len() > 0 {
		num, err := strconv.ParseFloat(numStr.String(), 64)
		if err != nil {
			return nil, core.ErrInvalidInput
		}
		tokens = append(tokens, Token{Number: num, IsNum: true})
	}

	return tokens, nil
}

func evaluate(tokens []Token) (float64, error) {
	if len(tokens) == 0 {
		return 0, core.ErrEmptyExpression
	}

	var newTokens []Token
	i := 0

	// First pass: handle * and /
	for i < len(tokens) {
		t := tokens[i]

		if !t.IsNum && (t.Operator == "*" || t.Operator == "/") {
			if len(newTokens) == 0 || i+1 >= len(tokens) {
				return 0, core.ErrInvalidExpression
			}

			prev := newTokens[len(newTokens)-1]
			next := tokens[i+1]

			if !next.IsNum {
				return 0, core.ErrInvalidExpression
			}

			var res float64
			var err error

			if t.Operator == "*" {
				res = core.Mul(prev.Number, next.Number)
			} else if t.Operator == "/" {
				res, err = core.Div(prev.Number, next.Number)
				if err != nil {
					return 0, core.ErrDivisionByZero
				}
			} else {
				return 0, core.ErrUnsupportedOp
			}

			newTokens[len(newTokens)-1] = Token{Number: res, IsNum: true}
			i += 2
		} else {
			newTokens = append(newTokens, t)
			i++
		}
	}

	if len(newTokens) == 0 {
		return 0, core.ErrInvalidExpression
	}

	// Second pass: handle + and -
	result := newTokens[0].Number
	i = 1

	for i < len(newTokens) {
		if i+1 >= len(newTokens) {
			return 0, core.ErrInvalidExpression
		}

		op := newTokens[i].Operator
		next := newTokens[i+1].Number

		switch op {
		case "+":
			result = core.Add(result, next)
		case "-":
			result = core.Sub(result, next)
		default:
			return 0, core.ErrUnsupportedOp
		}

		i += 2
	}

	return result, nil
}
