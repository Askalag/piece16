package repository

import (
	"fmt"
	"github.com/Askalag/piece16/src/log"
	"github.com/Askalag/piece16/src/utils"
	"github.com/spf13/viper"
)
import "github.com/jmoiron/sqlx"

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func (c *Config) connStr() string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		c.Host, c.Port, c.Username, c.DBName, c.Password, c.SSLMode)
}

func NewPostgresDB(c *Config) (*sqlx.DB, error) {
	db, err := sqlx.Connect(viper.GetString("db.postgres.driver"), c.connStr())
	if err != nil {
		log.FatalWithCode(3000, err.Error())
	}
	return db, err
}

func LoadPostgresConfig() *Config {
	return &Config{
		Host:     utils.GetEnv("TREE_POSTGRES_HOST", ""),
		Port:     utils.GetEnv("TREE_POSTGRES_PORT", ""),
		Username: utils.GetEnv("TREE_POSTGRES_USERNAME", ""),
		Password: utils.GetEnv("TREE_POSTGRES_PASSWORD", ""),
		DBName:   utils.GetEnv("TREE_POSTGRES_DBNAME", ""),
		SSLMode:  utils.GetEnv("TREE_POSTGRES_SLLMODE", "disable"),
	}
}
