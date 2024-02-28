package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://postgres:123321@127.0.0.1:6432/simple_bank?sslmode=disable"
)

var testStore Store

// !TestMain
func TestMain(m *testing.M) {
	connPool, err := pgxpool.New(context.Background(), dbSource)
	if err != nil {
		log.Fatal("Cannot connect to db ", err)
	}
	testStore = NewStore(connPool)
	os.Exit(m.Run())
}
