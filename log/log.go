package log

import (
	"io"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

var logger = log.New()

func init() {
	logger.SetFormatter(NewFormatter())
	logger.SetOutput(os.Stdout)
	logger.SetReportCaller(true)
}

func SetLevel(level log.Level) {
	logger.SetLevel(level)
}

func SetOutput(console bool, logFilePath string) {
	writers := make([]io.Writer, 0)

	if console {
		writers = append(writers, os.Stdout)
	}

	if logFilePath != "" {
		err := os.MkdirAll(filepath.Dir(logFilePath), os.ModePerm)
		if err != nil {
			panic(err)
		}

		file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			panic(err)
		}
		writers = append(writers, file)
	}

	logger.SetOutput(io.MultiWriter(writers...))
}

func Printf(format string, args ...interface{}) {
	logger.Printf(format, args...)
}

func Tracef(format string, args ...interface{}) {
	logger.Tracef(format, args...)
}

func Debugf(format string, args ...interface{}) {
	logger.Debugf(format, args...)
}

func Infof(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

func Warnf(format string, args ...interface{}) {
	logger.Warnf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	logger.Errorf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	logger.Fatalf(format, args...)
}

func Panicf(format string, args ...interface{}) {
	logger.Panicf(format, args...)
}
