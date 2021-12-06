package logger

import "github.com/sirupsen/logrus"

var JSONFormatter = &logrus.JSONFormatter{
	TimestampFormat: "2006-01-02 15:04:05",
}

var TextFormatter = &logrus.TextFormatter{
	TimestampFormat: "2006-01-02 15:04:05",
	FullTimestamp: true,
}