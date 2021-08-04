package log

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var language = "en"

const (
	GinLog = "gin"
	AppLog = "app"
)

// InitLogger logger initialization
// Format, Level, TimestampFormat...
func InitLogger(level string) error {
	log.SetFormatter(&log.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	lv, err := log.ParseLevel(level)
	if err != nil {
		return err
	}
	log.SetLevel(lv)
	return nil
}

func InfoWithCode(c int, args ...interface{}) {
	key := fmt.Sprintf("log.codes.%d%s", c, language)
	log.WithFields(log.Fields{
		"logType": AppLog,
		"logCode": c,
		"logMsg":  viper.GetString(key),
	}).Info(args)
}

func WarnWithCode(c int, args ...interface{}) {
	key := fmt.Sprintf("log.codes.%d%s", c, language)
	log.WithFields(log.Fields{
		"logType": AppLog,
		"logCode": c,
		"logMsg":  viper.GetString(key),
	}).Warn(args)
}

func FatalWithCode(c int, args ...interface{}) {
	key := fmt.Sprintf("log.codes.%d%s", c, language)
	log.WithFields(log.Fields{
		"logType": AppLog,
		"logCode": c,
		"logMsg":  viper.GetString(key),
	}).Fatal(args)
}

func GinLogInfo(lv string, cip string, code int, method string, uri string, latSec float64) {
	log.WithFields(log.Fields{
		"logType":    GinLog,
		"level":      lv,
		"clientIP":   cip,
		"statusCode": code,
		"reqMethod":  method,
		"reqUri":     uri,
		"latencySec": latSec,
	}).Info()
}
