package migrations

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/sqlite3"
	"github.com/golang-migrate/migrate/source/file"
)

func GetMigrator(db *sql.DB) *migrate.Migrate {
	fmt.Println("Running migrate up command")

	dbDriver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		fmt.Printf("instance error: %v \n", err)
	}

	fileSource, err := (&file.File{}).Open("file://migrations/sql")
	if err != nil {
		fmt.Printf("opening file error: %v \n", err)
	}

	m, err := migrate.NewWithInstance("file", fileSource, "core", dbDriver)
	if err != nil {
		fmt.Printf("migrate error: %v \n", err)
	}

	return m
}

func MigrateUp(db *sql.DB) {
	migrator := GetMigrator(db)

	if err := migrator.Up(); err != nil {
		fmt.Printf("migrate up error: %v \n", err)
	}

	fmt.Println("Migrate up done with success")
}

func MigrateDown(db *sql.DB) {
	migrator := GetMigrator(db)

	if err := migrator.Down(); err != nil {
		fmt.Printf("migrate down error: %v \n", err)
	}

	fmt.Println("Migrate up done with success")
}
