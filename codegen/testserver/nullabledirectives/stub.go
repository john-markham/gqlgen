// Code generated by github.com/john-markham/gqlgen, DO NOT EDIT.

package followschema

import (
	"context"

	"github.com/john-markham/gqlgen/codegen/testserver/nullabledirectives/generated"
)

type Stub struct {
	QueryResolver struct {
		DirectiveSingleNullableArg func(ctx context.Context, arg1 *string) (*string, error)
	}
}

func (r *Stub) Query() generated.QueryResolver {
	return &stubQuery{r}
}

type stubQuery struct{ *Stub }

func (r *stubQuery) DirectiveSingleNullableArg(ctx context.Context, arg1 *string) (*string, error) {
	return r.QueryResolver.DirectiveSingleNullableArg(ctx, arg1)
}
