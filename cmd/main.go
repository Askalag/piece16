package main

import (
	"errors"
	"github.com/Askalag/piece16"
	"github.com/Askalag/piece16/src/handler"
	"github.com/Askalag/piece16/src/log"
	"github.com/Askalag/piece16/src/repository"
	"github.com/Askalag/piece16/src/service"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
)

func main() {

	loadRootConfig()

	// postgres...
	postgresConfig := repository.LoadPostgresConfig()
	postgresDB, ok := repository.NewPostgresDB(postgresConfig)
	if ok != nil {
		log.FatalWithCode(3000, ok.Error())
	}

	// repositories...
	repo := repository.NewTreeRepository(postgresDB)

	// services...
	srv := service.NewService(repo)

	// uow unit of work...
	//uow := src.NewUOW(srv)

	// handler...
	h := handler.NewRootHandler(handler.MakeHandlers(srv))

	// server...
	serverStart(h, postgresDB)
}

func serverStart(h http.Handler, db *sqlx.DB) {
	cfg := piece16.LoadConfig()
	srv := piece16.NewServer(cfg)

	log.InfoWithCode(4001)
	go func() {
		log.InfoWithCode(4002, srv.GetAddr())
		if err := srv.Run(h); err != http.ErrServerClosed && err != nil {
			log.FatalWithCode(4000, err.Error())
		}
	}()

	srv.GracefulShutdown()

	if err := db.Close(); err != nil {
		log.InfoWithCode(3003, err.Error())
	}
}

func loadRootConfig() {
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
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		return errors.New("error while reading configuration file")
	}
	return nil
}
