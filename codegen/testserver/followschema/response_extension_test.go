package followschema

import (
	"context"
	"testing"

	"github.com/john-markham/gqlgen/graphql/handler/transport"
	"github.com/stretchr/testify/require"

	"github.com/john-markham/gqlgen/client"
	"github.com/john-markham/gqlgen/graphql"
	"github.com/john-markham/gqlgen/graphql/handler"
)

func TestResponseExtension(t *testing.T) {
	resolvers := &Stub{}
	resolvers.QueryResolver.Valid = func(ctx context.Context) (s string, e error) {
		return "Ok", nil
	}

	srv := handler.New(NewExecutableSchema(Config{Resolvers: resolvers}))
	srv.AddTransport(transport.POST{})
	srv.AroundResponses(func(ctx context.Context, next graphql.ResponseHandler) *graphql.Response {
		graphql.RegisterExtension(ctx, "example", "value")

		return next(ctx)
	})

	c := client.New(srv)

	raw, _ := c.RawPost(`query { valid }`)
	require.Equal(t, "value", raw.Extensions["example"])
}
