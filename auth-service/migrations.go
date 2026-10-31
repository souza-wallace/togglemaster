package main

import (
	"context"
	"database/sql"
	_ "embed"
	"log"
)

// O init.sql vai embutido no binário: a imagem não precisa copiar a pasta db/.
//
//go:embed db/init.sql
var schemaSQL string

// Identificador arbitrário do lock. Só precisa ser o mesmo entre as réplicas do auth-service.
const migrationLockID = 820001

// runMigrations aplica o schema antes de a API começar a atender.
// O init.sql é idempotente (CREATE TABLE IF NOT EXISTS), então roda em todo start.
// O advisory lock faz réplicas que sobem ao mesmo tempo esperarem uma pela outra.
func runMigrations(db *sql.DB) error {
	ctx := context.Background()

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if _, err := tx.ExecContext(ctx, "SELECT pg_advisory_xact_lock($1)", migrationLockID); err != nil {
		return err
	}

	if _, err := tx.ExecContext(ctx, schemaSQL); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	log.Println("Schema do banco aplicado.")
	return nil
}
