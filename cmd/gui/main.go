package main

import (
	"calculator/internal/ui"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2"
)

func main() {
	a := app.New()
	w := a.NewWindow("Calculator")
	w.Resize(fyne.NewSize(300, 400))

	// Build the UI from ui package
	w.SetContent(ui.BuildUI())
	w.ShowAndRun()
}
