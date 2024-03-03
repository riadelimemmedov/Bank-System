package main

import (
	"context"
	"log"
	"simplebank/api"
	db "simplebank/db/sqlc"
	"simplebank/util"

	"github.com/jackc/pgx/v5/pgxpool"
)

// !main
func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config file: ", err)
	}
	connPool, err := pgxpool.New(context.Background(), config.DbSource)
	if err != nil {
		log.Fatal("Cannot connect to db ", err)
	}

	store := db.NewStore(connPool)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)

	if err != nil {
		log.Fatal("cannot start server", err)
	}

}
