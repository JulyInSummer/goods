package pg_storage

import (
	"fmt"
	"log/slog"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/jmoiron/sqlx"
	"goods_project/internal/config"
	"goods_project/internal/storage"
)

type pgStorage struct {
	log *slog.Logger
	db  *sqlx.DB
}

func NewPGStorage(cfg *config.Config, log *slog.Logger) storage.Storage {
	connString := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.Name,
		cfg.DB.SSL,
	)

	db, err := sqlx.Open("postgres", connString)
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	return &pgStorage{
		log: log,
		db:  db,
	}
}
