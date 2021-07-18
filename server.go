package piece16

import (
	"context"
	"fmt"
	"github.com/Askalag/piece16/src/log"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Config struct {
	Addr         string
	IdleTimeout  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type Server struct {
	httpServer *http.Server
	config     *Config
}

func (s *Server) GetAddr() string {
	return s.config.Addr
}

func NewServer(c *Config) *Server {
	srv := new(Server)
	srv.config = c
	return srv
}

func (s *Server) Run(h http.Handler) error {
	s.httpServer = &http.Server{
		Addr:         s.config.Addr,
		Handler:      h,
		IdleTimeout:  s.config.IdleTimeout,
		ReadTimeout:  s.config.ReadTimeout,
		WriteTimeout: s.config.WriteTimeout,
	}
	return s.httpServer.ListenAndServe()
}

func LoadConfig() *Config {
	config := &Config{
		Addr: fmt.Sprintf("%s:%s",
			GetEnv("TREE_SERVER_HOST", ""),
			GetEnv("TREE_SERVER_PORT", "")),
		IdleTimeout:  viper.GetDuration("server.idleTimeout"),
		ReadTimeout:  viper.GetDuration("server.readTimeout"),
		WriteTimeout: viper.GetDuration("server.writeTimeout"),
	}
	return config
}

func (s *Server) GracefulShutdown() {
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT)
	<-sigChan
	log.InfoWithCode(4003)
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	_ = s.httpServer.Shutdown(tc)
}
