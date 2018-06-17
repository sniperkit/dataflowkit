package parse

import (
	// external
	"github.com/sirupsen/logrus"

	// internal
	"github.com/sniperkit/dataflowkit/pkg/logger"
)

var logger *logrus.Logger

func init() {
	logger = log.NewLogger(true)
}

/*
var logger *log.Logger

func init() {
	logger = log.New(os.Stdout, "Parser: ", log.Lshortfile)
}
*/
