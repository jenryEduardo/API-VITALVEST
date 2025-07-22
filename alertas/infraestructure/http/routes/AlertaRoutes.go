package routes

import (
	"API-VITALVEST/alertas/infraestructure/dependencies"
	"API-VITALVEST/core/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterAlertasEndpoints(router *gin.Engine) {
	alertas := router.Group("/alertas")
	alertas.Use(middleware.AuthMiddleware())
	{
		alertas.POST("", dependencies.NewSaveAlertaController().Run)
		alertas.GET("", dependencies.NewFindAllAlertaController().Run)
		alertas.GET("/:id", dependencies.NewFindByIDAlertaController().Run)
		alertas.PUT("/:id", dependencies.NewUpdateAlertaController().Run)
		alertas.DELETE("/:id", dependencies.NewDeleteAlertaController().Run)
}
}
