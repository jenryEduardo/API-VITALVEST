package controllers

import (
	"API-VITALVEST/users/application"
	"API-VITALVEST/users/infraestructure"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {

	repo := infraestructure.NewMysqlRepo()
	use_case := application.NewGetUsers(repo)

	data,err := use_case.Execute()

		if err!=nil{
			c.JSON(http.StatusBadGateway,gin.H{"error":"verifique la solicitud a la BD"})
		}

	c.JSON(http.StatusOK,data)
}