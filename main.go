package main

import (
	"flag"
	ui "github.com/gizak/termui"
	"os"
	"time"
)

var LOGS_HEIGHT = 7

var logger = NewLogger("./logs")
var logs = NewLogsWidget(LOGS_HEIGHT - 2)

var WordListFile string
var AdminName string

func init() {
	flag.StringVar(&WordListFile, "wordlist", "./wordlist.txt", "Word List File")
	flag.StringVar(&AdminName, "admin", "admin", "username of admin")
}

func main() {
	flag.Parse()

	err := ui.Init()
	defer ui.Close()

	if err != nil {
		Panic(err)
	}

	ui.Body.AddRows(
		ui.NewRow(
			ui.NewCol(12, 0, logs.widget),
		),
	)

	ui.Body.Align()
	ui.Render(ui.Body)

	getInfo()
	wl := initWordlist()

	ui.Handle("/sys/kbd/q", func(ui.Event) {
		ui.StopLoop()
	})

	ui.Handle("/sys/kbd/C-c", func(ui.Event) {
		ui.StopLoop()
	})

	ui.Handle("/sys/wnd/resize", func(e ui.Event) {
		wl.Resize(ui.TermHeight() - LOGS_HEIGHT)
		Rerender()
	})

	go processWordlist(wl)
	ui.Loop()
}

func initWordlist() *Wordlist {
	logs.AppendLog("Initializing wordlist..")

	wordlist := NewWordList(WordListFile, ui.TermHeight()-LOGS_HEIGHT)

	ui.Body.AddRows(
		ui.NewRow(
			ui.NewCol(12, 0, wordlist.WordlistWidget.widget),
		),
		ui.NewRow(
			ui.NewCol(12, 0, wordlist.Process),
		),
	)

	wordlist.WordlistWidget.Render()
	logs.Render()

	return wordlist
}

func getInfo() {
	cwd, err := os.Getwd()

	if err != nil {
		Panic(err)
	}

	logs.AppendLog("Starting TP-Link Dictionary Attack..")
	logs.AppendLog("CWD: " + cwd)
	logs.AppendLog("Wordlist for use: " + WordListFile)
	logs.Render()
}

func processWordlist(wl *Wordlist) {
	logs.AppendLog("Started Attack...")
	logs.Render()
	_, end := wl.NextWord()

	for end == false {
		_, end = wl.NextWord()
		time.Sleep(20 * time.Millisecond)
	}

}
