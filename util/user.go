package util

import (
	"context"
	userclient "github.com/best-expendables-v2/user-service-client"
)

func GetUserIDFromContext(ctx context.Context) string {
	user := userclient.GetCurrentUserFromContext(ctx)
	var userID string
	if user != nil {
		userID = user.Id
	}
	return userID
}

func GetUserFromContext(ctx context.Context) *userclient.User {
	return userclient.GetCurrentUserFromContext(ctx)
}
