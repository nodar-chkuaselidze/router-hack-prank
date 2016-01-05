package main

import (
	"errors"
	"fmt"
	ui "github.com/gizak/termui"
	"io"
	"os"
	"strconv"
)

type Wordlist struct {
	wordlistInfo   os.FileInfo
	wordlistFD     *os.File
	wordlistName   string
	WordlistWidget *LogsWidget
	Process        *ui.Gauge
	readBytes      int64
}

func (wl *Wordlist) Init(filename string, totalHeight int) {
	wl.wordlistName = filename

	stats, err := os.Stat(filename)

	if os.IsNotExist(err) {
		Panic(errors.New("Could not find wordlist file:" + filename))
	} else if err != nil {
		Panic(err)
	}

	wl.wordlistInfo = stats
	wl.wordlistFD, err = os.Open(filename)

	if err != nil {
		Panic(err)
	}

	wl.WordlistWidget = NewLogsWidget(totalHeight - 3)
	wl.WordlistWidget.widget.Border = false

	wl.Process = ui.NewGauge()
	wl.Process.Height = 1
	wl.Process.Border = false

	wl.UpdateProgress(0)
}

func (wl *Wordlist) Resize(totalHeight int) {
	wl.WordlistWidget.ChangeMax(totalHeight - 3)
	wl.WordlistWidget.Render()
}

func (wl *Wordlist) NextWord() (string, bool) {
	var word string

	_, err := fmt.Fscanln(wl.wordlistFD, &word)
	end := false

	if err == io.EOF {
		return "", true
	}

	if err != nil {
		Panic(err)
	}

	wl.WordlistWidget.AppendLog(word)
	wl.UpdateProgress(wl.readBytes + int64(len(word)+1))
	wl.WordlistWidget.Render()

	return word, end
}

func (wl *Wordlist) UpdateProgress(readBytes int64) {
	wl.readBytes = readBytes
	totalBytes := wl.wordlistInfo.Size()

	wl.Process.Percent = int(float64(readBytes) / float64(totalBytes) * 100)
	wl.Process.Label = "{{percent}}% (" + strconv.FormatInt(readBytes, 10) + " / " + strconv.FormatInt(totalBytes, 10) + ")"
}

func NewWordList(filename string, totalHeight int) *Wordlist {
	wordlist := new(Wordlist)
	wordlist.Init(filename, totalHeight)

	return wordlist
}
