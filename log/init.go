package log

import "go.uber.org/zap"

var logger *zap.Logger
var sugar *zap.SugaredLogger

func init() {
	logger, _ = zap.NewProduction()
	defer logger.Sync()
	sugar = logger.Sugar()
}

func Sugar() *zap.SugaredLogger {
	return sugar
}
