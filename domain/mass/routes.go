package mass

import "github.com/gin-gonic/gin"

func InitRoutes(r *gin.Engine) {
	massGroup := r.Group("/mass")

	massGroup.POST("/", StoreMassController)
	massGroup.DELETE("/:id", DeleteMassController)
}
