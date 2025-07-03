package routes

import (
	"github.com/gin-gonic/gin"
	"API-VITALVEST/MLX/infraestructure/dependencies"
)

func RegisterMLXEndpoints(router *gin.Engine) {
	router.POST("/mlx", dependencies.NewSaveMLXController().Run)
	router.GET("/mlx", dependencies.NewFindAllMLXController().Run)
	router.GET("/mlx/:id", dependencies.NewFindByIDMLXController().Run)
	router.PUT("/mlx/:id", dependencies.NewUpdateMLXController().Run)
	router.DELETE("/mlx/:id", dependencies.NewDeleteMLXController().Run)
}
