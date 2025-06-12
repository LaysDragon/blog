package web

import "github.com/gin-gonic/gin"

func SetupRouter(router *gin.Engine,
	jwt *JwtHandler,
	account *AccountController,
	post *PostController) {
	router.Use(jwt.ParseMiddleware())

	router.POST("/account", account.HandlePost)
	router.POST("/account/login", account.HandleLogin)
	router.GET("/post", post.HandleGet)

	authGroup := router.Group("", RequiredAuthMiddware())
	authGroup.GET("/account/:id", account.HandleGet)

}
