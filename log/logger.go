package log

import (
	"io"
	stdlog "log"
	"os"

	"bytes"

	"time"

	"gopkg.in/natefinch/lumberjack.v2"
)

var commonLogger *Logger

// Logger is an an io.WriteCloser rolling logger with compression and backup management.
//
// This logger is a light wrapper around:
//    "gopkg.in/natefinch/lumberjack.v2"
type Logger struct {
	delegate io.WriteCloser
}

type rotator interface {
	Rotate() error
}

// LoggerCfg is logger configuration container.
type LoggerCfg struct {
	// Filename is the file to write logs to.
	Filename string

	// MaxSize is the maximum size in megabytes of the log file before it gets
	// rotated. It defaults to 100 megabytes.
	MaxSize int

	// MaxAge is the maximum number of days to retain old log files. The default is not to remove old log files
	// based on age.
	MaxAge int

	// MaxBackups is the maximum number of old log files to retain.
	MaxBackups int

	// Compress determines if the rotated log files should be compressed
	// using gzip.
	Compress bool

	// Console determines if current log should be written to console too.
	LogToConsole bool

	// RotateOnStartup determines if logs should be rotated on application startup.
	RotateOnStartup bool
}

// ConsoleLogger is io.WriteCloser wrapper that writes to the standard console and wrapped logger.
type ConsoleLogger struct {
	delegate io.WriteCloser
}

// NewLogger creates new logger using specified configuration.
func NewLogger(cfg *LoggerCfg) (*Logger, error) {
	ljLogger := &lumberjack.Logger{
		Filename:   cfg.Filename,
		MaxSize:    cfg.MaxSize,
		MaxAge:     cfg.MaxAge,
		MaxBackups: cfg.MaxBackups,
		Compress:   cfg.Compress,
	}

	if cfg.RotateOnStartup {
		if err := ljLogger.Rotate(); err != nil {
			return nil, err
		}
	}

	if cfg.LogToConsole {
		return &Logger{&ConsoleLogger{ljLogger}}, nil
	}

	return &Logger{ljLogger}, nil
}

// InstallLogger creates new logger and installs its instance as common application logger.
func InstallLogger(cfg *LoggerCfg) error {
	logger, err := NewLogger(cfg)

	if err != nil {
		return err
	}

	stdlog.SetOutput(logger)
	stdlog.SetFlags(0)

	commonLogger = logger

	return nil
}

// Close closed common application logger if any.
func Close() error {
	if commonLogger != nil {
		if err := commonLogger.Close(); err != nil {
			return err
		}
	}

	return nil
}

// Write writes to logger.
func (l *Logger) Write(p []byte) (n int, err error) {
	buf := bytes.NewBufferString(time.Now().UTC().Format("2006-01-02T15:04:05.000Z"))

	if err := buf.WriteByte(0x20); err != nil {
		return 0, err
	}

	if n, err := buf.Write(p); err != nil {
		return n, err
	}

	return l.delegate.Write(buf.Bytes())
}

// Close closes logger.
func (l *Logger) Close() error {
	return l.delegate.Close()
}

// Rotate rotates current log.
func (l *Logger) Rotate() error {
	if rot, ok := l.delegate.(rotator); ok {
		return rot.Rotate()
	}

	return nil
}

// Write writes to console and to logger.
func (cl *ConsoleLogger) Write(p []byte) (n int, err error) {
	if _, err := os.Stdout.Write(p); err != nil {
		return 0, err
	}

	return cl.delegate.Write(p)
}

// Close closes logger.
func (cl *ConsoleLogger) Close() error {
	return cl.delegate.Close()
}

// Rotate rotates logger.
func (cl *ConsoleLogger) Rotate() error {
	if rot, ok := cl.delegate.(rotator); ok {
		return rot.Rotate()
	}

	return nil
}
