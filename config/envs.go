package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type Config struct {
	PublicHost     string
	Port           string
	DBUser         string
	DBName         string
	DBAddress      string
	DBPassword     string
	JWTSecret      string
	TestDBUser     string
	TestDBName     string
	TestDBAddress  string
	TestDBPassword string
}

var Envs = initConfig()

func initConfig() Config {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal("Error getting current working directory:", err)
	}

	envPath := filepath.Join(wd, ".env")
	err = godotenv.Load(envPath)
	if err != nil {
		log.Println("Warning: No .env file found, using default values")
	} else {
		log.Println(".env file loaded successfully")
	}

	return Config{
		PublicHost: getEnv("PUBLIC_HOST", "http://localhost"),
		Port:       getEnv("PORT", "8080"),
		DBUser:     getEnv("DB_USER", "admin"),
		DBName:     getEnv("DB_NAME", "JobSeeker"),
		DBAddress: fmt.Sprintf(
			"%s:%s",
			getEnv("DB_HOST", "127.0.0.1"),
			getEnv("DB_PORT", "3306"),
		),
		DBPassword: getEnv("DB_PASSWORD", "admin"),
		JWTSecret:  getEnv("JWT_SECRET", "secret-key"),
		TestDBUser: getEnv("TEST_DB_USER", "admin"),
		TestDBName: getEnv("TEST_DB_NAME", "JobSeeker_test"),
		TestDBAddress: fmt.Sprintf(
			"%s:%s",
			getEnv("TEST_DB_HOST", "127.0.0.1"),
			getEnv("TEST_DB_PORT", "3306"),
		),
		TestDBPassword: getEnv("TEST_DB_PASSWORD", "admin"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
