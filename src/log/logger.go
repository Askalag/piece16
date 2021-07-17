package log

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var language = "en"

// InitLogger logger initialization
// Format, Level, TimestampFormat...
func InitLogger(level string) {
	log.SetFormatter(&log.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	lv, err := log.ParseLevel(level)
	if err != nil {
		WarnWithCode(1002)
		log.SetLevel(log.InfoLevel)
	}
	log.SetLevel(lv)
}

func InfoWithCode(c int, args ...interface{}) {
	key := fmt.Sprintf("log.codes.%d%s", c, language)
	log.WithFields(log.Fields{
		"logCode": c,
		"logMsg":  viper.GetString(key),
	}).Warn(args)
}

func WarnWithCode(c int, args ...interface{}) {
	key := fmt.Sprintf("log.codes.%d%s", c, language)
	log.WithFields(log.Fields{
		"logCode": c,
		"logMsg":  viper.GetString(key),
	}).Warn(args)
}

func FatalWithCode(c int, args ...interface{}) {
	key := fmt.Sprintf("log.codes.%d%s", c, language)
	log.WithFields(log.Fields{
		"logCode": c,
		"logMsg":  viper.GetString(key),
	}).Fatal(args)
}
