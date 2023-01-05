package main

import (
	"fmt"
	"log"
	"os"
)

var LOGFILE_1 = "./tmp/go1.log"
var LOGFILE_2 = "./tmp/go2.log"

type Log struct {
	logPath []string
}

func main() {
	log := Log{
		logPath: []string{LOGFILE_1, LOGFILE_2},
	}
	log.Println("Hello there!")
	log.Println("Another log entry!")
	log.Println("Another log from vim!")
}

func (l *Log) Println(msgs ...any) {
	for i := 0; i < len(l.logPath); i++ {
		f, err := os.OpenFile(l.logPath[i], os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println(err)
		}
		defer f.Close()
		iLog := log.New(f, fmt.Sprintf("Logger: %d - ", i), log.LstdFlags)
		iLog.Println(msgs...)
	}
}
