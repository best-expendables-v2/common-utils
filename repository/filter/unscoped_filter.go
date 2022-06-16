package filter

import (
	"context"
	"net/http"
)

var isUnscopedKey = "is_unscoped"

func SetUnscoped(ctx context.Context) context.Context {
	return context.WithValue(ctx, isUnscopedKey, true)
}

func GetUnscoped(ctx context.Context) bool {
	if v := ctx.Value(isUnscopedKey); v != nil {
		return v.(bool)
	}
	return false
}

func DBUnscopedFilter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		values := r.URL.Query()
		if values.Get("unscoped") == "true" {
			r = r.WithContext(SetUnscoped(r.Context()))
		}
		next.ServeHTTP(w, r)
	})
}
