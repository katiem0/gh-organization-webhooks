package log

import (
	"testing"

	"go.uber.org/zap/zapcore"
)

func TestNewLogger(t *testing.T) {
	tests := []struct {
		name      string
		debug     bool
		wantLevel zapcore.Level
	}{
		{
			name:      "Debug enabled",
			debug:     true,
			wantLevel: zapcore.DebugLevel,
		},
		{
			name:      "Debug disabled",
			debug:     false,
			wantLevel: zapcore.InfoLevel,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger, err := NewLogger(tt.debug)

			if err != nil {
				t.Errorf("NewLogger() error = %v", err)
				return
			}

			if logger == nil {
				t.Error("Expected logger, got nil")
				return
			}

			// Check if the logger has the expected level
			// Note: This is a bit of an implementation detail, but it's a reasonable check
			if got := logger.Core().Enabled(tt.wantLevel); !got {
				t.Errorf("Logger.Core().Enabled(%v) = %v, want %v", tt.wantLevel, got, true)
			}
		})
	}
}
