package main

import (
	"errors"
	"github.com/Askalag/piece16/src"
	"github.com/Askalag/piece16/src/handler"
	"github.com/Askalag/piece16/src/log"
	"github.com/Askalag/piece16/src/repository"
	"github.com/Askalag/piece16/src/service"
	"github.com/Askalag/piece16/src/utils"
	"github.com/gin-gonic/gin"
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
	// initial schema and tables if not exists...
	repository.CreateTables(postgresDB)

	// repositories...
	repo := repository.NewTreeRepository(postgresDB)

	// services...
	srv := service.NewService(repo)

	// uow - unit of work...
	//uow := src.NewUOW(srv)

	// handler...
	h := handler.NewEngine(handler.MakeHandlers(srv))

	// server...
	serverStart(h, postgresDB)
}

func serverStart(h http.Handler, db *sqlx.DB) {
	cfg := src.LoadConfig()
	srv := src.NewServer(cfg)

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

	// .Env
	if err := godotenv.Load(); err != nil {
		log.FatalWithCode(2002, err.Error())
	}

	// setting log level
	level := utils.GetEnv("TREE_LOG_LEVEL", "")
	if level != "" {
		if err := log.InitLogger(level); err != nil {
			level = "debug"
			log.WarnWithCode(2001, err)
			err = log.InitLogger(level)
		}
		log.InfoWithCode(2003, "log Level: "+level)
	}

	// gin mode
	ginMode := utils.GetEnv("GIN_MODE", "debug")
	gin.SetMode(ginMode)
}

func loadConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		return errors.New("error while reading configuration file")
	}
	return nil
}
