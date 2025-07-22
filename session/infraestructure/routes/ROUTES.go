package routes

import (
	"github.com/gin-gonic/gin"
	"API-VITALVEST/session/infraestructure/controllers"
)

func SetUproutesSession(c *gin.Engine){

	router := c.Group("/session")

	router.POST("/",controllers.SaveSession)
	router.GET("/",controllers.GetAllSessions)
}