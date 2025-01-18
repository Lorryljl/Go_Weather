package main

import (
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:  true,
		PadLevelText: true,
	})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)
	logrus.Debug("Debug Message")
}
