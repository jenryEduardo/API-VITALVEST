package routes

import (
	"github.com/gin-gonic/gin"
	"API-VITALVEST/login/infraestructure/controllers"
)

func SetUpRoutes(routes *gin.Engine){

router := routes.Group("/login")

router.POST("/",controllers.Login_apps)

}