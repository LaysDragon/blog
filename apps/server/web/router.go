package web

import "github.com/gin-gonic/gin"

func SetupRouter(router *gin.Engine,
	account *AccountController,
	post *PostController) {
	router.GET("/post", post.HandleGet)

	router.POST("/account", account.HandlePost)
	router.GET("/account/:id", account.HandleGet)
	router.POST("/account/login", account.HandleLogin)

}
