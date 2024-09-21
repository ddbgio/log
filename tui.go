package log

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	"errors"

	"golang.org/x/crypto/ssh/terminal"
)

type TUI struct {
	level slog.Leveler
}

const (
	iconDebug = "🔷"
	iconInfo  = "🟢"
	iconWarn  = "🟨"
	iconError = "❌"
	iconRun   = "▶"
	// iconRun   = "💲"
)

var (
	errTuiFail       = errors.New("tui print failed")
	errInvalidFields = errors.New("expected even number of key-value pairs")
	formatTime       = "15:04:05"
	kvIndent         = 12
)

// NewTUI returns a new TUI printer
func NewTUI(opts slog.HandlerOptions) (*TUI, error) {
	t := &TUI{
		level: opts.Level,
	}
	return t, nil
}

func printTUI(msg string, icon string, fields ...interface{}) {
	headers, values, err := kvParse(fields...)
	if err != nil {
		err := fmt.Errorf("%w: %v", errTuiFail, errInvalidFields)
		fmt.Println(err)
	}
	// print just the time, include seconds
	now := time.Now().Format(formatTime)
	message := fmt.Sprintf("%s %s %s", now, icon, msg)
	if len(headers) != len(values) {
		err := fmt.Errorf("%w: %v", errTuiFail, errInvalidFields)
		fmt.Println(err)
		return
	}
	fmt.Println(message)
	for i, header := range headers {
		fmt.Printf("%*s%12s = %-40s\n",
			kvIndent, "", header, values[i],
		)
	}

}

func (t *TUI) Debug(msg string, fields ...interface{}) {
	if t.level.Level() <= slog.LevelDebug {
		printTUI(msg, iconDebug, fields...)
	}
}

func (t *TUI) Info(msg string, fields ...interface{}) {
	if t.level.Level() <= slog.LevelInfo {
		printTUI(msg, iconInfo, fields...)
	}
}

func (t *TUI) Warn(msg string, fields ...interface{}) {
	if t.level.Level() <= slog.LevelWarn {
		printTUI(msg, iconWarn, fields...)
	}
}

func (t *TUI) Error(msg string, fields ...interface{}) {
	if t.level.Level() <= slog.LevelError {
		printTUI(msg, iconError, fields...)
	}
}

// kvParse parses key-value pairs into separate slices
func kvParse(fields ...interface{}) ([]string, []string, error) {
	if len(fields)%2 != 0 {
		return nil, nil, fmt.Errorf("expected even number of key-value pairs, got %d", len(fields))
	}
	var headers []string
	var values []string
	for i, field := range fields {
		if i%2 == 0 {
			headers = append(headers, fmt.Sprintf("%v", field))
		} else {
			values = append(values, fmt.Sprintf("%v", field))
		}
	}
	if len(headers) != len(values) {
		return nil, nil, errInvalidFields
	}
	return headers, values, nil
}

// termInfo returns the terminal width and height
func termInfo() (int, int) {
	width, height, err := terminal.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return 0, 0
	}
	return width, height
}
