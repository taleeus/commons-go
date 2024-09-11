package logconfig

import (
	"log/slog"
	"os"
	"testing"
)

func TestLogger(t *testing.T) {
	ConfigureLogger(os.Stderr, slog.LevelInfo)
	slog.Info("Test!", "foo", "bar")
}
