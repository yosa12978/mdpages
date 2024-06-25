package data

import (
	"database/sql"
	"os"
	"sync"
)

var (
	pg     *sql.DB
	pgOnce sync.Once
)

func Postgres() *sql.DB {
	pgOnce.Do(func() {
		conn, err := sql.Open("postgres", os.Getenv("POSTGRES"))
		if err != nil {
			panic(err)
		}
		if err := conn.Ping(); err != nil {
			panic(err)
		}
		pg = conn
	})
	return pg
}
