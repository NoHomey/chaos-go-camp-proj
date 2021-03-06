package open

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

//DB represents information for the database to be opened.
type DB struct {
	User string
	Pass string
	Name string
}

//Open opens a MySQL database.
func Open(conn DB) (*sql.DB, error) {
	connURI := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", conn.User, conn.Pass, conn.Name)
	db, err := sql.Open("mysql", connURI)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxIdleTime(time.Minute)
	db.SetConnMaxLifetime(5 * time.Minute)
	db.SetMaxIdleConns(25)
	db.SetMaxOpenConns(100)
	return db, nil
}

//Ping pings the given db.
func Ping(ctx context.Context, db *sql.DB) error {
	pCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	return db.PingContext(pCtx)
}
