package database

import (
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
)

var (
	db *sqlx.DB
)

func getConnect() string {
	return fmt.Sprintf(
		`%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&maxAllowedPacket=%d`,
		"user",
		"password",
		"host",
		"database",
		"utf8mb4",
		0,
	)
}

func Connect() *sqlx.DB {
	if db == nil {
		conn, err := sqlx.Connect("mysql", getConnect())
		if err != nil {
			log.Panicln(err.Error())
			return nil
		}
		db = conn

		db.SetConnMaxLifetime(time.Minute)
		db.SetConnMaxIdleTime(time.Minute)

		db.SetMaxIdleConns(30)
		db.SetMaxOpenConns(30)
	}
	return db
}
