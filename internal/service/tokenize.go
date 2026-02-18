package service

import (
	"strings"
	"calculator/internal/core"
	"strconv"
)

type Token struct {
	Number   float64
	Operator string
	IsNum    bool
}

func Tokenize(expr string) ([]Token, error) {
	var tokens []Token
	var numStr strings.Builder

	if len(expr) == 0 {
		return []Token{}, nil
	}

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

