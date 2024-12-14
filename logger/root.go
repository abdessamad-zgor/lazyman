package logger

import (
	"log"
	"os"
	"path"
)

var logger *log.Logger

func init() {
	_, ok := os.LookupEnv("LAZYMAN_DEV")
	if ok {
		cwd, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		logFile, err := os.OpenFile(path.Join(cwd, "./logs/dev.log"), os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
		if err != nil {
			panic(err)
		}
		logger = log.New(logFile, "", log.Ltime)
	}
}

func Info(v ...any) {
	if logger != nil {
		logger.Println("INFO: ", v)
	}
}

func Infof(format string, v ...any) {
	if logger != nil {
		logger.Printf(format, v)
	}
}

func Error(v ...any) {
	if logger != nil {
		logger.Panic("Error: ", v)
	}
}

func Close() {
	if logger != nil {
		file := logger.Writer().(*os.File)
		file.Close()
	}
}
