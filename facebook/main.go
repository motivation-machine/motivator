package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
)

func main() {

	// Load configuration
	viper.SetConfigFile("config.json")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Failed to load configuration %v", err))
	}
	viper.AutomaticEnv()

	// DB connection
	db, err := gorm.Open("postgres",
		fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
			viper.Get("FACEBOOK_DB_USER"),
			viper.Get("FACEBOOK_DB_PASSWORD"),
			viper.Get("FACEBOOK_DB_HOST"),
			viper.Get("FACEBOOK_DB_NAME")))
	if err != nil {
		panic(fmt.Errorf("Failed to connect to database %v", err))
	}
	defer db.Close()

	// Get Facebook page sources
}
