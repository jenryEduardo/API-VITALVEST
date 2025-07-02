package routes

import (
	"github.com/gin-gonic/gin"
	"API-VITALVEST/GSR/infraestructure/dependencies"
)

func RegisterGSREndpoints(router *gin.Engine) {
	router.POST("/gsr", dependencies.NewSaveGSRController().Run)
	router.GET("/gsr", dependencies.NewFindAllGSRController().Run)
	router.GET("/gsr/:id", dependencies.NewFindByIDGSRController().Run)
	router.PUT("/gsr/:id", dependencies.NewUpdateGSRController().Run)
	router.DELETE("/gsr/:id", dependencies.NewDeleteGSRController().Run)
}
