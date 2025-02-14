package logrus

import (
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

func InitializeLogrusLogger() {
	log.SetFormatter(&log.JSONFormatter{})

	// currentTime := time.Now()
	// date := currentTime.Format("20060102")
	// path := fmt.Sprintf("%s/%s-%s.%s", os.Getenv("LOG_PATH"), os.Getenv("LOG_PREFIX"), date, os.Getenv("LOG_EXT"))
	// err := os.MkdirAll(filepath.Dir(path), 0770)
	// if err == nil {
	// 	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	// 	if err != nil {
	// 		log.SetOutput(os.Stdout)
	// 		return
	// 	}
	// 	log.SetOutput(file)
	// } else {
	// 	log.SetOutput(os.Stdout)
	// }

	if strings.ToLower(os.Getenv("ENV")) == "dev" || strings.ToLower(os.Getenv("ENVIRONMENT")) == "dev" {
		log.SetOutput(os.Stdout)
	}
	// else {
	// 	err := os.MkdirAll(filepath.Dir(path), 0770)
	// 	if err == nil {
	// 		file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	// 		if err != nil {
	// 			log.SetOutput(os.Stdout)
	// 			return
	// 		}
	// 		log.SetOutput(file)
	// 	} else {
	// 		log.SetOutput(os.Stdout)
	// 	}
	// }
}

func LogInfo(logUse, logtype, message string) {
	log.WithFields(log.Fields{
		"app_name":    os.Getenv("APP_NAME"),
		"app_version": os.Getenv("APP_VERSION"),
		"log_type":    logtype,
		"log":         logUse,
	}).Info(message)
}

func LogInfoWithData(logUse string, data interface{}, logtype, message string) {
	log.WithFields(log.Fields{
		"app_name":    os.Getenv("APP_NAME"),
		"app_version": os.Getenv("APP_VERSION"),
		"data":        data,
		"log_type":    logtype,
		"log":         logUse,
	}).Info(message)
}

func LogError(logUse, logtype, errtype, message string) {
	log.WithFields(log.Fields{
		"app_name":    os.Getenv("APP_NAME"),
		"app_version": os.Getenv("APP_VERSION"),
		"error_type":  errtype,
		"log_type":    logtype,
		"log":         logUse,
	}).Error(message)
}
