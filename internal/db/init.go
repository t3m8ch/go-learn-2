package db

import (
	"context"
	"os"

	pgxdecimal "github.com/jackc/pgx-shopspring-decimal"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

func InitDb(dbUrl string, logger *zap.Logger) (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), "postgres://t3m8ch@localhost/productsdb")
	if err != nil {
		return nil, err
	}

	pgxdecimal.Register(conn.TypeMap())

	createTableSql, err := os.ReadFile("./db/schema.sql")
	if err != nil {
		return nil, err
	}

	_, err = conn.Exec(context.Background(), string(createTableSql))
	if err != nil {
		return nil, err
	}

	return conn, nil
}
