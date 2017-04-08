package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("config.json")
	err := viper.ReadInConfig()
	if err != nil {
		panic("Failed to read config file")
	}
	fmt.Printf("%s", viper.GetString("serviceName"))
}
