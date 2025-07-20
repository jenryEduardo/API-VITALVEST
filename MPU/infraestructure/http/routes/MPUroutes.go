package routes

import (
	"github.com/gin-gonic/gin"
	"API-VITALVEST/MPU/infraestructure/dependencies"
	"API-VITALVEST/core/middleware"
)

func RegisterMPUEndpoints(router *gin.Engine) {
	mpu := router.Group("/mpu")
	mpu.Use(middleware.AuthMiddleware())
	{
		mpu.POST("", dependencies.NewSaveMPUController().Run)
		mpu.GET("", dependencies.NewFindAllMPUController().Run)
		mpu.GET("/:id", dependencies.NewFindByIDMPUController().Run)
		mpu.PUT("/:id", dependencies.NewUpdateMPUController().Run)
		mpu.DELETE("/:id", dependencies.NewDeleteMPUController().Run)
	}
}
