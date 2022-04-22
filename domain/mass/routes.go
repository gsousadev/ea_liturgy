package mass

import "github.com/gin-gonic/gin"

func InitRoutes(r *gin.Engine) {
	r.POST("/mass", StoreMassController)
}
