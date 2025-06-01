package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/LaysDragon/blog/apps/server/db/pgrepo"
	"github.com/LaysDragon/blog/apps/server/internal"
	"github.com/LaysDragon/blog/apps/server/usecase"
	"github.com/LaysDragon/blog/apps/server/web"
	"github.com/gin-gonic/gin"
)

func main() {
	config := internal.LoadConfig()

	db, err := sql.Open(config.DBType, config.DataSourceName)
	if err != nil {
		log.Fatalf("unable to connect to database, %v", err)
	}

	postRepo := pgrepo.NewPost(db)
	// postService := service.NewPostService()
	postUseCase := usecase.NewPost(postRepo)
	postController := web.NewPostController(postUseCase)

	router := gin.Default()
	router.GET("/post", postController.HandleGetPost)
	router.Run()
}
