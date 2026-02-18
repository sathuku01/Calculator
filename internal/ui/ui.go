package ui

import (
	"image/color"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)


type customMinSize struct {
	fyne.CanvasObject
	size fyne.Size
}

func (m *customMinSize) MinSize() fyne.Size {
	return m.size
}

func withHeight(obj fyne.CanvasObject, height float32) fyne.CanvasObject {
	return container.New(layout.NewMaxLayout(), &customMinSize{obj, fyne.NewSize(0, height)}, obj)
}

func BuildUI() *fyne.Container {
	Display := widget.NewEntry()
	Display.TextStyle = fyne.TextStyle{Bold: true}
	Display.SetPlaceHolder("0")
	SetDisplay(Display)

	// Colorful Display Area
	displayBG := canvas.NewRectangle(color.NRGBA{33, 37, 41, 255})
	displayArea := container.NewStack(displayBG, withHeight(Display, 100))

	orange := color.NRGBA{255, 165, 0, 255}
	red := color.NRGBA{220, 53, 69, 255}

	btn := func(label string, textColor color.Color, tapped func()) fyne.CanvasObject {
		txt := canvas.NewText(label, textColor)
		txt.TextStyle = fyne.TextStyle{Bold: true}
		txt.Alignment = fyne.TextAlignCenter
		// Stack the text on top of a standard button
		content := container.NewStack(widget.NewButton("", tapped), txt)
		return withHeight(content, 75)
	}

	grid := container.NewGridWithColumns(4,
		btn("7", color.Black, func() { HandleNumber("7") }),
		btn("8", color.Black, func() { HandleNumber("8") }),
		btn("9", color.Black, func() { HandleNumber("9") }),
		btn("/", orange, func() { HandleOperator("/") }),

		btn("4", color.Black, func() { HandleNumber("4") }),
		btn("5", color.Black, func() { HandleNumber("5") }),
		btn("6", color.Black, func() { HandleNumber("6") }),
		btn("*", orange, func() { HandleOperator("*") }),

		btn("1", color.Black, func() { HandleNumber("1") }),
		btn("2", color.Black, func() { HandleNumber("2") }),
		btn("3", color.Black, func() { HandleNumber("3") }),
		btn("-", orange, func() { HandleOperator("-") }),

		btn("C", red, HandleClear),
		btn("0", color.Black, func() { HandleNumber("0") }),
		btn("=", orange, HandleEqual),
		btn("+", orange, func() { HandleOperator("+") }),
	)

	return container.NewVBox(displayArea, grid)
}