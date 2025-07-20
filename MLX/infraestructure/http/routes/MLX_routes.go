package routes

import (
	"github.com/gin-gonic/gin"
	"API-VITALVEST/MLX/infraestructure/dependencies"
	"API-VITALVEST/core/middleware"
)

func RegisterMLXEndpoints(router *gin.Engine) {
	mlx := router.Group("/mlx")
	mlx.Use(middleware.AuthMiddleware())
	{
		mlx.POST("", dependencies.NewSaveMLXController().Run)
		mlx.GET("", dependencies.NewFindAllMLXController().Run)
		mlx.GET("/:id", dependencies.NewFindByIDMLXController().Run)
		mlx.PUT("/:id", dependencies.NewUpdateMLXController().Run)
		mlx.DELETE("/:id", dependencies.NewDeleteMLXController().Run)
	}
}
