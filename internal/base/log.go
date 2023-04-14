package base

import (
	"log"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
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

	// logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	logrus.SetOutput(os.Stdout)
	// INFO[2022-05-16T20:51:44+08:00] hello world
	logrus.Info("hello world")

	// logrus.SetReportCaller(true)
	// INFO[2022-05-16T20:51:44+08:00]/Users/yaolei/go/pkg/mod/github.com/sirupsen/logrus@v1.4.2/entry.go:359 github.com/sirupsen/logrus.(*Entry).Logln() hello world
	logrus.Info("hello world")

	standardFields := logrus.Fields{
		"hostname": "Yao's-MacbookPro",
		"appname":  "golang-study",
		"session":  "1ce3f6v",
	}
	logrus.WithFields(standardFields).WithFields(logrus.Fields{"xxx": 1}).Info("My first log from Golang")

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	defer logger.Sync()

	atomicLevel := zap.NewAtomicLevelAt(zap.InfoLevel)
	config := zap.Config{
		Level:       atomicLevel,
		Development: false,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding:         "json",
		EncoderConfig:    zap.NewProductionEncoderConfig(),
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}
	logger2, err := config.Build()
	if err != nil {
		log.Fatalf("can't initialize zap logger2: %v", err)
	}
	defer logger2.Sync()

	http.HandleFunc("log/level", atomicLevel.ServeHTTP)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("listen and serve error: %v", err)
	}
}
