package core

import "errors"

var (
	ErrInvalidInput      = errors.New("invalid input")
	ErrDivisionByZero    = errors.New("Cannot divide by zero")
	ErrUnsupportedOp     = errors.New("unsupported operation")
	ErrEmptyExpression   = errors.New("empty expression")
	ErrInvalidExpression = errors.New("invalid expression")
)
