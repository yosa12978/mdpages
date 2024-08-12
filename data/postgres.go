package data

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/lib/pq"
	"github.com/yosa12978/mdpages/config"
)

var (
	pg     *sql.DB
	pgOnce sync.Once
)

func Postgres() *sql.DB {
	pgOnce.Do(func() {
		cfg := config.Get()
		addr := fmt.Sprintf(
			"postgres://%s:%s@%s",
			cfg.Postgres.Username,
			cfg.Postgres.Password,
			cfg.Postgres.Addr,
		)
		conn, err := sql.Open("postgres", addr)
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
