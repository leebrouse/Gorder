package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"time"
)

func StructuredLog(l *logrus.Entry) gin.HandlerFunc {
	return func(c *gin.Context) {
		//first todo
		logrus.Info("Request In")
		now := time.Now()
		c.Next()

		// back todo
		elapsed := time.Since(now).Milliseconds()
		l.WithFields(logrus.Fields{
			"time_elapsed_ms": elapsed,
			"request_uri":     c.Request.RequestURI,
			"client_ip":       c.ClientIP(),
			"full_path":       c.FullPath(),
		}).Info("request_out")
	}
}
