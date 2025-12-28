package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func panel(title string, body fyne.CanvasObject) fyne.CanvasObject {
	header := canvas.NewText(title, theme.Color(theme.ColorNameForeground))
	header.TextStyle = fyne.TextStyle{Bold: true}
	header.TextSize = theme.TextSize() * 1.15

	bg := canvas.NewRectangle(theme.Color(theme.ColorNameInputBackground))
	bg.StrokeColor = theme.Color(theme.ColorNameSeparator)
	bg.StrokeWidth = 1

	top := container.NewVBox(header, widget.NewSeparator())
	content := container.NewBorder(top, nil, nil, nil, body)
	return container.NewMax(bg, container.NewPadded(content))
}

func chip(text string, fill color.Color) fyne.CanvasObject {
	label := widget.NewLabelWithStyle(text, fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	label.Wrapping = fyne.TextWrapWord

	bg := canvas.NewRectangle(fill)
	bg.StrokeColor = theme.Color(theme.ColorNameSeparator)
	bg.StrokeWidth = 1

	row := container.NewHBox(layout.NewSpacer(), label, layout.NewSpacer())
	return container.NewMax(bg, container.NewPadded(row))
}

func fieldLabel(text string) *widget.Label {
	l := widget.NewLabelWithStyle(text, fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	// Keep labels in a single line so border/grid layouts don't collapse the label column.
	l.Wrapping = fyne.TextWrapOff
	return l
}

func formRow(label string, field fyne.CanvasObject) fyne.CanvasObject {
	l := fieldLabel(label)
	return container.NewBorder(nil, nil, l, nil, field)
}

func metricTile(title string, value fyne.CanvasObject) fyne.CanvasObject {
	titleLabel := widget.NewLabelWithStyle(title, fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	titleLabel.Wrapping = fyne.TextWrapOff

	return metricTileWithHeader(titleLabel, value)
}

func metricTileWithIcon(title string, icon fyne.Resource, value fyne.CanvasObject) fyne.CanvasObject {
	titleLabel := widget.NewLabelWithStyle(title, fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	titleLabel.Wrapping = fyne.TextWrapOff

	header := container.NewHBox(widget.NewIcon(icon), titleLabel)
	return metricTileWithHeader(header, value)
}

func metricTileWithHeader(header fyne.CanvasObject, value fyne.CanvasObject) fyne.CanvasObject {
	bg := canvas.NewRectangle(theme.Color(theme.ColorNameButton))
	bg.StrokeColor = theme.Color(theme.ColorNameSeparator)
	bg.StrokeWidth = 1

	return container.NewMax(bg, container.NewPadded(container.NewVBox(header, value)))
}
