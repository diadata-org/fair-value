package scrapers

import (
	"github.com/sirupsen/logrus"
)

// TO DO: Add log levels.

var log *logrus.Logger

func init() {
	log = logrus.New()
}
