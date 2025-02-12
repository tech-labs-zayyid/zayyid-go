package zap

import (
	"os"
	"time"

	"go.uber.org/zap"
)

var log *zap.SugaredLogger

func InitializeZapLogger() {
	logger := zap.NewExample()
	defer func() {
		if err := logger.Sync(); err != nil {
			// Handle the error, e.g., log it or panic
			panic(err)
		}
	}()
	log = logger.Sugar()
}

func LogInfo(logUse, logtype, message string) {
	log.Infow(message,
		"app_name", os.Getenv("APP_NAME"),
		"app_version", os.Getenv("APP_VERSION"),
		"log_type", logtype,
		"time", time.Now().Format(time.RFC3339),
		"log", logUse,
	)
}

func LogInfoWithData(logUse string, data interface{}, logtype, message string) {
	log.Infow(message,
		"app_name", os.Getenv("APP_NAME"),
		"app_version", os.Getenv("APP_VERSION"),
		"data", data,
		"log_type", logtype,
		"time", time.Now().Format(time.RFC3339),
		"log", logUse,
	)

}

func LogError(logUse, logtype, errtype, message string) {
	log.Errorw(message,
		"app_name", os.Getenv("APP_NAME"),
		"app_version", os.Getenv("APP_VERSION"),
		"error_type", errtype,
		"log_type", logtype,
		"time", time.Now().Format(time.RFC3339),
		"log", logUse,
	)
}
