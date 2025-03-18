package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

func (db *Database) QueryContext(ctx context.Context, query string) (*sql.Rows, error) {
	err := db.conn.PingContext(ctx)
	if err != nil {
		return nil, errors.New("failed to ping database")
	}

	rows, err := db.conn.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %v", err)
	}

	return rows, nil
}
