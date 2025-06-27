package controllers

import (
	"API-VITALVEST/sesion/application"
	"API-VITALVEST/sesion/infraestructure"
	"net/http"
	"github.com/gin-gonic/gin"
)

func GetAllSession(c *gin.Context){

	repo := infraestructure.NewMysqlRepo()
	use_case := application.NewGetAllUser(repo)

	
	datos,err:=use_case.Execute()

	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":"no se pudo obtener los datos de la BD"})
	}

	c.JSON(http.StatusOK,datos)

}