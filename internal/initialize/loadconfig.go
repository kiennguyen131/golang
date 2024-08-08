package initialize

import (
	"ecommerce-backend-api/init/global"
	"fmt"

	"github.com/spf13/viper"
)

func LoadConfig() {
	viper := viper.New()
	viper.AddConfigPath("./config/") //path to config
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	// read configuration
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Failed to read configuration %w \n", err))
	}

	// read server configuration
	fmt.Println("Server Port::", viper.GetInt("server.port"))
	fmt.Println("Server JWT Key::", viper.GetString("security.jwt.key"))

	// configure structure
	if err := viper.Unmarshal(&global.Config); err != nil {
		fmt.Printf("Unable to decode configuration %v", err)
	}
}
