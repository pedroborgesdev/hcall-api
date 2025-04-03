package logger

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"hcall/api/config"
)

// Códigos ANSI para cores de texto
const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPurple = "\033[35m"
	colorCyan   = "\033[36m"
	colorWhite  = "\033[37m"
)

// LogLevel represents the severity level of a log message
type LogLevel string

const (
	// Log levels with text colors
	LevelDebug   LogLevel = colorBlue + "DEBUG" + colorReset
	LevelInfo    LogLevel = colorGreen + "INFO" + colorReset
	LevelWarning LogLevel = colorYellow + "WARNING" + colorReset
	LevelError   LogLevel = colorRed + "ERROR" + colorReset
	LevelFatal   LogLevel = colorPurple + "FATAL" + colorReset
)

// Logger handles logging with different levels and formats
type Logger struct {
	useJSON bool
	debug   bool
	colors  bool // Controle para cores
}

var (
	instance    *Logger
	initialized bool
)

// InitLogger initializes the logger instance after config is loaded
func InitLogger() {
	if initialized {
		return
	}

	instance = &Logger{
		useJSON: os.Getenv("LOG_FORMAT") == "json",
		debug:   config.AppConfig.Debug,
		colors:  os.Getenv("LOG_COLORS") != "false", // Ativado por padrão
	}
	initialized = true
}

// GetLogger returns the singleton logger instance
func GetLogger() *Logger {
	if !initialized {
		return &Logger{
			useJSON: false,
			debug:   true,
			colors:  true,
		}
	}
	return instance
}

// getLevelWithColor retorna o nível formatado
func (l *Logger) getLevelWithColor(level LogLevel) string {
	if !l.colors {
		switch level {
		case LevelDebug:
			return "DEBUG"
		case LevelInfo:
			return "INFO"
		case LevelWarning:
			return "WARNING"
		case LevelError:
			return "ERROR"
		case LevelFatal:
			return "FATAL"
		}
	}
	return string(level)
}

// log writes a log message with the given level and fields
func (l *Logger) log(level LogLevel, message string, fields map[string]interface{}) {
	if !config.AppConfig.Debug {
		return
	}

	if fields == nil {
		fields = make(map[string]interface{})
	}

	fields["timestamp"] = time.Now().Format(time.RFC3339)
	fields["level"] = level
	fields["message"] = message

	if l.useJSON {
		jsonData, err := json.Marshal(fields)
		if err != nil {
			fmt.Printf("Error marshaling log to JSON: %v\n", err)
			return
		}
		fmt.Println(string(jsonData))
	} else {
		levelStr := l.getLevelWithColor(level)

		// Formato texto melhorado com campos adicionais
		logMsg := fmt.Sprintf("[%s] %s %s",
			fields["timestamp"],
			levelStr,
			message)

		// Adiciona campos extras (exceto os já usados)
		for k, v := range fields {
			if k != "timestamp" && k != "level" && k != "message" {
				logMsg += fmt.Sprintf(" %s=%v", k, v)
			}
		}

		fmt.Println(logMsg)
	}
}

// Debug logs a debug message
func (l *Logger) Debug(message string, fields map[string]interface{}) {
	if fields == nil {
		fields = make(map[string]interface{})
	}
	l.log(LevelDebug, message, fields)
}

// Info logs an info message
func (l *Logger) Info(message string, fields map[string]interface{}) {
	if fields == nil {
		fields = make(map[string]interface{})
	}
	l.log(LevelInfo, message, fields)
}

// Warning logs a warning message
func (l *Logger) Warning(message string, fields map[string]interface{}) {
	if fields == nil {
		fields = make(map[string]interface{})
	}
	l.log(LevelWarning, message, fields)
}

// Error logs an error message
func (l *Logger) Error(message string, fields map[string]interface{}) {
	if fields == nil {
		fields = make(map[string]interface{})
	}
	l.log(LevelError, message, fields)
}

// Fatal logs a fatal message and exits
func (l *Logger) Fatal(message string, fields map[string]interface{}) {
	if fields == nil {
		fields = make(map[string]interface{})
	}
	l.log(LevelFatal, message, fields)
	os.Exit(1)
}

// Convenience functions for package-level logging
func Debug(message string, fields map[string]interface{}) {
	GetLogger().Debug(message, fields)
}

func Info(message string, fields map[string]interface{}) {
	GetLogger().Info(message, fields)
}

func Warning(message string, fields map[string]interface{}) {
	GetLogger().Warning(message, fields)
}

func Error(message string, fields map[string]interface{}) {
	GetLogger().Error(message, fields)
}

func Fatal(message string, fields map[string]interface{}) {
	GetLogger().Fatal(message, fields)
}
