package initalizer

import (
	"log"
	"os"
	"time"
)

var logFile *os.File

func SetupLog() error {
	logDir := "logs"
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		os.Mkdir(logDir, os.ModePerm)
	}

	date := time.Now().Format("2006-01-02")
	var err error
	logFile, err = os.OpenFile("logs/"+date+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
		return err
	}

	log.SetOutput(logFile)

	return nil
}

func CloseLog() {
	if logFile != nil {
		logFile.Close()
	}
}
