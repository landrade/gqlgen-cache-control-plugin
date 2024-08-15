package cache

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

type Extension struct{}

var _ interface {
	graphql.HandlerExtension
	graphql.ResponseInterceptor
} = Extension{}

func (c Extension) ExtensionName() string {
	return "cache"
}

func (c Extension) Validate(_ graphql.ExecutableSchema) error {
	return nil
}

func (c Extension) InterceptResponse(ctx context.Context, next graphql.ResponseHandler) *graphql.Response {
	if !graphql.HasOperationContext(ctx) {
		return next(ctx)
	}

	cache := CacheControl(ctx)
	if cache == nil {
		ctx = WithCacheControlExtension(ctx)

		cache = CacheControl(ctx)
	}

	graphql.RegisterExtension(ctx, "cacheControl", cache)

	return next(ctx)
}
