package routes

import (
	"github.com/gin-gonic/gin"
	USERS "API-VITALVEST/users/infraestructure/controllers"
)

func UserRoutes(router *gin.Engine){

		routes:= router.Group("/users")
		routes.POST("/",USERS.Create_user)
		router.DELETE("/:id_user",USERS.Delete)
}

