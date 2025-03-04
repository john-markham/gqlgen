package singlefile

import (
	"context"
	"testing"

	"github.com/john-markham/gqlgen/client"
	"github.com/john-markham/gqlgen/graphql/handler"
	"github.com/john-markham/gqlgen/graphql/handler/transport"
	"github.com/stretchr/testify/require"
)

type FieldsOrderPayloadResults struct {
	OverrideValueViaInput struct {
		FirstFieldValue *string `json:"firstFieldValue"`
	} `json:"overrideValueViaInput"`
}

func TestFieldsOrder(t *testing.T) {
	resolvers := &Stub{}

	srv := handler.New(NewExecutableSchema(Config{Resolvers: resolvers}))
	srv.AddTransport(transport.POST{})
	c := client.New(srv)
	resolvers.FieldsOrderInputResolver.OverrideFirstField = func(ctx context.Context, in *FieldsOrderInput, data *string) error {
		if data != nil {
			in.FirstField = data
		}
		return nil
	}
	resolvers.MutationResolver.OverrideValueViaInput = func(ctx context.Context, in FieldsOrderInput) (ret *FieldsOrderPayload, err error) {
		ret = &FieldsOrderPayload{
			FirstFieldValue: in.FirstField,
		}
		return
	}

	t.Run("firstField", func(t *testing.T) {
		var resp FieldsOrderPayloadResults

		err := c.Post(`mutation {
			overrideValueViaInput(input: { firstField:"newName" }) {
				firstFieldValue
			}
		}`, &resp)
		require.NoError(t, err)

		require.NotNil(t, resp.OverrideValueViaInput.FirstFieldValue)
		require.Equal(t, "newName", *resp.OverrideValueViaInput.FirstFieldValue)
	})

	t.Run("firstField/override", func(t *testing.T) {
		var resp FieldsOrderPayloadResults

		err := c.Post(`mutation { overrideValueViaInput(input: {
				firstField:"newName",
				overrideFirstField: "override"
			}) {
				firstFieldValue
			}
		}`, &resp)
		require.NoError(t, err)

		require.NotNil(t, resp.OverrideValueViaInput.FirstFieldValue)
		require.NotEqual(t, "newName", *resp.OverrideValueViaInput.FirstFieldValue)
		require.Equal(t, "override", *resp.OverrideValueViaInput.FirstFieldValue)
	})
}
