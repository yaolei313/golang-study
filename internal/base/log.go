package base

import (
	logrus "github.com/sirupsen/logrus"
	"log"
)

func StudyLog() {
	/*
	1.Make calls to the logger from within your main application process, not within goroutines.
	2.Write logs from your application to a local file, even if you’ll ship them to a central platform later.
	3.Standardize your logs with a set of predefined messages.
	4.Send your logs to a central platform so you can analyze and aggregate them.
	5.Use HTTP headers and unique IDs to log user behavior across microservices.
	*/
	log.Println("hello world")

	logrus.SetFormatter(&logrus.JSONFormatter{})

	standardFields := logrus.Fields{
		"hostname": "staging-1",
		"appname":  "study-log",
		"session":  "1ce3f6v",
	}
	logrus.WithFields(standardFields).WithFields(logrus.Fields{"xxx": 1}).Info("My first log from Golang")
}
