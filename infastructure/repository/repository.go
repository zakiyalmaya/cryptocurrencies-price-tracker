package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/go-redis/redis/v8"
	_ "github.com/mattn/go-sqlite3"
	"github.com/zakiyalmaya/cryptocurrencies-price-tracker/constant"
	"github.com/zakiyalmaya/cryptocurrencies-price-tracker/infastructure/repository/tracker"
	"github.com/zakiyalmaya/cryptocurrencies-price-tracker/infastructure/repository/user"
)

type Repositories struct {
	db      *sql.DB
	Redcl   *redis.Client
	User    user.UserRepository
	Tracker tracker.TrackerRepository
}

func NewRespository(db *sql.DB, redcl *redis.Client) *Repositories {
	return &Repositories{
		db:      db,
		Redcl:   redcl,
		User:    user.NewUserRepository(db),
		Tracker: tracker.NewTrackerRepository(db),
	}
}

func DBConnection() *sql.DB {
	db, err := sql.Open("sqlite3", constant.SQLiteDBName)
	if err != nil {
		log.Panicln("error connecting to database: ", err.Error())
		return nil
	}

	createUserTable(db)
	createUserTrackedCoinTable(db)
	return db
}

func createUserTable(db *sql.DB) {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name VARCHAR(255) NOT NULL,
		email VARCHAR(255) NOT NULL,
		username VARCHAR(255) UNIQUE NOT NULL,
		password TEXT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	)`)
	if err != nil {
		log.Panicln("error creating users table: ", err.Error())
	}
}

func createUserTrackedCoinTable(db *sql.DB) {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS user_tracked_coins (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id VARCHAR(255) NOT NULL,
		coin_id VARCHAR(255) NOT NULL,
		coin_symbol VARCHAR(255) NOT NULL,
		coin_name VARCHAR(255) NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	)`)
	if err != nil {
		log.Panicln("error creating user_tracked_coins table: ", err.Error())
	}
}

func RedisClient() *redis.Client {
	option := &redis.Options{
		Addr:     constant.RedisHost,
		Password: constant.RedisPass,
		DB:       0,
	}

	redcl := redis.NewClient(option)
	if err := redcl.Ping(context.Background()).Err(); err != nil {
		log.Panicln("error connect to redis: ", err.Error())
		return nil
	}

	return redcl
}
