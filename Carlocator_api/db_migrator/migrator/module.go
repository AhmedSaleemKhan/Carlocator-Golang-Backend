package migrator

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/pkg/errors"
	migrate "github.com/rubenv/sql-migrate"
)

type Migrator struct {
	Source        migrate.FileMigrationSource
	DB            *sql.DB
	Operation     string
	MaxMigrations int
}

func ProvideMigrator(db *sql.DB, operation string, maxMigrations int) *Migrator {
	migrations := &migrate.FileMigrationSource{
		Dir: "./migrations",
	}

	return &Migrator{
		Source:        *migrations,
		DB:            db,
		Operation:     operation,
		MaxMigrations: maxMigrations,
	}
}

func (m *Migrator) RunMigrations() (int, error) {
	switch m.Operation {
	case "up":
		return m.migrateUp()
	case "down":
		return m.migrateDown()
	default:
		fmt.Println("No Migration flag provided")
	}
	return 0, nil
}

func (m *Migrator) migrateUp() (int, error) {
	// migrate.SetIgnoreUnknown(false)
	n, err := migrate.ExecMax(m.DB, "postgres", m.Source, migrate.Up, m.MaxMigrations)
	if err != nil {
		return 0, errors.Wrap(err, "migrator could not complete")
	}
	log.Printf("Applied %d migrations", n)
	return n, nil
}

func (m *Migrator) migrateDown() (int, error) {
	// migrate.SetIgnoreUnknown(false)
	n, err := migrate.ExecMax(m.DB, "postgres", m.Source, migrate.Down, m.MaxMigrations)
	if err != nil {
		return 0, errors.Wrap(err, "migrator could not complete")
	}
	log.Printf("Applied %d migrations", n)
	return n, nil
}
