package infra

import (
	"database/sql"
	"fmt"

	"survey/config"

	_ "github.com/lib/pq"
)

func postgreConnect(config config.PostgresDB) (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.Username, config.Password, config.Database)

	conn, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}

	if err := conn.Ping(); err != nil {
		return nil, err
	}

	return conn, nil
}
