package db

import (
	"github.com/fazriridwan19/service-employee/config"
	"github.com/jmoiron/sqlx"
)

func ConnectMysql(cfg *config.Config) (*sqlx.DB, error) {
	db, error := sqlx.Connect("mysql", cfg.DatabaseConnection())
	if error != nil {
		return nil, error
	}
	return db, nil
}
