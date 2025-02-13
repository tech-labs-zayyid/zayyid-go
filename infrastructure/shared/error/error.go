package error

import (
	"errors"
	"fmt"
	"strings"
	"zayyid-go/infrastructure/logger"
	sharedConstant "zayyid-go/infrastructure/shared/constant"
)

func New(tipe string, msg string, err error) error {
	if err == nil {
		logger.LogError(sharedConstant.ERR, tipe, fmt.Sprintf(msg+" undefined error"))
	} else {
		logger.LogError(sharedConstant.ERR, tipe, fmt.Sprintf(msg+err.Error()))
	}

	return fmt.Errorf("%s | %s: %w", tipe, msg, err)
}

func TrimMesssage(err error) (tipe string, newErr error) {
	errs := strings.Split(err.Error(), "|")
	tipe = strings.TrimSpace(errs[0])

	newErr = errors.New(strings.TrimSpace(errs[1]))
	return
}
