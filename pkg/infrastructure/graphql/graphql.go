package graphql

import (
	"webreader/ent"
	"webreader/pkg/adapter/controller"
	"webreader/pkg/adapter/resolver"

	"entgo.io/contrib/entgql"

	"github.com/99designs/gqlgen/graphql/handler"
)

// NewServer generates graphql server
func NewServer(client *ent.Client, controller controller.Controller) *handler.Server {
	srv := handler.NewDefaultServer(resolver.NewSchema(client, controller))
	srv.Use(entgql.Transactioner{TxOpener: client})

	return srv
}
