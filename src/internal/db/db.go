package db

import (
	"database/sql"
	"github.com/Antoniel03/social-network-api/internal/env"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func runMigrations(db *sql.DB) {
	migrateDBName := env.GetString("MIGRATE_DB_NAME", "postgres")

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("Error while getting migration driver: %v", err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://./internal/db/migrations",
		migrateDBName, driver)
	if err != nil {
		log.Fatalf("Error while getting migration instance: %v", err)
	}
	if err := m.Migrate(1); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
}

func SetupDB() *sql.DB {
	user := env.GetString("DATABASE_USER", "postgres")
	host := env.GetString("DATABASE_HOST", "")
	db_name := env.GetString("DATABASE_NAME", "mydb")
	sslmode := env.GetString("SSLMODE", "disable")
	password := env.GetString("DATABASE_PASSWORD", "")

	log.Printf("%s, %s", user, db_name)

	connStr := "user=" + user + " dbname=" + db_name + " sslmode=" + sslmode
	if host != "" && password != "" {
		connStr = connStr + " host=" + host + " password=" + password
	}
	log.Println(connStr)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	runMigrations(db)
	return db
}

// Replaced by migrations
func CreateTables(connection *sql.DB) error {
	query, err := os.ReadFile("./internal/db/scripts/initDB.sql")
	if err != nil {
		return err
	}
	_, err = connection.Exec(string(query))
	if err != nil {
		return err
	}
	return nil
}
