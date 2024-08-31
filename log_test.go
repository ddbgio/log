package log

import (
	"log/slog"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	opts := slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: false,
	}
	log, err := New(opts)
	require.NoError(t, err)
	require.NotNil(t, log)
	log.Debug("debug message test", "key", "value")
	log.Info("info message test", "key", "value")
	log.Warn("warn message test", "key", "value")
	log.Error("error message test", "key", "value")
}
