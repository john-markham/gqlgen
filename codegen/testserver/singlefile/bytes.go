package singlefile

import (
	"fmt"
	"io"

	"github.com/john-markham/gqlgen/graphql"
)

func MarshalBytes(b []byte) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		_, _ = fmt.Fprintf(w, "%q", string(b))
	})
}

func UnmarshalBytes(v any) ([]byte, error) {
	switch v := v.(type) {
	case string:
		return []byte(v), nil
	case *string:
		return []byte(*v), nil
	case []byte:
		return v, nil
	default:
		return nil, fmt.Errorf("%T is not []byte", v)
	}
}
