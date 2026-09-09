package db

import (
    "context"
    "embed"
    "io/fs"
    "path/filepath"
)

//go:embed ../../../../db/migrations/*.sql
var migrationsFS embed.FS

func RunMigrations(ctx context.Context, pool interface{}) error {
    // Basitleştirilmiş — üretimde golang-migrate kullan
    return nil
}
