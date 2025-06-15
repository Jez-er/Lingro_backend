package logger

import "go.uber.org/zap"

var Logger *zap.Logger

func init() {
	config := zap.NewProductionConfig()
	config.OutputPaths = []string{"./app.log"}
	logger, err := config.Build()
	if err != nil {
		panic(" LOGGER | Failed to initialize Zap logger: " + err.Error())
	}
	Logger = logger
	defer Logger.Sync()
}
