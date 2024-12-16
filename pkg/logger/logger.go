// Package logger menyediakan fungsi logging kustom
package logger

import (
	"log"
	"os"
)

var (
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
)

func init() {
	InfoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// Info logs informational messages
func Info(format string, v ...interface{}) {
	InfoLogger.Printf(format, v...)
}

// Error logs error messages
func Error(format string, v ...interface{}) {
	ErrorLogger.Printf(format, v...)
}

// Fatal logs error message and exits
func Fatal(format string, v ...interface{}) {
	ErrorLogger.Fatalf(format, v...)
}