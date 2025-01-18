package main

import (
	"github.com/sirupsen/logrus"
	"src/Gin/_init"
)

func main() {

	logger := logrus.New()

	logger.AddHook(&logr.CustomHook{})

	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	logger.SetLevel(logrus.DebugLevel)

	logger.Debug("test")
}
