package main

import (
	ui "github.com/gizak/termui"
)

type LogsWidget struct {
	latests []string
	max     int
	widget  *ui.List
}

func (l *LogsWidget) AppendLog(log string) {
	l.latests = append(l.latests[1:l.max], log)
}

func (l *LogsWidget) changeUiList() {
	l.widget.Items = l.latests
}

func (l *LogsWidget) Render() {
	l.changeUiList()
	Rerender()
}

func (l *LogsWidget) ChangeMax(max int) {
	l.max = max
	l.widget.Height = max + 2

	latestLen := len(l.latests)

	if l.max > latestLen {
		l.latests = append(make([]string, l.max-latestLen), l.latests...)
	}

	if l.max < latestLen {
		l.latests = l.latests[latestLen-l.max:]
	}
}

func NewLogsWidget(max int) *LogsWidget {
	logs := new(LogsWidget)
	logs.max = max
	logs.latests = make([]string, max)
	logs.widget = ui.NewList()
	logs.widget.Height = max + 2

	logs.changeUiList()

	return logs
}
