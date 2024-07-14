package logging

import "go.uber.org/zap"

var sugar *zap.SugaredLogger

func NewLogger() *zap.SugaredLogger {
	if sugar != nil {
		return sugar
	}
	logger, _ := zap.NewProduction()
	sugar = logger.Sugar()
	return sugar
}
