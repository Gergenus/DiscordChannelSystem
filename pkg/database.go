package pkg

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Gergenus/config"
	_ "github.com/lib/pq"
)

type PostgresDatabase struct {
	Db *sql.DB
}

func NewPostgresDatabase(cfg *config.Config) PostgresDatabase {
	z := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.Db.User,
		cfg.Db.Password,
		cfg.Db.Host,
		cfg.Db.Port,
		cfg.Db.Dbname,
		cfg.Db.Sslmode,
	)

	conn, err := sql.Open("postgres", z)
	if err != nil {
		log.Fatal("db connecting error", err)
	}
	zv := PostgresDatabase{Db: conn}
	return zv
}

func (p PostgresDatabase) GetDB() *sql.DB {
	return p.Db
}
