package main

import (
	"log"
	"time"

	"github.com/natefinch/lumberjack"
)

var consumerLogger *log.Logger

func main() {

	consumerLogger = log.Default()

	// Lumberjack Logger
	consumerLogger.SetOutput(&lumberjack.Logger{
		Filename:   "some_loging_filename.log",
		MaxSize:    1024,
		MaxBackups: 30,
		MaxAge:     30,
		Compress:   true,
	})

	for {

		consumerLogger.Printf(time.Now().String())

	}

}
