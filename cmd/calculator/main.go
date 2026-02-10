package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"calculator/internal/core"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Example: 6 + 7 * 5 / 14")
		return
	}

	expr := strings.Join(os.Args[1:], "")
	expr = strings.ReplaceAll(expr, " ", "")

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

type Token struct {
	Number   float64
	Operator string
	IsNum    bool
}

func tokenize(expr string) ([]Token, error) {
	var tokens []Token
	var numStr strings.Builder

	for i := 0; i < len(expr); i++ {
		c := expr[i]
		if c >= '0' && c <= '9' || c == '.' {
			numStr.WriteByte(c)
		} else if strings.ContainsRune("+-*/", rune(c)) {
			if numStr.Len() == 0 {
				return nil, fmt.Errorf("operator %c cannot be at this position", c)
			}
			num, err := strconv.ParseFloat(numStr.String(), 64)
			if err != nil {
				return nil, err
			}
			tokens = append(tokens, Token{Number: num, IsNum: true})
			numStr.Reset()
			tokens = append(tokens, Token{Operator: string(c), IsNum: false})
		} else {
			return nil, fmt.Errorf("invalid character: %c", c)
		}
	}

	if numStr.Len() > 0 {
		num, err := strconv.ParseFloat(numStr.String(), 64)
		if err != nil {
			return nil, err
		}
		tokens = append(tokens, Token{Number: num, IsNum: true})
	}

	return tokens, nil
}

func evaluate(tokens []Token) (float64, error) {
	var newTokens []Token
	i := 0
	for i < len(tokens) {
		t := tokens[i]
		if !t.IsNum && (t.Operator == "*" || t.Operator == "/") {
			if len(newTokens) == 0 {
				return 0, fmt.Errorf("invalid expression")
			}
			prev := newTokens[len(newTokens)-1]
			next := tokens[i+1]
			if !next.IsNum {
				return 0, fmt.Errorf("invalid expression")
			}

			var res float64
			var err error
			if t.Operator == "*" {
				res = core.Mul(prev.Number, next.Number)
			} else {
				res, err = core.Div(prev.Number, next.Number)
				if err != nil {
					if err == core.ErrDivisionByZero {
						return 0, core.ErrDivisionByZero
					}
					return 0, err
				}
			}

			newTokens[len(newTokens)-1] = Token{Number: res, IsNum: true}
			i += 2
		} else {
			newTokens = append(newTokens, t)
			i++
		}
	}

	result := newTokens[0].Number
	i = 1
	for i < len(newTokens) {
		op := newTokens[i].Operator
		next := newTokens[i+1].Number
		switch op {
		case "+":
			result = core.Add(result, next)
		case "-":
			result = core.Sub(result, next)
		default:
			return 0, fmt.Errorf("unexpected operator: %s", op)
		}
		i += 2
	}

	return result, nil
}
