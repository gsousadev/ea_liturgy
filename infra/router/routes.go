package router

import (
	"github.com/gsousadev/ea_liturgy/domain/mass"

	"github.com/gin-gonic/gin"
)

func initRoutes(r *gin.Engine) {
	mass.InitRoutes(r)
}
