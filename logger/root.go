package logger

import (
	"log"
	"os"
)

var logger *log.Logger

func init() {
    logFile, err := os.OpenFile("./logs/dev.log", os.O_RDWR|os.O_CREATE, 664)
    if err != nil {
        panic(err)
    }
    logger = log.New(logFile, "", log.Ltime)
}

func Info(v ...any) {
    logger.Println("INFO: ", v)
}

func Infof(format string, v ...any) {
    logger.Printf(format, v)
}

func Error(v ...any) {
    logger.Panic("Error: ", v)
}
