package clickhouse_storage

import (
	"context"
	"crypto/tls"
	"fmt"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"goods_project/internal/config"
)

func NewCh(cfg *config.Config) driver.Conn {
	ctx := context.Background()

	address := fmt.Sprintf("%s:%s", cfg.Ch.Host, cfg.Ch.Port)

	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{address},
		Auth: clickhouse.Auth{
			Database: cfg.Ch.DB,
			Username: cfg.Ch.Username,
			Password: cfg.Ch.Password,
		},
		ClientInfo: clickhouse.ClientInfo{Products: []struct {
			Name    string
			Version string
		}{
			{Name: "goods-logs", Version: "0.1"},
		}},
		Debugf: func(format string, v ...interface{}) {
			fmt.Printf(format, v)
		},
		TLS: &tls.Config{
			InsecureSkipVerify: true,
		},
	})

	if err != nil {
		panic(err)
	}

	if err = conn.Ping(ctx); err != nil {
		if exception, ok := err.(*clickhouse.Exception); ok {
			fmt.Printf("Exception [%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
		}
		panic(err)
	}

	return conn
}
