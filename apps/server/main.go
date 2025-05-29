package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"

	"github.com/LaysDragon/blog/apps/server/db/pgrepo"
	"github.com/LaysDragon/blog/apps/server/service"
	"github.com/LaysDragon/blog/apps/server/usecase"
	"github.com/LaysDragon/blog/apps/server/web"
	"github.com/gin-gonic/gin"
)

// TODO: add config check
type Config struct {
	DBType         string
	DataSourceName string
}

func main() {
	fmt.Println("hello world")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileAlreadyExistsError); ok {
			log.Fatalf("找不到配置檔:%v", err)
		}
		log.Fatalf("fatal error config file: %v", err)
	}

	var config Config
	fmt.Println(viper.AllSettings())
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("unable to decode into struct, %v ", err)
	}
	log.Printf("config: %v", config)

	db, err := sql.Open(config.DBType, config.DataSourceName)
	if err != nil {
		log.Fatalf("unable to connect to database, %v", err)
	}

	postRepo := pgrepo.NewPostRepo(db)
	postService := service.NewPostService()
	postUseCase := usecase.NewPostUseCase(postRepo, postService)
	postController := web.NewPostController(postUseCase)

	router := gin.Default()
	router.GET("/post", postController.HandleGetPost)
	router.Run()
}
