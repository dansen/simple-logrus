package log

import (
	"io"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

var logger = log.New()

func init() {
	setLogger()
	logger.SetFormatter(NewFormatter())
	logger.SetOutput(os.Stdout)
	logger.SetReportCaller(true)
}

// log
var (
	Printf  func(format string, v ...interface{})
	Print   func(v ...interface{})
	Println func(v ...interface{})

	Fatalf  func(format string, v ...interface{})
	Fatal   func(v ...interface{})
	Fatalln func(v ...interface{})

	Panicf  func(format string, v ...interface{})
	Panic   func(v ...interface{})
	Panicln func(v ...interface{})

	Tracef  func(format string, v ...interface{})
	Trace   func(v ...interface{})
	Traceln func(v ...interface{})

	Debugf  func(format string, v ...interface{})
	Debug   func(v ...interface{})
	Debugln func(v ...interface{})

	Infof  func(format string, v ...interface{})
	Info   func(v ...interface{})
	Infoln func(v ...interface{})

	Warnf  func(format string, v ...interface{})
	Warn   func(v ...interface{})
	Warnln func(v ...interface{})

	Warningf  func(format string, v ...interface{})
	Warning   func(v ...interface{})
	Warningln func(v ...interface{})

	Errorf  func(format string, v ...interface{})
	Error   func(v ...interface{})
	Errorln func(v ...interface{})
)

func setLogger() {
	Tracef = logger.Tracef
	Trace = logger.Trace
	Traceln = logger.Traceln

	Debugf = logger.Debugf
	Debug = logger.Debug
	Debugln = logger.Debugln

	Infof = logger.Infof
	Info = logger.Info
	Infoln = logger.Infoln

	Printf = logger.Printf
	Print = logger.Print
	Println = logger.Println

	Warnf = logger.Warnf
	Warn = logger.Warn
	Warnln = logger.Warnln

	Warningf = logger.Warningf
	Warning = logger.Warning
	Warningln = logger.Warningln

	Errorf = logger.Errorf
	Error = logger.Error
	Errorln = logger.Errorln

	Fatalf = logger.Fatalf
	Fatal = logger.Fatal
	Fatalln = logger.Fatalln

	Panicf = logger.Panicf
	Panic = logger.Panic
	Panicln = logger.Panicln
}

func SetLevelName(level string) {
	switch level {
	case "info":
		logger.SetLevel(log.InfoLevel)
	case "debug":
		logger.SetLevel(log.DebugLevel)
	case "warn":
		logger.SetLevel(log.WarnLevel)
	case "error":
		logger.SetLevel(log.ErrorLevel)
	case "fatal":
		logger.SetLevel(log.FatalLevel)
	}
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
