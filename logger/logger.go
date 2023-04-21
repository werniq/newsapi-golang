package logger

import (
	"log"
	"os"
)

// NewLogger returns a new logger
func NewLogger() *log.Logger {
	return log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
}
