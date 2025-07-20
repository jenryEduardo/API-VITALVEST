package routes

import (
	"github.com/gin-gonic/gin"
	USERS "API-VITALVEST/users/infraestructure/controllers"
	"API-VITALVEST/core/middleware"
)

func UserRoutes(router *gin.Engine) {
	// Ruta pública
	router.POST("/login", USERS.LoginController)

	// Rutas protegidas
	protected := router.Group("/users")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.POST("/", USERS.Create_user)
		protected.GET("/", USERS.GetUsers)
		protected.PUT("/:user_id", USERS.UpdateUser)
		protected.DELETE("/:user_id", USERS.Delete)
		protected.GET("/perfil", func(c *gin.Context) {
			usuarioID, _ := c.Get("usuario_id")
			c.JSON(200, gin.H{"usuario_id": usuarioID})
		})
	}
}
