package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.uber.org/fx"

	"github.com/SiddhantAgarwal/GoServerTemplate/internal/config"
	"github.com/SiddhantAgarwal/GoServerTemplate/internal/db"
	"github.com/SiddhantAgarwal/GoServerTemplate/internal/handler"
	"github.com/SiddhantAgarwal/GoServerTemplate/internal/repository"
	"github.com/SiddhantAgarwal/GoServerTemplate/internal/router"
	"github.com/SiddhantAgarwal/GoServerTemplate/internal/service"
	applogger "github.com/SiddhantAgarwal/GoServerTemplate/pkg/logger"
)

type Server struct {
	server *http.Server
	log    *slog.Logger
}

func NewServer(lc fx.Lifecycle, router *mux.Router, cfg *config.Config, log *slog.Logger) *Server {
	s := &Server{
		server: &http.Server{
			Addr:    cfg.Addr(),
			Handler: router,
		},
		log: log,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				log.Info("server starting", "addr", s.server.Addr)

				if err := s.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
					log.Error("server error", "error", err)
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Info("shutting down server")

			shutdownCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
			defer cancel()

			return s.server.Shutdown(shutdownCtx)
		},
	})

	return s
}

func main() {
	fx.New(
		fx.Provide(config.Load),
		fx.Provide(func(cfg *config.Config) *slog.Logger {
			return applogger.New(cfg.LogLevel)
		}),
		db.Module(),
		repository.Module(),
		service.Module(),
		handler.Module(),
		router.Module(),
		fx.Provide(NewServer),
		fx.Invoke(func(s *Server) {}),
	).Run()
}
