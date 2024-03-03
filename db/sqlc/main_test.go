package db

import (
	"context"
	"log"
	"os"
	"simplebank/util"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

var testStore Store

// !TestMain
func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config", err)
	}
	connPool, err := pgxpool.New(context.Background(), config.DbSource)
	if err != nil {
		log.Fatal("Cannot connect to db ", err)
	}
	testStore = NewStore(connPool)
	os.Exit(m.Run())
}
