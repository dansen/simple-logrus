package log

import (
	logrus "github.com/sirupsen/logrus"
)

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{})
}

func Infof(format string, args ...interface{}) {
	logrus.Infof(format, args...)
}
