package routes

import (
	"github.com/gin-gonic/gin"
	"API-VITALVEST/MPU/infraestructure/dependencies"
)

func RegisterMPUEndpoints(router *gin.Engine) {
	router.POST("/mpu", dependencies.NewSaveMPUController().Run)
	router.GET("/mpu", dependencies.NewFindAllMPUController().Run)
	router.GET("/mpu/:id", dependencies.NewFindByIDMPUController().Run)
	router.PUT("/mpu/:id", dependencies.NewUpdateMPUController().Run)
	router.DELETE("/mpu/:id", dependencies.NewDeleteMPUController().Run)
}
