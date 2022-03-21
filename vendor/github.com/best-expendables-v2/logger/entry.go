package logger

import (
	"github.com/sirupsen/logrus"
)

type (
	// entry extends logrus
	// Implementation of our logger contracts you can find logger.go file.
	// In entry.go should be implementation for copying, creation and etc, I want to keep this file clean.
	entry struct {
		level Level

		// Meta contains top field properties
		// Metadata doesn't use shallow copy as content, therefore should be immutable after creation.
		// e.g.: user-id, trace-id and can not be edited manually.
		meta meta

		// content contains everything what come to the logger
		content []field

		logger *logrus.Logger
	}

	field struct {
		Key   string
		Value interface{}
	}
)

func newEntry(level Level, meta meta, logger *logrus.Logger) *entry {
	return &entry{
		level:   level,
		meta:    meta,
		content: []field{},
		logger:  logger,
	}
}

func (e *entry) WithField(key string, value interface{}) Entry {
	return e.WithFields(Fields{key: value})
}

func (e *entry) GetFields() Fields {
	fields := make(map[string]interface{})
	for _, c := range e.content {
		fields[c.Key] = c.Value
	}
	return fields
}

func (e *entry) WithFields(fields Fields) Entry {
	e2 := *e

	for k, v := range fields {
		e2.content = append(e2.content, field{k, v})
	}

	return &e2
}

func (e *entry) entry(level Level) *logrus.Entry {
	// Convert field back to Fields (map)
	// so it can be serialized into
	// "Test key: Test value"
	// instead of "{ "Key": "TestKey", "Value": "TestValue" }"
	contentMap := make(map[string]interface{})
	for _, content := range e.content {
		contentMap[content.Key] = content.Value
	}

	entry := logrus.Entry{
		Data: logrus.Fields{
			metaTmpField: e.meta,
			FieldLevel:   level,
			FieldContent: contentMap,
		},
		Logger: e.logger,
	}

	return &entry
}

func (e *entry) log(level Level, args ...interface{}) {
	if e.level >= level {
		e.entry(level).Print(args...)
	}
}

func (e *entry) logf(level Level, format string, args ...interface{}) {
	if e.level >= level {
		e.entry(level).Printf(format, args...)
	}
}
