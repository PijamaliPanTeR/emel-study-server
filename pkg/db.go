package pkg

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
)

// NewDB opens a PostgreSQL connection using the same pattern as main-service (database/sql + pgx driver).
func NewDB(cfg DatabaseConnConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName)
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("open db: %w", err)
	}
	if err := db.Ping(); err != nil {
		_ = db.Close()
		return nil, fmt.Errorf("ping db: %w", err)
	}
	return db, nil
}
