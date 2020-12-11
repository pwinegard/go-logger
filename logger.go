package logger

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	LogLevel Level = 2
)

type Level int

const (
	trace Level = iota
	debug
	info
	warn
	err
	fatal
)

// TRACE trace log level
func TRACE(message string, args ...interface{}) error {
	tmp := fmt.Errorf(message, args...)
	logger(trace, tmp.Error())
	return tmp
}

// DEBUG debug log level
func DEBUG(message string, args ...interface{}) error {
	tmp := fmt.Errorf(message, args...)
	logger(debug, tmp.Error())
	return tmp
}

// INFO info log level
func INFO(message string, args ...interface{}) error {
	tmp := fmt.Errorf(message, args...)
	logger(info, tmp.Error())
	return tmp
}

// WARN warn log level
func WARN(message string, args ...interface{}) error {
	tmp := fmt.Errorf(message, args...)
	logger(warn, tmp.Error())
	return tmp
}

// ERROR error log level
func ERROR(message string, args ...interface{}) error {
	tmp := fmt.Errorf(message, args...)
	logger(err, tmp.Error())
	return tmp
}

// FATAL fatal log level
func FATAL(message string, args ...interface{}) {
	tmp := fmt.Errorf(message, args...)
	logger(fatal, tmp.Error())
	os.Exit(1)
}

// TRACEJSON marshals object into JSON and logs with trace level
func TRACEJSON(object interface{}, prefix string) {
	logJSON(trace, object, prefix)
}

// DEBUGJSON marshals object into JSON and logs with debug level
func DEBUGJSON(object interface{}, prefix string) {
	logJSON(debug, object, prefix)
}

// INFOJSON marshals object into JSON and logs with info level
func INFOJSON(object interface{}, prefix string) {
	logJSON(info, object, prefix)
}

// WARNJSON marshals object into JSON and logs with warn level
func WARNJSON(object interface{}, prefix string) {
	logJSON(warn, object, prefix)
}

func logger(level Level, message string) {
	if level >= LogLevel {
		switch level {
		case trace:
			tmp := fmt.Sprintf("TRACE: %s", message)
			log.Println(tmp)
		case debug:
			tmp := fmt.Sprintf("DEBUG: %s", message)
			log.Println(tmp)
		case info:
			tmp := fmt.Sprintf("INFO: %s", message)
			log.Println(tmp)
		case warn:
			tmp := fmt.Sprintf("WARN: %s", message)
			log.Println(tmp)
		case err:
			tmp := fmt.Sprintf("ERROR: %s", message)
			log.Println(tmp)
		case fatal:
			tmp := fmt.Sprintf("FATAL: %s", message)
			log.Println(tmp)
			os.Exit(1)
		default:
			log.Println(message)
		}
	}
}

func logJSON(level Level, jsonMessage interface{}, prefix string) {
	message, e := json.MarshalIndent(jsonMessage, "", "    ")
	if e != nil {
		logger(err, e.Error())
	}
	logger(level, fmt.Sprintf("%s %s", prefix, string(message[:])))
}

func GetLogLevel(level string) Level {
	switch strings.ToLower(level) {
	case "trace":
		return 0
	case "debug":
		return 1
	case "info":
		return 2
	case "warn":
		return 3
	case "error":
		return 4
	case "fatal":
		return 5
	default:
		return 2
	}
}
