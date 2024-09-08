package utils

import (
	"fmt"
	"log"

	"github.com/dealense7/market-price-go/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func ConnectDatabase() error {
	var err error
	DB, err = sqlx.Connect(
		config.Envs.DBDriver,
		fmt.Sprintf(
			"%s:%s@(localhost:3306)/%s?parseTime=true",
			config.Envs.DBUsername,
			config.Envs.DBPassword,
			config.Envs.DBName,
		),
	)
	if err != nil {
		return fmt.Errorf("could not connect to the database: %w", err)
	}

	err = DB.Ping()
	if err != nil {
		return fmt.Errorf("failed to ping the database: %w", err)
	}

	log.Println("Connected to the MySQL database successfully!")
	return nil
}
