package db

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/SiddhantAgarwal/GoServerTemplate/internal/config"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module("db",
		fx.Provide(NewDB),
	)
}

func NewDB(lc fx.Lifecycle, cfg *config.Config, log *slog.Logger) (*sqlx.DB, error) {
	if cfg.DBDSN == "" {
		log.Info("DB_DSN not set, skipping database connection")
		return nil, nil
	}

	db, err := sqlx.Connect("pgx", cfg.DBDSN)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Info("database connected")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Info("closing database connection")
			return db.Close()
		},
	})

	return db, nil
}
