package initialize

import (
	"ecommerce-backend-api/init/global"
	"ecommerce-backend-api/init/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}
