package config

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Lg = logrus.New()

func LogrusConfigInit() {
	logFile, err := os.OpenFile("gok8srun.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0664)
	if err != nil {
		Lg.Fatalf("Can not open log file %v !", err)

	}

	Lg.SetOutput(logFile)
	Lg.SetFormatter(&logrus.JSONFormatter{})
	Lg.SetLevel(logrus.InfoLevel)
	Lg.Info("Logrus started....")
}
