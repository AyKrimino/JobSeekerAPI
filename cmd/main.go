package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/AyKrimino/JobSeekerAPI/cmd/api"
	"github.com/AyKrimino/JobSeekerAPI/config"
	"github.com/AyKrimino/JobSeekerAPI/db"
	"github.com/go-sql-driver/mysql"
)

// @title JobSeeker API
// @version 1.0
// @description API for managing JobSeeker and Company user registrations.
// @host localhost:8080
// @BasePath /
func main() {
	cfg := mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	db, err := db.NewMySQLStorage(cfg)
	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)

	server := api.NewAPIServer(fmt.Sprintf(":%s", config.Envs.Port), db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB: Successfully Connected!")
}
