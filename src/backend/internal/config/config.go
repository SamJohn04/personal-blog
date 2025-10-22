package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type Config struct {
	JWTSecret string
}

var (
	Cfg Config
	DB  *sql.DB
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env found. Ignore if it is not needed.")
	}
	initDB()

	Cfg = Config{
		JWTSecret: os.Getenv("JWT_SECRET"),
	}
}

func initDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error opening DB:", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Error connecting to DB:", err)
	}

	DB = db
	log.Println("Connected to MariaDB")
}
