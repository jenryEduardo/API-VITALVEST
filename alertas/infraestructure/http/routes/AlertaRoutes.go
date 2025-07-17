package routes

import (
	"github.com/gin-gonic/gin"
	"API-VITALVEST/alertas/infraestructure/dependencies"
)

func RegisterAlertasEndpoints(router *gin.Engine) {
	router.POST("/alertas", dependencies.NewSaveAlertaController().Run)
	router.GET("/alertas", dependencies.NewFindAllAlertaController().Run)
	router.GET("/alertas/:id", dependencies.NewFindByIDAlertaController().Run)
	router.PUT("/alertas/:id", dependencies.NewUpdateAlertaController().Run)
	router.DELETE("/alertas/:id", dependencies.NewDeleteAlertaController().Run)
}
