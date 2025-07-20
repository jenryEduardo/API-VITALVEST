package routes

import (
	"github.com/gin-gonic/gin"
	"API-VITALVEST/BME/infraestructure/dependencies"
	"API-VITALVEST/core/middleware"
)

func RegisterBMEEndpoints(router *gin.Engine) {
	bme := router.Group("/bme")
	bme.Use(middleware.AuthMiddleware()) // ← Protege todas las rutas BME
	{
		bme.POST("", dependencies.NewSaveBMEController().Run)
		bme.GET("", dependencies.NewFindAllBMEController().Run)
		bme.GET("/:id", dependencies.NewFindByIDBMEController().Run)
		bme.PUT("/:id", dependencies.NewUpdateBMEController().Run)
		bme.DELETE("/:id", dependencies.NewDeleteBMEController().Run)
	}
}
