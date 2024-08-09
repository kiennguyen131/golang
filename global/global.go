package global

import (
	"ecommerce-backend-api/init/pkg/logger"
	"ecommerce-backend-api/init/pkg/setting"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
)
