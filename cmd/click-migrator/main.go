package main

import (
	"fmt"
	"log"

	_ "github.com/ClickHouse/clickhouse-go"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/clickhouse"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/pkg/errors"
	"goods_project/internal/config"
)

func main() {
	cfg := config.NewConfig()
	connStr := fmt.Sprintf("clickhouse://%s:%s?username=%s&password=%s&database=%s&x-multi-statement=true&x-migrations-table-engine=MergeTree",
		cfg.Ch.Host,
		cfg.Ch.Port,
		cfg.Ch.Username,
		cfg.Ch.Password,
		cfg.Ch.DB,
	)

	m, err := migrate.New("file://clickhouse-migrations", connStr)
	if err != nil {
		panic(err)
	}

	if err = m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			log.Println("no migrations to apply")
		}
		panic(err)
	}

	log.Println("migrations has successfully applied")
}
