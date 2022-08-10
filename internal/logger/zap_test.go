package logger

import (
	"os"
	"testing"
)

func TestLogger_Debug(t *testing.T) {
	log := New(os.Stdout, DebugLevel)
	log.Debug("test", String("key", "value"))
	log.Info("test", String("key", "value"))
	log.Error("test", String("key", "value"))
}
