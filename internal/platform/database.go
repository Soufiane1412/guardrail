package platform

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

// NewDatabasePool creates a connection pool that the sentinel will use to talk to Postgres
func NewDatabasePool(ctx context.Context) (*pgxpool.pool, error) {
	// 1. Fetch the Connection String from the environment.
	// IMPACT: Security. Always hide DB password
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		return nil, fmt.Errorf("DATABASE_URL environment variable is not set")
	}

	// 2. Create the Pool
	// IMPACT: Efficiency. this sets up the 'pool' of operators waiting for calls.
	pool, err := pgxpool.New(ctx, connStr)
	if err != nil {
		return nil, fmt.Errorf("unable to create connection pool: %w", err)
	}

	// 3. The 'Ping'.
	// IMPACT: reliability we verify the heart is beating before we move on.
	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("unabunable to ping database: %w", err)
	}

	return pool, nil
}
