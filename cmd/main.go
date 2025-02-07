package main

import (
	"database/sql"
	"log"

	"github.com/AyKrimino/JowseekerAPI/cmd/api"
	"github.com/AyKrimino/JowseekerAPI/db"
	"github.com/go-sql-driver/mysql"
)

func main() {
	cfg := mysql.Config{
		User: "admin",
		Passwd: "admin",
		Addr: "127.0.0.1:3306",
		DBName: "JobSeeker",
		Net: "tcp",
		AllowNativePasswords: true,
		ParseTime: true,
	}

	db, err := db.NewMySQLStorage(cfg)
	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)

	server := api.NewAPIServer(":8080")
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