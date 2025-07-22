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

	datos,err := use_case.Execute()

		if err!=nil{
			c.JSON(http.StatusBadGateway,gin.H{"error":err})
		}

	c.JSON(http.StatusOK,datos)
}