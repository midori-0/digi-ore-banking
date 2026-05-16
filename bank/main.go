package main

import (
	"database/sql"
	"log/slog"
	"os"

	_ "modernc.org/sqlite"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	db, err := sql.Open("sqlite", "/Users/carson/Documents/digi-ore-banking/digi-ore-banking.db")
	if err != nil {
		logger.Error("failed to open database", "error", err)
		os.Exit(1)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		logger.Error("failed to ping database", "error", err)
		os.Exit(1)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS test (id INTEGER PRIMARY KEY)`)
	if err != nil {
		logger.Error("failed to create table", "error", err)
		os.Exit(1)
	}

	logger.Info("database opened successfully")
}
