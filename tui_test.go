package log

import (
	"log/slog"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTUI(t *testing.T) {
	tui, err := NewTUI(slog.HandlerOptions{
		Level: slog.LevelDebug,
	})
	require.NoError(t, err)
	require.NotNil(t, tui)
	tui.Debug("debug message test", "key1", "value1", "key2", "value2")
	tui.Info("info message test", "key", "value", "key2", "value2")
	tui.Warn("warn message test", "key", "value", "key2", "value2")
	tui.Error("error message test", "key", "value", "key2", "value2")
}

func TestHeaders(t *testing.T) {
	t.Run("invalid - odd number of fields", func(t *testing.T) {
		keys, values, err := kvParse("key", "value", "valuelessKey")
		require.Error(t, err)
		require.Empty(t, keys)
		require.Empty(t, values)
	})
	t.Run("valid", func(t *testing.T) {
		keys, values, err := kvParse("key", "value", "key2", "value2")
		require.NoError(t, err)
		require.NotEmpty(t, keys)
		require.NotEmpty(t, values)
		t.Logf("keys: %s", keys)
		t.Logf("vas: %s", values)
	})
}

func TestTermInfo(t *testing.T) {
	w, h := termInfo()
	require.Greater(t, w, 0)
	require.Greater(t, h, 0)
}
