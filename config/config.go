package config

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

type Config struct {
	MasterDbUrl string
	MasterDB    *sqlx.DB
}

var (
	config Config
)

func Initialize() error {
	if (config != Config{}) {
		return fmt.Errorf("config already initialized")
	}

	err := godotenv.Load()
	if err != nil {
		return err
	}

	masterDB, err := sqlx.Connect("postgres", os.Getenv("MASTER_DB_URL"))
	if err != nil {
		return err
	}

	config = Config{
		MasterDbUrl: os.Getenv("MASTER_DB_URL"),
		MasterDB:    masterDB,
	}

	return nil
}

func Get() Config {
	return config
}
