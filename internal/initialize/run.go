package initialize

import (
	"ecommerce-backend-api/init/global"
	"fmt"
)

func Run() {
	// load configuration
	LoadConfig()
	m := global.Config.Mysql
	fmt.Println("Load configuration mysql", m.Username, m.Password)
	InitLogger()
	InitMysql()
	InitRedis()

	r := InitRouter()
	r.Run(":8002")
}
