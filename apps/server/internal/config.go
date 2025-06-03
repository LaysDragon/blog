package internal

import (
	"log"

	"github.com/spf13/viper"
)

// TODO: add config check
type Config struct {
	DBType         string
	DataSourceName string
	//TODO: first run auto produce with crypto/rand and overwrite config
	JwtSecret string
}

func LoadConfig() Config {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileAlreadyExistsError); ok {
			log.Fatalf("找不到配置檔:%v", err)
		}
		log.Fatalf("fatal error config file: %v", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("unable to decode into struct, %v ", err)
	}
	log.Printf("config: %vF", config)
	return config
}
