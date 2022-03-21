package logger

import (
	"context"
	"io"
	"math"
	"os"

	"github.com/sirupsen/logrus"
)

type factory struct {
	level  Level
	logger *logrus.Logger
}

type opt func(logger *logrus.Logger)

// SetOut output log writer
func SetOut(out io.Writer) opt {
	return func(logger *logrus.Logger) {
		logger.Out = out
	}
}

// NewLoggerProvider returns logger provider
func NewLoggerFactory(level Level, opts ...opt) *factory {
	engine := &logrus.Logger{
		// Logrus doesn't have public methods which ignores the level.
		// Turn the logging for all levels, verbosity checks on our side now.
		Level: math.MaxUint32,

		Hooks:     make(logrus.LevelHooks),
		Formatter: new(JsonFormatter),
	}

	for _, opt := range opts {
		opt(engine)
	}

	if engine.Out == nil {
		engine.Out = os.Stderr
	}

	return &factory{level: level, logger: engine}
}

// Logger provides an Entry
// Received Context to populate userId and traceId
// Or received Options to populate it manually via WithUserId and WithTraceId
func (l *factory) Logger(ctx context.Context) Entry {
	meta := metaFromContext(ctx)
	return newEntry(l.level, meta, l.logger)
}
