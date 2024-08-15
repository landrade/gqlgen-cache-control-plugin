package cache

import (
	"context"
	"sync"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/ast"
)

type (
	Scope string

	ctxKey string

	OverallCachePolicy struct {
		MaxAge float64
		Scope  Scope
	}

	CacheControlExtension struct {
		Version int    `json:"version"`
		Hints   []Hint `json:"hints"`
		mu      sync.Mutex
	}

	Hint struct {
		Path   ast.Path `json:"path"`
		MaxAge float64  `json:"maxAge"`
		Scope  Scope    `json:"scope"`
	}
)

const (
	ScopePublic  = Scope("PUBLIC")
	ScopePrivate = Scope("PRIVATE")

	cacheCtxKey ctxKey = "key"
)

func (cache *CacheControlExtension) AddHint(h Hint) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	cache.Hints = append(cache.Hints, h)
}

// OverallPolicy return a calculated cache policy
func (cache *CacheControlExtension) OverallPolicy() OverallCachePolicy {
	var (
		scope     = ScopePublic
		maxAge    float64
		hasMaxAge bool
	)

	for _, c := range cache.Hints {

		if c.Scope == ScopePrivate {
			scope = c.Scope
		}

		if !hasMaxAge || c.MaxAge < maxAge {
			hasMaxAge = true
			maxAge = c.MaxAge
		}
	}

	return OverallCachePolicy{
		MaxAge: maxAge,
		Scope:  scope,
	}
}

func WithCacheControlExtension(ctx context.Context) context.Context {
	cache := &CacheControlExtension{Version: 1}

	return context.WithValue(ctx, cacheCtxKey, cache)
}

func CacheControl(ctx context.Context) *CacheControlExtension {
	c := ctx.Value(cacheCtxKey)
	if c, ok := c.(*CacheControlExtension); ok {
		return c
	}

	return nil
}

func SetHint(ctx context.Context, scope Scope, maxAge time.Duration) {
	c := ctx.Value(cacheCtxKey)
	if c, ok := c.(*CacheControlExtension); ok {
		c.AddHint(Hint{
			Path:   graphql.GetFieldContext(ctx).Path(),
			MaxAge: maxAge.Seconds(),
			Scope:  scope,
		})
	}
}

// GetOverallCachePolicy is responsible to extract cache policy from a Response.
// If does not have any cacheControl in Extensions, it will return (empty, false)
func GetOverallCachePolicy(cache *CacheControlExtension) (OverallCachePolicy, bool) {
	overallPolicy := cache.OverallPolicy()
	if overallPolicy.MaxAge > 0 {
		return overallPolicy, true
	}

	return OverallCachePolicy{}, false
}
