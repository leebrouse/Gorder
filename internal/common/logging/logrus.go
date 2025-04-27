package logging

import (
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
)

// config logrus
func Init() {
	SetFormatter(logrus.StandardLogger())
	logrus.SetLevel(logrus.DebugLevel)
}

func SetFormatter(logger *logrus.Logger) {
	logger.SetFormatter(&logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyLevel: "severity",
			logrus.FieldKeyTime:  "time",
			logrus.FieldKeyMsg:   "message",
		},
	})
	if isLocal, _ := strconv.ParseBool(os.Getenv("LOCAL_ENV")); isLocal {
		//logrus.SetFormatter(&prefixed.TextFormatter{
		//	ForceFormatting: true,
		//})
	}
}
