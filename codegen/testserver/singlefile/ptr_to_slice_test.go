package singlefile

import (
	"context"
	"testing"

	"github.com/john-markham/gqlgen/client"
	"github.com/john-markham/gqlgen/graphql/handler"
	"github.com/john-markham/gqlgen/graphql/handler/transport"
	"github.com/stretchr/testify/require"
)

func TestPtrToSlice(t *testing.T) {
	resolvers := &Stub{}

	srv := handler.New(NewExecutableSchema(Config{Resolvers: resolvers}))
	srv.AddTransport(transport.POST{})
	c := client.New(srv)

	resolvers.QueryResolver.PtrToSliceContainer = func(ctx context.Context) (wrappedStruct *PtrToSliceContainer, e error) {
		ptrToSliceContainer := PtrToSliceContainer{
			PtrToSlice: &[]string{"hello"},
		}
		return &ptrToSliceContainer, nil
	}

	t.Run("pointer to slice", func(t *testing.T) {
		var resp struct {
			PtrToSliceContainer struct {
				PtrToSlice []string
			}
		}

		err := c.Post(`query { ptrToSliceContainer {  ptrToSlice }}`, &resp)
		require.NoError(t, err)

		require.Equal(t, []string{"hello"}, resp.PtrToSliceContainer.PtrToSlice)
	})
}
