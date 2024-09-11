package logconfig

import (
	"io"

	"log/slog"
)

// ConfigureLogger configures a PrettyHandler with the given Writer and log Level.
//
//	ConfigureLogger(os.Stderr, slog.LevelDebug)
func ConfigureLogger(out io.Writer, level slog.Level) {
	opts := PrettyHandlerOptions{
		SlogOpts: slog.HandlerOptions{
			Level: level,
		},
	}

	handler := NewPrettyHandler(out, opts)
	logger := slog.New(handler)

	slog.SetDefault(logger)
}
