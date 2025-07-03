package controllers

import (
	"API-VITALVEST/login/application"
	"API-VITALVEST/login/domain"
	"API-VITALVEST/login/infraestructure"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login_apps(c *gin.Context) {


	var inicio domain.Login

	if err := c.ShouldBindJSON(&inicio);err!=nil{
		c.JSON(http.StatusNotFound,gin.H{"error":"verfique el formato que esta enviando"})
	}

	repo := infraestructure.NewMysqlRepo()
	use_case := application.NewLogin(repo)

	data,err:= use_case.Execute(inicio)
	
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err})
	}

	c.JSON(http.StatusOK,data)
	

}