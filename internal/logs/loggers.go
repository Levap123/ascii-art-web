package logs

import (
	"fmt"
	"log"
	"os"
)

const logPath = "logs.log"

var (
	Logger       *log.Logger
	ErrLogger    *log.Logger
	LoggerOut    *log.Logger
	ErrLoggerOut *log.Logger
)

func init() {
	logFile, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0o666)
	if err != nil {
		fmt.Println("errorMessage opening file:", err)
		os.Exit(1)
	}
	Logger = log.New(logFile, "Info:  ", log.Ldate|log.Ltime)
	ErrLogger = log.New(logFile, "Error: ", log.Ldate|log.Ltime)
	LoggerOut = log.New(os.Stdout, "\033[32mInfo:\033[0m  ", log.Ldate|log.Ltime)
	ErrLoggerOut = log.New(os.Stdout, "\033[31mError:\033[0m ", log.Ldate|log.Ltime)
}

func LoggerOk(urlPath, httpMethod, handlerFunc string, statucCode int) {
	Logger.Printf("%-15s | %-6s | %s | %d |", urlPath, httpMethod, handlerFunc, statucCode)
	LoggerOut.Printf("%-15s | %-6s | %s | %d |", urlPath, httpMethod, handlerFunc, statucCode)
}

func LoggerErr(urlPath, httpMethod, handlerFunc, errText string, errStatusCode int) {
	ErrLoggerOut.Printf("%-15s | %-6s | %s | \033[31m%d\033[0m | %s", urlPath, httpMethod, handlerFunc, errStatusCode, errText)
	ErrLogger.Printf("%-15s | %-6s | %s | %d | %s", urlPath, httpMethod, handlerFunc, errStatusCode, errText)
}
