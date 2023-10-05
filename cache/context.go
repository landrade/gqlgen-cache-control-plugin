package cache

import "context"

type contextKey string

const forceCacheControlKey = contextKey("forceCacheControl")

// ContextWithForceCacheControl creates a new context with the forceCacheControl option.
func ContextWithForceCacheControl(ctx context.Context, forceCacheControl bool) context.Context {
	return context.WithValue(ctx, forceCacheControlKey, forceCacheControl)
}
