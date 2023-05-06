package log

import (
	"bytes"
	"errors"
	"fmt"
	"runtime"
	"strings"
	"time"

	logrus "github.com/sirupsen/logrus"
)

type HookContextKey string

type TextFormatter struct {
	logrus.TextFormatter
	TimeLocation *time.Location
}

var (
	ErrFormatterNotFound     = errors.New("formatter not found")
	ErrMethodNotValid        = errors.New("method not valid")
	ErrFormatOptionsNotFound = errors.New("format options not found")
)

func NewFormatter() *TextFormatter {
	formatter := &TextFormatter{
		TextFormatter: logrus.TextFormatter{
			ForceColors:   true,
			FullTimestamp: true,
			CallerPrettyfier: func(f *runtime.Frame) (string, string) {
				s := f.File
				fileIndex := strings.LastIndex(s, "/")
				packageIndex := strings.LastIndex(s[:fileIndex], "/")
				atIndex := strings.LastIndex(s[packageIndex+1:fileIndex], "@")
				var packageFile string
				if atIndex >= 0 {
					packageFile = s[packageIndex+1:]
				} else {
					packageFile = s[packageIndex+1 : fileIndex][atIndex+1:] + s[fileIndex:]
				}

				funcIndex := strings.LastIndex(f.Function, ".")
				structIndex := strings.LastIndex(f.Function[:funcIndex], ".")
				var function string
				if structIndex >= 0 {
					function = f.Function[structIndex+1:]
				} else {
					function = f.Function[funcIndex+1:]
				}

				return fmt.Sprintf("%s:", function), fmt.Sprintf(" %s:%d", packageFile, f.Line)
			},
		},
		TimeLocation: time.Local,
	}
	return formatter
}

func (f *TextFormatter) generateTimeFormat(date bool, time bool, nanosecond bool, timezone bool) string {
	var buf bytes.Buffer
	if date {
		if buf.Len() > 0 {
			buf.WriteByte(' ')
		}

		buf.WriteString("2006-01-02")
	}

	if time {
		if buf.Len() > 0 {
			buf.WriteByte(' ')
		}

		buf.WriteString("15:04:05")
	}

	if nanosecond {
		buf.WriteString(".000")
	}

	if timezone {
		if buf.Len() > 0 {
			buf.WriteByte(' ')
		}

		buf.WriteString("-0700")
	}

	return buf.String()
}

func (f *TextFormatter) generateFile(frame *runtime.Frame) string {
	s := frame.File
	fileIndex := strings.LastIndex(s, "/")
	packageIndex := strings.LastIndex(s[:fileIndex], "/")
	atIndex := strings.LastIndex(s[packageIndex+1:fileIndex], "@")

	if atIndex >= 0 {
		return fmt.Sprintf(" %s:%d", s[packageIndex+1:], frame.Line)
	}

	return fmt.Sprintf(" %s%s:%d", s[packageIndex+1 : fileIndex][atIndex+1:], s[fileIndex:], frame.Line)
}

func (f *TextFormatter) generateFunction(frame *runtime.Frame) string {
	s := frame.Function

	return fmt.Sprintf(" (%s)", s[strings.LastIndex(s, "/")+1:])
}

func (f *TextFormatter) generateCallerPrettierfier(file bool, function bool) func(*runtime.Frame) (string, string) {
	switch {
	case file && function:
		return func(frame *runtime.Frame) (string, string) {
			return f.generateFile(frame) + f.generateFunction(frame) + ":", ""
		}
	case file:
		return func(frame *runtime.Frame) (string, string) {
			return f.generateFile(frame) + ":", ""
		}
	case function:
		return func(frame *runtime.Frame) (string, string) {
			return f.generateFunction(frame) + ":", ""
		}
	default:
		return func(frame *runtime.Frame) (string, string) {
			return "", ""
		}
	}
}

func (f *TextFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	f.TextFormatter.DisableTimestamp = false
	f.TextFormatter.TimestampFormat = f.generateTimeFormat(true, true, true, false)
	f.TextFormatter.CallerPrettyfier = f.generateCallerPrettierfier(true, false)

	entry.Time = entry.Time.In(time.Local)
	data, err := f.TextFormatter.Format(entry)
	if err != nil {
		return nil, err
	}

	return data, nil
}
