package followschema

import (
	"context"
	"testing"

	"github.com/john-markham/gqlgen/client"
	"github.com/john-markham/gqlgen/graphql/handler"
	"github.com/john-markham/gqlgen/graphql/handler/transport"
	"github.com/stretchr/testify/require"
)

func TestSlices(t *testing.T) {
	resolvers := &Stub{}

	srv := handler.New(NewExecutableSchema(Config{Resolvers: resolvers}))
	srv.AddTransport(transport.POST{})
	c := client.New(srv)

	t.Run("nulls vs empty slices", func(t *testing.T) {
		resolvers.QueryResolver.Slices = func(ctx context.Context) (slices *Slices, e error) {
			return &Slices{}, nil
		}

		var resp struct {
			Slices Slices
		}
		c.MustPost(`query { slices { test1, test2, test3, test4 }}`, &resp)
		require.Nil(t, resp.Slices.Test1)
		require.Nil(t, resp.Slices.Test2)
		require.NotNil(t, resp.Slices.Test3)
		require.NotNil(t, resp.Slices.Test4)
	})

	t.Run("custom scalars to slices work", func(t *testing.T) {
		resolvers.QueryResolver.ScalarSlice = func(ctx context.Context) ([]byte, error) {
			return []byte("testing"), nil
		}

		var resp struct {
			ScalarSlice string
		}
		c.MustPost(`query { scalarSlice }`, &resp)
		require.Equal(t, "testing", resp.ScalarSlice)
	})
}
