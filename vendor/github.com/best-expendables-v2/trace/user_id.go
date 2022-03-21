package trace

import "context"

const (
	userCtxKey ctxKey = iota
)

func GetCurrentUserFromContext(ctx context.Context) *User {
	if user, ok := ctx.Value(userCtxKey).(*User); ok {
		return user
	}
	return nil
}

func ContextWithUser(ctx context.Context, user *User) context.Context {
	return context.WithValue(ctx, userCtxKey, user)
}

func GetUserIDFromContext(ctx context.Context) string {
	user := GetCurrentUserFromContext(ctx)
	var userID string
	if user != nil {
		userID = user.Id
	}
	return userID
}