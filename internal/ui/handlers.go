package ui

import (
	"calculator/internal/parser"
	"calculator/internal/service"
	"fmt"
	"strings"

	"fyne.io/fyne/v2/widget"
)

var display *widget.Entry
var justEvaluated bool

// SetDisplay sets the calculator display (called from ui.go)
func SetDisplay(entry *widget.Entry) {
	display = entry
}

// HandleNumber appends a number to the display
func HandleNumber(num string) {
	if display == nil {
		return
	}

	if justEvaluated {
		display.SetText(num)
		justEvaluated = false
		return
	}

	if display.Text == "0" {
		display.SetText(num)
		return
	}

	display.SetText(display.Text + num)
}

// HandleOperator appends an operator
func HandleOperator(op string) {
	if display == nil {
		return
	}

	text := display.Text
	if text == "" {
		return
	}

	justEvaluated = false

	// Prevent duplicate operators
	lastChar := text[len(text)-1:]
	if strings.ContainsAny(lastChar, "+-*/") {
		return
	}

	display.SetText(text + op)
}

// HandleClear resets the display
func HandleClear() {
	if display == nil {
		return
	}

	display.SetText("0")
}

// HandleEqual evaluates the current expression
func HandleEqual() {
	if display == nil {
		return
	}

	expression := display.Text
	if expression == "" {
		return
	}

	// Business logic (shared)
	tokens, err := service.Tokenize(expression)
	if err != nil {
		display.SetText(fmt.Sprintf("Error"))
		justEvaluated = true
		return
	}

	result, err := parser.Evaluate(tokens)
	if err != nil {
		display.SetText(fmt.Sprintf("Error"))
		justEvaluated = true
		return
	}

	display.SetText(fmt.Sprintf("%v", result))
	justEvaluated = true
}
