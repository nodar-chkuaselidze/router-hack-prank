package main

import (
	"fmt"
	"os"
)

type Logger struct {
	filename string
	logFD    *os.File
	info     os.FileInfo
}

func (l *Logger) Init() {
	_, err := os.Stat(l.filename)

	if os.IsNotExist(err) {
		l.logFD, err = os.Create(l.filename)

		if err != nil {
			Panic(err)
		}

		return
	}

	//if exists just append
	open, err := os.OpenFile(l.filename, os.O_RDWR, 0666)

	open.Seek(0, 2)

	if err != nil {
		Panic(err)
	}

	l.logFD = open
}

func (l *Logger) Log(data ...interface{}) {
	fmt.Fprintln(l.logFD, data)
}

func NewLogger(filename string) *Logger {
	logger := new(Logger)
	logger.filename = filename
	logger.Init()

	return logger
}
