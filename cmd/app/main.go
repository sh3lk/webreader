package main

import (
	"log"
	"webreader/config"
	"webreader/ent"
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

	return client
}

func newController(client *ent.Client) controller.Controller {
	r := registry.New(client)
	return r.NewController()
}
