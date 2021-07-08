package main

import (
	"errors"
	piece "github.com/Askalag/piece16/src"
	"github.com/Askalag/piece16/src/log"
	"github.com/Askalag/piece16/src/repository"
	"github.com/Askalag/piece16/src/service"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

func main()  {

	Init()

	postgresConfig := &repository.Config{
		Host: GetEnv("TREE_POSTGRES_HOST", "localhost"),
		Port: GetEnv("TREE_POSTGRES_PORT", "5432"),
		Username: GetEnv("TREE_POSTGRES_USERNAME", ""),
		Password: GetEnv("TREE_POSTGRES_PASSWORD", ""),
		DBName: GetEnv("TREE_POSTGRES_DBNAME", ""),
		SSLMode: GetEnv("TREE_POSTGRES_SLLMODE", "disable"),
	}
	// dataBases...
	postgresDB, ok := repository.NewPostgresDB(postgresConfig)
	if ok != nil {
		log.FatalWithCode(3000, ok.Error())
	}

	// repositories...
	repo := repository.NewTreeRepository(postgresDB)

	// services...
	srv := service.NewService(repo)

	// uow unit of work...
	uow := piece.NewUOW(srv)
	uow.S1.Tree.Create()
}

func Init()  {
	// Config...
	if err := loadConfig(); err != nil {
		logrus.Fatal(err)
		return
	}
	log.InitLogger(viper.GetString("log.level"))

	// .Env
	if err := godotenv.Load(); err != nil {
		log.FatalWithCode(2002, err.Error())
	}
}

func loadConfig() error {
	viper.AddConfigPath("src/config")
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err !=nil {
		return errors.New("error while reading configuration file")
	}
	return nil
}

func GetEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
