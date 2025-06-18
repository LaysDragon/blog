package config

import (
	"log"

	"github.com/LaysDragon/blog/apps/server/db"
	"github.com/LaysDragon/blog/apps/server/perm"
	"github.com/LaysDragon/blog/apps/server/web"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

var Module = fx.Module("config",
	fx.Provide(
		fx.Annotate(
			LoadConfig,
			fx.As(new(db.DbConfig)),
			fx.As(new(perm.PermConfig)),
			fx.As(new(web.WebConfig)),
		),
	),
)

// TODO: add config check
type Config struct {
	DBType         string
	DataSourceName string
	//TODO: first run auto produce with crypto/rand and overwrite config
	JwtSecret string
}

func (c Config) GetDBType() string {
	return c.DBType
}
func (c Config) GetDataSourceName() string {
	return c.DataSourceName
}

func (c Config) GetJwtSecret() string {
	return c.JwtSecret
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
