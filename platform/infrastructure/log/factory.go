package log

import (
	"go.uber.org/zap"
)

//Build construct the logger
func Build() (*zap.SugaredLogger, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}
	return logger.Sugar(), nil
}
