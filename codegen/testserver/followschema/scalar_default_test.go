package followschema

import (
	"context"
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/stretchr/testify/require"
)

func TestDefaultScalarImplementation(t *testing.T) {
	resolvers := &Stub{}

	srv := handler.New(NewExecutableSchema(Config{Resolvers: resolvers}))
	srv.AddTransport(transport.POST{})
	c := client.New(srv)

	resolvers.QueryResolver.DefaultScalar = func(ctx context.Context, arg string) (i string, e error) {
		return arg, nil
	}

	t.Run("with arg value", func(t *testing.T) {
		var resp struct{ DefaultScalar string }
		c.MustPost(`query { defaultScalar(arg: "fff") }`, &resp)
		require.Equal(t, "fff", resp.DefaultScalar)
	})

	t.Run("with default value", func(t *testing.T) {
		var resp struct{ DefaultScalar string }
		c.MustPost(`query { defaultScalar  }`, &resp)
		require.Equal(t, "default", resp.DefaultScalar)
	})
}
