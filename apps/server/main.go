package main

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
	prettyconsole "github.com/thessem/zap-prettyconsole"
	"go.uber.org/zap"

	"github.com/LaysDragon/blog/apps/server/db/pgrepo"
	"github.com/LaysDragon/blog/apps/server/internal"
	"github.com/LaysDragon/blog/apps/server/usecase"
	"github.com/LaysDragon/blog/apps/server/web"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
)

func main() {
	config := internal.LoadConfig()
	log := prettyconsole.NewLogger(zap.DebugLevel).Sugar()

	db, err := sql.Open(config.DBType, config.DataSourceName)
	if err != nil {
		log.Fatalf("unable to connect to database, %v", err)
	}

	postRepo := pgrepo.NewPost(db)
	postUseCase := usecase.NewPost(postRepo)
	postController := web.NewPostController(postUseCase)

	gin.Default()
	router := gin.New()
	router.Use(ginzap.Ginzap(log.Desugar(), time.RFC3339, true))
	router.Use(ginzap.RecoveryWithZap(log.Desugar(), true))
	router.GET("/post", postController.HandleGetPost)

	log.Info("Server setup complete!!")
	router.Run()
}
