package main

import (
	"context"
	"log"
	"webreader/config"
	"webreader/ent"
	"webreader/ent/migrate"
	"webreader/pkg/adapter/controller"
	"webreader/pkg/infrastructure/datastore"
	"webreader/pkg/infrastructure/graphql"
	"webreader/pkg/infrastructure/router"
	"webreader/pkg/registry"
)

func main() {
	config.ReadConfig(config.ReadConfigOption{})

	client := newDBClient()
	ctrl := newController(client)

	srv := graphql.NewServer(client, ctrl)
	e := router.New(srv)

	e.Logger.Fatal(e.Start(":" + config.C.Server.Address))
}

func newDBClient() *ent.Client {
	client, err := datastore.NewClient()
	if err != nil {
		log.Fatalf("failed opening mysql client: %v", err)
	}

	if err := client.Schema.Create(
		context.Background(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
		migrate.WithForeignKeys(true),
	); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return client
}

func newController(client *ent.Client) controller.Controller {
	r := registry.New(client)
	return r.NewController()
}
