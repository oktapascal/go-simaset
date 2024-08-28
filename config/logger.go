package config

import (
	"fmt"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type LoggerFileHook struct {
	file      *os.File
	flag      int
	chmod     os.FileMode
	formatter *logrus.TextFormatter
}

func NewLoggerFileHook(file string, flag int, chmod os.FileMode) (*LoggerFileHook, error) {
	textFormatter := &logrus.TextFormatter{
		DisableColors:   true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	}

	logFile, err := os.OpenFile(file, flag, chmod)
	if err != nil {
		fmt.Printf("Failed to open log file: %s\n", err)
		return nil, err
	}

	return &LoggerFileHook{
		file:      logFile,
		flag:      flag,
		chmod:     chmod,
		formatter: textFormatter,
	}, nil
}

func createLogsDir() error {
	logsDir := "storage/logs"

	if _, err := os.Stat(logsDir); os.IsNotExist(err) {
		return os.MkdirAll(logsDir, 0755)
	}

	return nil
}

func getLogFileName() string {
	currentDate := time.Now().Format("2006-01-02")

	logFilePath := filepath.Join("storage/logs", currentDate+".log")

	return logFilePath
}

func (hook *LoggerFileHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
		logrus.WarnLevel,
		logrus.InfoLevel,
		logrus.DebugLevel,
	}
}

func (hook *LoggerFileHook) Fire(entry *logrus.Entry) error {
	var format, err = hook.formatter.Format(entry)
	var str = string(format)

	_, err = hook.file.WriteString(str)
	if err != nil {
		fmt.Printf("Failed to write log file: %s\n", err)
		return err
	}

	return nil
}

func CreateLoggers(request *http.Request) *logrus.Entry {
	logger := logrus.New()

	logger.SetOutput(os.Stdout)
	logger.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logger.SetLevel(logrus.DebugLevel)

	var errDirectory = createLogsDir()
	if errDirectory != nil {
		fmt.Println(errDirectory)
		return nil
	}

	logFileName := getLogFileName()

	var loggerFileHook, err = NewLoggerFileHook(logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		logger.Hooks.Add(loggerFileHook)
	}

	if request == nil {
		return logger.WithFields(logrus.Fields{})
	}

	return logger.WithFields(logrus.Fields{
		"request_id": middleware.GetReqID(request.Context()),
		"at":         time.Now().Format("2006-01-02 15:04:05"),
		"method":     request.Method,
		"uri":        request.RequestURI,
		"ip":         request.RemoteAddr,
	})
}
