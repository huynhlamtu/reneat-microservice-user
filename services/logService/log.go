package logService

import (
	"github.com/sirupsen/logrus"
	"os"
)

func NewLogrus() {
	if os.Getenv("APP_ENV") != "production" {
		logrus.SetLevel(logrus.TraceLevel)
	}

	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetReportCaller(true)
}
