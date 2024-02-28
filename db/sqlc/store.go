package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

// !Store defines all functions to execute db queries and transactions
type Store interface {
	Querier
	TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error)
}

// !SqlStore provides all functions execute SQL queries and transaction
type SQLStore struct {
	connPool *pgxpool.Pool
	*Queries
}

// !NewStore creates a new store
func NewStore(connPool *pgxpool.Pool) Store {
	return &SQLStore{
		connPool: connPool,
		Queries:  New(connPool),
	}
}

// The NewStore function creates a new instance of SQLStore and returns it as a value of the Store interface type. By returning the SQLStore instance as a Store, the function ensures that the returned object adheres to the methods specified in the Store interface. This allows the caller of NewStore to work with the Store interface without being aware of the specific implementation details.
