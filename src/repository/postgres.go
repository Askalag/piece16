package repository

import (
	"fmt"
	"github.com/Askalag/piece16/src/log"
	"github.com/Askalag/piece16/src/utils"
	"github.com/spf13/viper"
	"io/ioutil"
	"path/filepath"
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

// connStr - Connection string
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

// LoadPostgresConfig - db config
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

// CreateTables create tables if not exists
func CreateTables(db *sqlx.DB) {
	path := filepath.Join("src/migration/sql/", "initial_tables.sql")
	f, err := ioutil.ReadFile(path)
	if err != nil {
		log.WarnWithCode(3004, err.Error())
	}

	sql := string(f)
	_, err = db.Exec(sql)
	if err != nil {
		log.WarnWithCode(3005, err.Error())
	}
}
