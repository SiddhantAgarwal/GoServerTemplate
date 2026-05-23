package repository

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type HealthRepository interface {
	Ping(ctx context.Context) error
}

type healthRepo struct {
	db *sqlx.DB
}

func NewHealthRepository(db *sqlx.DB) HealthRepository {
	return &healthRepo{db: db}
}

func (r *healthRepo) Ping(ctx context.Context) error {
	if r.db == nil {
		return sql.ErrConnDone
	}

	return r.db.PingContext(ctx)
}
