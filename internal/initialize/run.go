package initialize

import (
	"ecommerce-backend-api/init/global"
	"fmt"

	"go.uber.org/zap"
)

func Run() {
	// load configuration
	LoadConfig()
	m := global.Config.Mysql
	fmt.Println("Load configuration mysql", m.Username, m.Password)
	InitLogger()
	global.Logger.Info("Config log ok", zap.String("ok", "success"))
	InitMysql()
	InitRedis()

	r := InitRouter()
	r.Run(":8002")
}
