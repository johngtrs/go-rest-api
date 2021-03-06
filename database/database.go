package database

import (
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func Initialize() (*sqlx.DB, error) {
	cfg := mysql.Config{
		User:                 os.Getenv("MYSQL_USER"),
		Passwd:               os.Getenv("MYSQL_PASSWORD"),
		Net:                  "tcp",
		Addr:                 os.Getenv("MYSQL_HOST") + ":" + os.Getenv("MYSQL_LOCAL_PORT"),
		DBName:               os.Getenv("MYSQL_DATABASE"),
		AllowNativePasswords: true,
	}

	// Get a database handle
	var err error
	db, err := sqlx.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	return db, err
}
