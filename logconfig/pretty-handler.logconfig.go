// Thanks to https://betterstack.com/community/guides/logging/logging-in-go/#creating-custom-handlers
package logconfig

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"log/slog"
	"runtime"
	"strings"

	"github.com/fatih/color"
)

type PrettyHandlerOptions struct {
	SlogOpts slog.HandlerOptions
}

// PrettyHandler is a handler for slogger that prints additional information (file, line)
// with pretty colors and formatting
type PrettyHandler struct {
	slog.Handler
	l *log.Logger
}

func (h *PrettyHandler) Handle(ctx context.Context, r slog.Record) error {
	var level string

	switch r.Level {
	case slog.LevelDebug:
		level = color.MagentaString("DBG")
	case slog.LevelInfo:
		level = color.BlueString("INF")
	case slog.LevelWarn:
		level = color.YellowString("WRN")
	case slog.LevelError:
		level = color.RedString("ERR")
	}

	fields := make(map[string]interface{}, r.NumAttrs())
	r.Attrs(func(a slog.Attr) bool {
		fields[a.Key] = a.Value.Any()

		return true
	})

	b, err := json.MarshalIndent(fields, "", "  ")
	if err != nil {
		return err
	}

	filelineStr := "?:?"
	if r.PC != 0 {
		frame, _ := runtime.CallersFrames([]uintptr{r.PC}).Next()

		funcPaths := strings.Split(frame.Function, ".")
		shortFunc := funcPaths[len(funcPaths)-1]

		filePaths := strings.Split(frame.File, "/")
		shortFile := filePaths[len(filePaths)-1]

		filelineStr = fmt.Sprintf("%s:%d %s()", shortFile, frame.Line, shortFunc)
	}

	timeStr := r.Time.Format("[15:05:05.000]")
	msg := color.CyanString(r.Message)

	h.l.Println(timeStr, filelineStr, level, msg, color.WhiteString(string(b)))

	return nil
}

// NewPrettyHandler returns a new PrettyHandler
func NewPrettyHandler(
	out io.Writer,
	opts PrettyHandlerOptions,
) *PrettyHandler {
	h := &PrettyHandler{
		Handler: slog.NewJSONHandler(out, &opts.SlogOpts),
		l:       log.New(out, "", 0),
	}

	return h
}
