package main

import (
	"go.uber.org/zap"
)

var ZapLogger *zap.Logger

func init() {
	config := zap.NewProductionConfig()
	config.OutputPaths = []string{"stdout"}
	ZapLogger, _ = config.Build()
}
func getLogger() *zap.Logger {
	return ZapLogger
}
