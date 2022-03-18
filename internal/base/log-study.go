package base

import (
	logrus "github.com/sirupsen/logrus"
	"log"
)

func StudyLog() {
	log.Println("hello world")

	logrus.SetFormatter(&logrus.JSONFormatter{})
}
