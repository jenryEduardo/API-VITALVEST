package controllers

import (
	"API-VITALVEST/session/application"
	"API-VITALVEST/session/infraestructure"
	"net/http"
	"github.com/gin-gonic/gin"
)

func GetAllSessions(c *gin.Context){


	repo := infraestructure.NewMysqlRepo()
	useCase := application.NewGetAll(repo)

	 data,err:=useCase.Execute()

	 if err!=nil{
		c.JSON(http.StatusNotFound,gin.H{"err":"no se encontro ningun dato"})
	 }

	 c.JSON(http.StatusOK,data)

}