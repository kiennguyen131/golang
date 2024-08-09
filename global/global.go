package global

import (
	"ecommerce-backend-api/init/pkg/logger"
	"ecommerce-backend-api/init/pkg/setting"

	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Mdb    *gorm.DB
)
