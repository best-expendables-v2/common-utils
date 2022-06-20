package filter

import (
	"context"
	"net/http"
)

var isUnscopedKey = "is_unscoped"
var isPreloadUnscopedKey = "is_preload_unscoped"

func SetUnscoped(ctx context.Context) context.Context {
	return context.WithValue(ctx, isUnscopedKey, true)
}

func SetPreloadUnscoped(ctx context.Context) context.Context {
	return context.WithValue(ctx, isPreloadUnscopedKey, true)
}

func GetUnscoped(ctx context.Context) bool {
	if v := ctx.Value(isUnscopedKey); v != nil {
		return v.(bool)
	}
	return false
}

func GetPreloadUnscoped(ctx context.Context) bool {
	if v := ctx.Value(isPreloadUnscopedKey); v != nil {
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
		if values.Get("preload_unscoped") == "true" {
			r = r.WithContext(SetPreloadUnscoped(r.Context()))
		}
		next.ServeHTTP(w, r)
	})
}
