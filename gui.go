package main

import (
	ui "github.com/gizak/termui"
)

func Rerender() {
	ui.Body.Width = ui.TermWidth()
	ui.Body.Align()
	ui.Render(ui.Body)
}

func Panic(err error) {
	ui.StopLoop()
	panic(err)
}
