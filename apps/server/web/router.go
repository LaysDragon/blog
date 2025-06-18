package web

import "github.com/gin-gonic/gin"

func SetupRouter(router *gin.Engine,
	jwt *JwtHandler,
	account *AccountController,
	site *SiteController,
	post *PostController) {
	router.Use(jwt.ParseMiddleware())

	router.POST("/accounts", account.HandlePost)
	router.POST("/accounts/login", account.HandleLogin)

	router.GET("/posts", post.HandleList)
	router.GET("/posts/:id", post.HandleGet)

	authGroup := router.Group("", RequiredAuthMiddware())
	authGroup.GET("/accounts/:id", account.HandleGet)
	authGroup.GET("/accounts", account.HandleList)

	authGroup.GET("/sites", site.HandleList)

	authGroup.POST("/posts", post.HandlePost)

}
