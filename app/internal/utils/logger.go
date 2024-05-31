package utils

import (
	"log"
	"os"
)

// Logger struct that holds info and error loggers.
type Logger struct {
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
}

// NewLogger creates and returns a new Logger instance.
func NewLogger() *Logger {
	return &Logger{
		InfoLogger:  log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		ErrorLogger: log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}
