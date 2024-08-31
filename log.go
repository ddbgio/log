package log

import (
	"log/slog"
	"os"

	"github.com/lmittmann/tint"
	"github.com/mattn/go-isatty"
)

const ISO8601 = "2006-01-02T15:04:05.000Z"

func New(opts slog.HandlerOptions) (*slog.Logger, error) {
	w := os.Stderr
	logger := slog.New(
		tint.NewHandler(w, &tint.Options{
			Level:      opts.Level,
			AddSource:  opts.AddSource,
			NoColor:    !isatty.IsTerminal(w.Fd()),
			TimeFormat: ISO8601,
		}),
	)
	return logger, nil
}
