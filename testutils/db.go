package testutils

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	mysqlMigrate "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func SetupTestDB(t *testing.T) *sql.DB {
	dsn := "admin:admin@tcp(localhost:3306)/JobSeeker_test?parseTime=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		t.Fatal("Failed to connect to test DB:", err)
	}
	
	err = db.Ping()
	if err != nil {
		t.Fatal("Failed to ping DB:", err)
	}
	
	err = runMigrations(db)
	if err != nil {
		t.Fatal("Migrations failed:", err)
	}
	
	_, err = db.Exec("SET FOREIGN_KEY_CHECKS = 0")
	if err != nil {
		t.Fatal("Failed to disable FK checks:", err)
	}
	
	_, err = db.Exec("DELETE FROM Company")
	if err != nil {
		t.Fatal("Failed to clean Company:", err)
	}
	
	_, err = db.Exec("DELETE FROM JobSeeker")
	if err != nil {
		t.Fatal("Failed to clean JobSeeker:", err)
	}
	
	_, err = db.Exec("DELETE FROM `User`") 
	if err != nil {
		t.Fatal("Failed to clean User:", err)
	}
	
	_, err = db.Exec("SET FOREIGN_KEY_CHECKS = 1")
	if err != nil {
		t.Fatal("Failed to enable FK checks:", err)
	}
	
	return db
}

func runMigrations(db *sql.DB) error {
	dbName := "jobSeeker_test"

	driver, err := mysqlMigrate.WithInstance(db, &mysqlMigrate.Config{
		DatabaseName: dbName,
	})
	if err != nil {
		return err
	}
	
	m, err := migrate.NewWithDatabaseInstance(
		"file:///home/krimino/Documents/projects/go-projects/JobSeekerAPI/cmd/migrate/migrations",
		"mysql",
		driver,
	)
	if err != nil {
		return err
	}
	
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}
	
	return nil
}
