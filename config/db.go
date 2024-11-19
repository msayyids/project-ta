package config

import (
	"fmt"
	"os"
	"project-ta/helper"

	"github.com/jmoiron/sqlx"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

var (
	user     = os.Getenv("DB_USERNAME")
	dbname   = os.Getenv("DB_DBNAME")
	sslmode  = os.Getenv("DB_SSLMODE")
	password = os.Getenv("DB_PASSWORD")
	host     = os.Getenv("DB_HOST")
)

func ConnectionString() string {
	return fmt.Sprintf("user=%s dbname=%s sslmode=%s password=%s host=%s", user, dbname, sslmode, password, host)
}

func ConnecDb() sqlx.DB {
	db, err := sqlx.Connect("postgres", ConnectionString())
	helper.PanicIfError(err)

	return *db
}
