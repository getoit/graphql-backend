package xidgql

import (
	"fmt"
	"io"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
	"github.com/rs/xid"
)

func MarshalID(id xid.ID) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		_, _ = io.WriteString(w, strconv.Quote(id.String()))
	})
}

func UnmarshalID(v interface{}) (id xid.ID, err error) {
	s, ok := v.(string)
	if !ok {
		return id, fmt.Errorf("invalid type %T, expect string", v)
	}
	return xid.FromString(s)
}
