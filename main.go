package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.NewWithID("study.metanit.go")
	w := a.NewWindow("Go — учебная практика, главы 2–5")
	w.SetContent(buildUI())
	w.Resize(fyne.NewSize(1000, 760))
	w.ShowAndRun()
}
