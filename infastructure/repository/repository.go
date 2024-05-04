package repository

import (
	"log"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	_ "github.com/miekg/pkcs11"
	"github.com/zakiyalmaya/cryptocurrencies-price-tracker/infastructure/repository/user"
)

type Repositories struct {
	db   *sql.DB
	User user.UserRepository
}

func NewRespository(db *sql.DB) *Repositories {
	return &Repositories{
		db:   db,
		User: user.NewUserRepository(db),
	}
}

func DBConnection() *sql.DB {
	db, err := sql.Open("sqlite3", "./cryptocurrencies.db")
	if err != nil {
		log.Panicln("error connecting to database: ", err.Error())
		return nil
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name VARCHAR(255) NOT NULL,
		email VARCHAR(255) NOT NULL,
		username VARCHAR(255) UNIQUE NOT NULL,
		password TEXT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	)`)
	if err != nil {
		log.Panicln("error creating table: ", err.Error())
		return nil
	}

	return db
}
