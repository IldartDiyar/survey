package infra

import (
	"database/sql"
	"survey/config"
)

type Container struct {
	Conn *sql.DB
}

func New(conf *config.Config) (*Container, error) {
	conn, err := postgreConnect(conf.PostgresDB)

	if err != nil {
		return nil, err
	}

	return &Container{Conn: conn}, nil

}
