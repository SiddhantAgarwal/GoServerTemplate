# GoServerTemplate

[![codebeat badge](https://codebeat.co/badges/9ecb25b4-c1d6-4ce3-a84a-7f5fc1d363a0)](https://codebeat.co/projects/github-com-siddhantagarwal-goservertemplate-master)

*Production-ready Go server template with layered architecture, dependency injection, and PostgreSQL support.*

## Architecture

```
cmd/server         → Application entrypoint (FX bootstrap)
internal/config    → Environment-based configuration
internal/db        → PostgreSQL connection (sqlx + pgx)
internal/handler   → HTTP handlers (handler → service → repo)
internal/service   → Business logic
internal/repository → Data access layer
internal/router    → gorilla/mux routing + middleware
pkg/logger         → Structured logging (slog)
```

## Dependencies

- [gorilla/mux](https://github.com/gorilla/mux) — HTTP router
- [uber-go/fx](https://github.com/uber-go/fx) — Dependency injection & lifecycle
- [jmoiron/sqlx](https://github.com/jmoiron/sqlx) + [pgx](https://github.com/jackc/pgx) — PostgreSQL

## Environment Variables

| Variable   | Default | Description                |
|------------|---------|----------------------------|
| `PORT`     | `:8080` | Server listen address      |
| `DB_DSN`   | `""`    | Postgres DSN (optional)    |
| `LOG_LEVEL`| `info`  | slog level: debug/info/warn/error |

## Running Locally

```bash
# Run server
go run ./cmd/server

# Run tests
go test ./...

# Build binary
go build -o server ./cmd/server
./server
```

## Docker

```bash
docker build -t go-server-template .
docker run -p 8080:8080 -e PORT=:8080 go-server-template
```

## Routes

| Method | Path     | Description           |
|--------|----------|-----------------------|
| GET    | `/`      | Welcome / status      |
| GET    | `/health`| Health check (DB ping) |
