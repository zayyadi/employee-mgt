package middleware

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/sirupsen/logrus"
)

var (
	httpDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "http_request_duration_seconds",
		Help:    "Duration of HTTP requests.",
		Buckets: prometheus.DefBuckets,
	}, []string{"path", "method", "status_code"})
)

// Logger is a Gin middleware for structured logging with logrus and prometheus metrics
func Logger(logger *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Generate request ID
		requestID := uuid.New().String()
		c.Set("requestID", requestID)
		c.Writer.Header().Set("X-Request-ID", requestID)

		// Create a logger entry for the request
		entry := logger.WithFields(logrus.Fields{
			"requestID": requestID,
			"method":    c.Request.Method,
			"path":      path,
		})
		c.Set("logger", entry)

		c.Next()

		// After request
		latency := time.Since(start)
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		method := c.Request.Method
		if raw != "" {
			path = path + "?" + raw
		}

		// Update prometheus metrics
		httpDuration.With(prometheus.Labels{
			"path":        path,
			"method":      method,
			"status_code": strconv.Itoa(statusCode),
		}).Observe(latency.Seconds())

		logFields := logrus.Fields{
			"statusCode": statusCode,
			"latency":    latency.String(),
			"clientIP":   clientIP,
			"method":     method,
			"path":       path,
		}

		if len(c.Errors) > 0 {
			// Append error field if there are errors
			entry.WithFields(logFields).Error(c.Errors.String())
		} else {
			if statusCode > 499 {
				entry.WithFields(logFields).Error("Server error")
			} else if statusCode > 399 {
				entry.WithFields(logFields).Warn("Client error")
			} else {
				entry.WithFields(logFields).Info("Request handled")
			}
		}
	}
}
