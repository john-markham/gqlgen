package followschema

import (
	"context"
	"testing"

	"github.com/john-markham/gqlgen/graphql/handler/transport"
	"github.com/stretchr/testify/require"

	"github.com/john-markham/gqlgen/client"
	"github.com/john-markham/gqlgen/graphql/handler"
)

type fakeUnexportedEmbeddedInterface struct{}

func (*fakeUnexportedEmbeddedInterface) UnexportedEmbeddedInterfaceExportedMethod() string {
	return "UnexportedEmbeddedInterfaceExportedMethod"
}

func TestEmbedded(t *testing.T) {
	resolver := &Stub{}
	resolver.QueryResolver.EmbeddedCase1 = func(ctx context.Context) (*EmbeddedCase1, error) {
		return &EmbeddedCase1{}, nil
	}
	resolver.QueryResolver.EmbeddedCase2 = func(ctx context.Context) (*EmbeddedCase2, error) {
		return &EmbeddedCase2{&unexportedEmbeddedPointer{}}, nil
	}
	resolver.QueryResolver.EmbeddedCase3 = func(ctx context.Context) (*EmbeddedCase3, error) {
		return &EmbeddedCase3{&fakeUnexportedEmbeddedInterface{}}, nil
	}

	srv := handler.New(NewExecutableSchema(Config{Resolvers: resolver}))
	srv.AddTransport(transport.POST{})
	c := client.New(srv)

	t.Run("embedded case 1", func(t *testing.T) {
		var resp struct {
			EmbeddedCase1 struct {
				ExportedEmbeddedPointerExportedMethod string
			}
		}
		err := c.Post(`query { embeddedCase1 { exportedEmbeddedPointerExportedMethod } }`, &resp)
		require.NoError(t, err)
		require.Equal(t, "ExportedEmbeddedPointerExportedMethodResponse", resp.EmbeddedCase1.ExportedEmbeddedPointerExportedMethod)
	})

	t.Run("embedded case 2", func(t *testing.T) {
		var resp struct {
			EmbeddedCase2 struct {
				UnexportedEmbeddedPointerExportedMethod string
			}
		}
		err := c.Post(`query { embeddedCase2 { unexportedEmbeddedPointerExportedMethod } }`, &resp)
		require.NoError(t, err)
		require.Equal(t, "UnexportedEmbeddedPointerExportedMethodResponse", resp.EmbeddedCase2.UnexportedEmbeddedPointerExportedMethod)
	})

	t.Run("embedded case 3", func(t *testing.T) {
		var resp struct {
			EmbeddedCase3 struct {
				UnexportedEmbeddedInterfaceExportedMethod string
			}
		}
		err := c.Post(`query { embeddedCase3 { unexportedEmbeddedInterfaceExportedMethod } }`, &resp)
		require.NoError(t, err)
		require.Equal(t, "UnexportedEmbeddedInterfaceExportedMethod", resp.EmbeddedCase3.UnexportedEmbeddedInterfaceExportedMethod)
	})
}
