package config

import "github.com/spf13/viper"

func InitConfig() {
	log := CreateLoggers(nil)

	viper.SetConfigType("dotenv")
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
		return
	}
}
