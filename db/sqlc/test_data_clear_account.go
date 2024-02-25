package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/gommon/log"
)

// !TruncateTestAccountData
func TruncateTestAccountData(ctx context.Context, connPool *pgxpool.Pool) {
	_, truncateResultErr := connPool.Exec(ctx, "TRUNCATE accounts RESTART IDENTITY")
	if truncateResultErr != nil {
		log.Error("Not account exists on database")
	} else {
		log.Info("Accounts table truncated")
	}
}
