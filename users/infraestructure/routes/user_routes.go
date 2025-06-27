package routes

import (
	"github.com/gin-gonic/gin"
	USERS "API-VITALVEST/users/infraestructure/controllers"
)

func UserRoutes(router *gin.Engine){

		routes:= router.Group("/users")
		routes.POST("/",USERS.Create_user)
		routes.DELETE("/:user_id",USERS.Delete)
		routes.PUT("/:user_id",USERS.UpdateUser)
		routes.GET("/",USERS.GetUsers)
}

