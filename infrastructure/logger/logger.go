package logger

import (
	"zayyid-go/infrastructure/logger/logrus"
	"zayyid-go/infrastructure/logger/zap"
)

const (
	LOGRUS = "logrus"
	ZAP    = "zap"
)

var useLog string

func InitializeLogger(log string) {
	switch log {
	case LOGRUS:
		logrus.InitializeLogrusLogger()
		useLog = LOGRUS
	case ZAP:
		zap.InitializeZapLogger()
		useLog = ZAP
	default:
		logrus.InitializeLogrusLogger()
		useLog = LOGRUS
	}
}

func LogInfo(logtype, message string) {
	if useLog == LOGRUS {
		logrus.LogInfo(useLog, logtype, message)
	} else if useLog == ZAP {
		zap.LogInfo(useLog, logtype, message)
	}
}

func LogInfoWithData(data interface{}, logtype, message string) {
	if useLog == LOGRUS {
		logrus.LogInfoWithData(useLog, data, logtype, message)
	} else if useLog == ZAP {
		zap.LogInfoWithData(useLog, data, logtype, message)
	}
}

func LogError(logtype, errtype, message string) {
	if useLog == LOGRUS {
		logrus.LogError(useLog, logtype, errtype, message)
	} else if useLog == ZAP {
		zap.LogInfoWithData(useLog, logtype, errtype, message)
	}
}
