package logger

import (
	"context"
)

var logger = NewLoggerFactory(DebugLevel).
	Logger(context.Background())

// SetDefaultEntry
func SetDefaultEntry(entry Entry) {
	logger = entry
}

// WithFields returns copy of default entry with fields
func WithFields(fields Fields) Entry {
	return logger.WithFields(fields)
}

func (e *entry) Debug(args ...interface{}) {
	e.log(DebugLevel, args...)
}

func (e *entry) Info(args ...interface{}) {
	e.log(InfoLevel, args...)
}

func (e *entry) Alert(args ...interface{}) {
	e.log(AlertLevel, args...)
}

func (e *entry) Notice(args ...interface{}) {
	e.log(NoticeLevel, args...)
}

func (e *entry) Warning(args ...interface{}) {
	e.log(WarningLevel, args...)
}

func (e *entry) Error(args ...interface{}) {
	e.log(ErrorLevel, args...)
}

func (e *entry) Emergency(args ...interface{}) {
	e.log(EmergencyLevel, args...)
}

func (e *entry) Critical(args ...interface{}) {
	e.log(CriticalLevel, args...)
}

func (e *entry) Debugf(format string, args ...interface{}) {
	e.logf(DebugLevel, format, args...)
}

func (e *entry) Alertf(format string, args ...interface{}) {
	e.logf(AlertLevel, format, args...)
}

func (e *entry) Infof(format string, args ...interface{}) {
	e.logf(InfoLevel, format, args...)
}

func (e *entry) Noticef(format string, args ...interface{}) {
	e.logf(NoticeLevel, format, args...)
}

func (e *entry) Warningf(format string, args ...interface{}) {
	e.logf(WarningLevel, format, args...)
}

func (e *entry) Errorf(format string, args ...interface{}) {
	e.logf(ErrorLevel, format, args...)
}

func (e *entry) Emergencyf(format string, args ...interface{}) {
	e.logf(EmergencyLevel, format, args...)
}

func (e *entry) Criticalf(format string, args ...interface{}) {
	e.logf(CriticalLevel, format, args...)
}

func Debug(args ...interface{}) {
	logger.Debug(args...)
}

func Alert(args ...interface{}) {
	logger.Alert(args...)
}

func Info(args ...interface{}) {
	logger.Info(args...)
}

func Notice(args ...interface{}) {
	logger.Notice(args...)
}

func Warning(args ...interface{}) {
	logger.Warning(args...)
}

func Error(args ...interface{}) {
	logger.Error(args...)
}

func Emergency(args ...interface{}) {
	logger.Emergency(args...)
}

func Critical(args ...interface{}) {
	logger.Critical(args...)
}

func Debugf(format string, args ...interface{}) {
	logger.Debugf(format, args...)
}

func Alertf(format string, args ...interface{}) {
	logger.Alertf(format, args...)
}

func Infof(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

func Noticef(format string, args ...interface{}) {
	logger.Noticef(format, args...)
}

func Warningf(format string, args ...interface{}) {
	logger.Warningf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	logger.Errorf(format, args...)
}

func Emergencyf(format string, args ...interface{}) {
	logger.Emergencyf(format, args...)
}

func Criticalf(format string, args ...interface{}) {
	logger.Criticalf(format, args...)
}
