# Cache control plugin for GQLGEN

Plugin to write cache-control directives. With this plugin, you can write a cache-control HTTP header and cache-control extension, following Apollo's rules. Refs:

- https://www.apollographql.com/docs/apollo-server/performance/caching/
- https://github.com/apollographql/apollo-server/tree/d5015f4ea00cadb2a74b09956344e6f65c084629/packages/apollo-cache-control


### Installation

Install Go module:

```bash
$ go get github.com/landrade/gqlgen-cache-control-plugin
```

### Configuring the plugin

```go
package main
import (
	"github.com/landrade/gqlgen-cache-control-plugin/cache"
	"github.com/99designs/gqlgen/graphql/handler"
)
func main() {
	// Building your server
	cfg := generated.Config{Resolvers: &graph.Resolver{}}
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(cfg))
	// Enable cache extensions
	srv.Use(cache.Extension{}) // <----
	//...
}
```

### Set cache hints

After you enable `cache.Extension`, you can set cache hints using `cache.SetHint` function in your resolvers.

```go
import (
	"github.com/landrade/gqlgen-cache-control-plugin/cache"
	// ...
)
func (r *commentResolver) Post(ctx context.Context, obj *model.Comment) (*model.Post, error) {
    post, err := // getting post by comment
    if err != nil {
        return nil, err
	}
	// Set a CacheHint
	cache.SetHint(ctx, cache.ScopePublic, 10*time.Second)
	return post, nil
}
```

### CDN Caching

It's possible to enable the Gqlgen to provide a `Cache-Control` header based on your cache hints in `GET` or `POST` requests.
To do it you need wrap your server using `cache.Middleware` function. It will add `Cache-Control` header to all responses.


```go
func main() {
	// ... setup server
	srv = cache.Middleware(srv)
	// ... do more things
}
````

Doing it, Gqlgen write the lowest max-age defined in cacheControl extensions.

For more informations, see `_example` folder.
