package parse

import (
	"io"
	"time"

	// external
	"github.com/sirupsen/logrus"

	// internal
	"github.com/sniperkit/dataflowkit/pkg/scrape"
)

// LoggingMiddleware logs Parse Service endpoints
func LoggingMiddleware(logger *logrus.Logger) ServiceMiddleware {
	return func(next Service) Service {
		return loggingMiddleware{next, logger}
	}
}

// Make a new type and wrap into Service interface
// Add logger property to this type
type loggingMiddleware struct {
	Service
	logger *logrus.Logger
}

// Logging Parse Service
func (mw loggingMiddleware) Parse(payload scrape.Payload) (output io.ReadCloser, err error) {
	defer func(begin time.Time) {
		output, err = mw.Service.Parse(payload)
		url := payload.Request.GetURL()
		if err != nil {
			mw.logger.WithFields(
				logrus.Fields{
					"err":     err,
					"fetcher": payload.FetcherType,
					"took":    time.Since(begin),
				}).Error("Parse URL: ", url)
		} else {
			mw.logger.WithFields(
				logrus.Fields{
					"fetcher": payload.FetcherType,
					"took":    time.Since(begin),
				}).Info("Parse URL: ", url)
		}
	}(time.Now())
	return
}
