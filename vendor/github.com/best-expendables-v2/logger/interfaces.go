package logger

import (
	"context"
)

type (
	Factory interface {
		Logger(ctx context.Context) Entry
	}

	Entry interface {
		Info(args ...interface{})
		Debug(args ...interface{})
		Notice(args ...interface{})
		Warning(args ...interface{})
		Alert(args ...interface{})
		Error(args ...interface{})
		Emergency(args ...interface{})
		Critical(args ...interface{})

		Infof(format string, args ...interface{})
		Debugf(format string, args ...interface{})
		Noticef(format string, args ...interface{})
		Warningf(format string, args ...interface{})
		Alertf(format string, args ...interface{})
		Errorf(format string, args ...interface{})
		Emergencyf(format string, args ...interface{})
		Criticalf(format string, args ...interface{})

		WithField(key string, value interface{}) Entry
		WithFields(fields Fields) Entry
		GetFields() Fields
	}

	Fields map[string]interface{}
)
