package logger

import (
	"encoding/json"
	"fmt"
	"time"

	userclient "github.com/best-expendables-v2/user-service-client"

	"context"

	"github.com/best-expendables-v2/trace"
	"github.com/sirupsen/logrus"
)

const (
	// FormatVersion log format version
	FormatVersion uint8 = 2
	// MessageLength max length of message
	MessageLength = 128
)

const (
	FieldVersion   = "@version"
	FieldTimestamp = "@timestamp"
	FieldLevel     = "level"
	FieldContent   = "content"
	FieldMessage   = "message"
	FieldUserId    = "userId"
	FieldTraceId   = "traceId"
	FieldRequestId = "context-id"

	metaTmpField = "meta"
)

type (
	JsonFormatter struct{}

	// message serializable result message
	message struct {
		meta

		Version   uint8       `json:"@version"`
		Timestamp string      `json:"@timestamp"`
		Message   string      `json:"message"`
		Level     string      `json:"level"`
		Content   interface{} `json:"content"`
	}

	meta struct {
		TraceID   string `json:"traceId"`
		UserID    string `json:"userId"`
		RequestID string `json:"context-id"`
	}
)

func (JsonFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	msg := message{
		Version:   FormatVersion,
		Timestamp: entry.Time.Format(time.RFC3339Nano),
		Level:     entry.Data["level"].(Level).String(),
		Content:   entry.Data[FieldContent],
	}

	if meta, ok := entry.Data[metaTmpField].(meta); ok {
		msg.meta = meta
	}

	if len(entry.Message) > MessageLength {
		msg.Message = entry.Message[:MessageLength]
	} else {
		msg.Message = entry.Message
	}

	serialized, err := json.Marshal(msg)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal fields to JSON: %s", err)
	}
	return append(serialized, '\n'), nil
}

func metaFromContext(ctx context.Context) meta {
	meta := meta{RequestID: trace.RequestIDFromContext(ctx)}

	if user := userclient.GetCurrentUserFromContext(ctx); user != nil {
		meta.UserID = user.Id
	}

	return meta
}
