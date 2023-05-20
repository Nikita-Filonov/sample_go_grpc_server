package database

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"sample_go_grpc_server/utils/config"
	"sample_go_grpc_server/utils/logger"
)

type Client struct {
	db     *sql.DB
	logger *logger.CtxLogger
}

func NewClient(config config.Config, loggerService logger.Service, space string) Client {
	db, err := sql.Open("sqlite3", config.Database.File)

	if err != nil {
		panic(err)
	}

	return Client{
		db:     db,
		logger: loggerService.NewPrefix(fmt.Sprintf("DATABASE.%s", space)),
	}
}

func (c *Client) QueryRow(query string, args ...any) *sql.Row {
	c.logger.Info(query)
	return c.db.QueryRow(query, args...)
}

func (c *Client) Exec(query string, args ...any) (sql.Result, error) {
	c.logger.Info(query)
	return c.db.Exec(query, args...)
}
