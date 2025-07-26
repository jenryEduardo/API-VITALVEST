package routes

import (
	"github.com/gin-gonic/gin"
	"API-VITALVEST/GSR/infraestructure/dependencies"
	"API-VITALVEST/core/middleware"
)

func RegisterGSREndpoints(router *gin.Engine) {
	gsr := router.Group("/gsr")
	gsr.POST("", dependencies.NewSaveGSRController().Run)
	gsr.GET("", dependencies.NewFindAllGSRController().Run)
	gsr.Use(middleware.AuthMiddleware()) 
	{
		
		
		gsr.GET("/:id", dependencies.NewFindByIDGSRController().Run)
		gsr.PUT("/:id", dependencies.NewUpdateGSRController().Run)		
		gsr.DELETE("/:id", dependencies.NewDeleteGSRController().Run)		
	}
}