package logger

import "context"

type ctxKey int

const entryCtxKey ctxKey = iota

// ContextWithEntry set entry to context
func ContextWithEntry(e Entry, ctx context.Context) context.Context {
	return context.WithValue(ctx, entryCtxKey, e)
}

// EntryFromContext get entry from context
func EntryFromContext(ctx context.Context) Entry {
	if logger, ok := ctx.Value(entryCtxKey).(Entry); ok {
		return logger
	}

	return nil
}

// EntryFromContextOrDefault returns the logger from context if not nil, otherwise, returns default entry
func EntryFromContextOrDefault(ctx context.Context) Entry {
	if entry := EntryFromContext(ctx); entry != nil {
		return entry
	}

	return logger
}
