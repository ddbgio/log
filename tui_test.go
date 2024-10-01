package log

import (
	"log/slog"
	"os"
	"testing"

	"github.com/mattn/go-isatty"
	"github.com/stretchr/testify/require"
)

func TestTUI(t *testing.T) {
	tui, err := NewTUI(slog.HandlerOptions{
		Level: slog.LevelDebug,
	})
	require.NoError(t, err)
	require.NotNil(t, tui)
	t.Run("print all message types", func(t *testing.T) {
		tui.Debug("debug message test", "key1", "value1", "key2", "value2")
		tui.Info("info message test", "key", "value", "key2", "value2")
		tui.Warn("warn message test", "key", "value", "key2", "value2")
		tui.Error("error message test", "key", "value", "key2", "value2")
	})
	t.Run("overflow", func(t *testing.T) {
		tui.Info("test overflow message",
			"fooKey", "fooValue",
			"keyForLongValue", `
I'd just like to interject for a moment. What you're refering to as Linux, is in fact, GNU/Linux, 
or as I've recently taken to calling it, GNU plus Linux. Linux is not an operating system unto itself, 
but rather another free component of a fully functioning GNU system made useful by the GNU corelibs, 
shell utilities and vital system components comprising a full OS as defined by POSIX.

Many computer users run a modified version of the GNU system every day, 
without realizing it. Through a peculiar turn of events, the version of GNU which is widely used today 
is often called Linux, and many of its users are not aware that it is basically the GNU system, 
developed by the GNU Project.

There really is a Linux, and these people are using it, but it is just a part of the system they use. 
Linux is the kernel: the program in the system that allocates the machine's resources to 
the other programs that you run. The kernel is an essential part of an operating system, but useless by itself; 
it can only function in the context of a complete operating system. Linux is normally used in combination with 
the GNU operating system: the whole system is basically GNU with Linux added, or GNU/Linux. 
All the so-called Linux distributions are really distributions of GNU/Linux!
`,
			"barKey", "barValue",
		)
	})
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
	t.Run("terminal size", func(t *testing.T) {
		term := os.Stderr
		if isatty.IsTerminal(term.Fd()) {
			w, h := termInfo()
			require.Greater(t, w, 0)
			require.Greater(t, h, 0)
			t.Logf("terminal size\nw: %d\nh: %d", w, h)
		} else {
			t.Logf("skipping test, not a terminal")
		}
	})
}
