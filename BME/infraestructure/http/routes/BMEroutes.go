package routes

import (
	"github.com/gin-gonic/gin"
	"API-VITALVEST/BME/infraestructure/dependencies"
)

func RegisterBMEEndpoints(router *gin.Engine) {
	router.POST("/bme", dependencies.NewSaveBMEController().Run)
	router.GET("/bme", dependencies.NewFindAllBMEController().Run)
	router.GET("/bme/:id", dependencies.NewFindByIDBMEController().Run)
	router.PUT("/bme/:id", dependencies.NewUpdateBMEController().Run)
	router.DELETE("/bme/:id", dependencies.NewDeleteBMEController().Run)
}
