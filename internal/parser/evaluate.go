package parser

import (
	"calculator/internal/core"
	"calculator/internal/service"
)

func Evaluate(tokens []service.Token) (float64, error) {
	if len(tokens) == 0 {
		return 0, core.ErrEmptyExpression
	}

	var newTokens []service.Token
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

			newTokens[len(newTokens)-1] = service.Token{Number: res, IsNum: true}
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
