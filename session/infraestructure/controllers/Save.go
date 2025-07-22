package controllers

import (
	"API-VITALVEST/session/domain"
	"net/http"
	"API-VITALVEST/session/infraestructure"
	"API-VITALVEST/session/application"
	"github.com/gin-gonic/gin"
)

func SaveSession(c *gin.Context){

	var data domain.Session

	err:=c.ShouldBindJSON(&data)

	if err!=nil{
		c.JSON(http.StatusNotFound,gin.H{"error":"no se encontro los datos correctos en el json"})
		return
	}

	repo := infraestructure.NewMysqlRepo()
	useCase:=application.NewSession(repo)

	if err:= useCase.Execute(data);err!=nil{
		c.JSON(http.StatusOK,gin.H{"ok":"se creo la sesion con exito"})
		return
	}

}