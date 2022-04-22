package router

import "github.com/gin-gonic/gin"

func InitServer() {
	r := gin.Default()

	initRoutes(r)

	r.Run()
}
