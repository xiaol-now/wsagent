package message

import (
	"os"
	"wsagent/internal/logger"
)

var log = logger.New(os.Stdout, logger.DebugLevel)
