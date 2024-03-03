package main

import (
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/pkg/errors"
	"goods_project/internal/config"
)

func main() {
	cfg := config.NewConfig()

	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", cfg.DB.User, cfg.DB.Password, cfg.DB.Host, cfg.DB.Port, cfg.DB.Name, cfg.DB.SSL)

	m, err := migrate.New("file://migrations", connString)
	if err != nil {
		panic(err)
	}

	if err = m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			log.Println("no migrations to apply")
			return
		}
		panic(err)
	}

	log.Println("migrations has been successfully applied")
}
