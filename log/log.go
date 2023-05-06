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
	Tracef = logger.Printf
	Trace = logger.Print
	Traceln = logger.Println

	Debugf = logger.Printf
	Debug = logger.Print
	Debugln = logger.Println

	Infof = logger.Printf
	Info = logger.Print
	Infoln = logger.Println

	Printf = logger.Printf
	Print = logger.Print
	Println = logger.Println

	Warnf = logger.Printf
	Warn = logger.Print
	Warnln = logger.Println

	Warningf = logger.Printf
	Warning = logger.Print
	Warningln = logger.Println

	Errorf = logger.Printf
	Error = logger.Print
	Errorln = logger.Println

	Fatalf = logger.Fatalf
	Fatal = logger.Fatal
	Fatalln = logger.Fatalln

	Panicf = logger.Panicf
	Panic = logger.Panic
	Panicln = logger.Panicln
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
