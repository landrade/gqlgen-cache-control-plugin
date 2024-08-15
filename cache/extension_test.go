package cache

import (
	"context"
	"testing"

	"github.com/99designs/gqlgen/graphql"

	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {
	t.Parallel()

	t.Run("Validate", func(t *testing.T) {
		ext := Extension{}
		require.NoError(t, ext.Validate(nil))
	})

	t.Run("InterceptResponse", func(t *testing.T) {
		query := "query{latestPost {id}}"

		setupCtx := func(t *testing.T) context.Context {
			t.Helper()

			ctxWithOperation := graphql.WithOperationContext(
				context.Background(),
				&graphql.OperationContext{
					RawQuery: query,
				})

			return graphql.WithResponseContext(
				ctxWithOperation,
				graphql.DefaultErrorPresenter,
				graphql.DefaultRecover,
			)
		}

		t.Run("Does not try to register cacheControl extension if context is not part of an ongoing operation", func(t *testing.T) {
			ext := Extension{}

			ctx := context.TODO()

			resp := ext.InterceptResponse(ctx, func(ctx context.Context) *graphql.Response {
				return &graphql.Response{}
			})

			require.Nil(t, resp.Extensions["cacheControl"])
		})

		t.Run("Injects CacheControl in context", func(t *testing.T) {
			ext := Extension{}

			ctx := setupCtx(t)

			_ = ext.InterceptResponse(ctx, func(ctx context.Context) *graphql.Response {
				cc := CacheControl(ctx)
				require.NotNil(t, cc)
				return &graphql.Response{}
			})
		})

		t.Run("Registers cacheControl extension around the graphql response operation", func(t *testing.T) {
			ext := Extension{}

			ctx := setupCtx(t)

			var registeredExtensions map[string]interface{}

			resp := ext.InterceptResponse(ctx, func(ctx context.Context) *graphql.Response {
				registeredExtensions = graphql.GetExtensions(ctx)

				return &graphql.Response{Extensions: registeredExtensions}
			})

			require.NotNil(t, resp.Extensions["cacheControl"])
		})
	})
}
