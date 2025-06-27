package controllers

import (
	"API-VITALVEST/sesion/application"
	"API-VITALVEST/sesion/domain"
	"API-VITALVEST/sesion/infraestructure"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdateSession(c *gin.Context){

	id_params := c.Param("id_Session")
	id,err := strconv.Atoi(id_params)

	if err!=nil{
		c.JSON(http.StatusNotFound,gin.H{"error":"no se pudo encontrar el id en el header"})
	}

	var Session domain.Sesion

	if err:=c.ShouldBindJSON(&Session);err!=nil{
		c.JSON(http.StatusNotFound,gin.H{"error":"no se pudo encontrar la sesion en el header"})
	}

	repo := infraestructure.NewMysqlRepo()
	use_case:=application.NewUpdateSesion(repo)

	if err:= use_case.Execute(Session,id);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":"verifique los datos de la solicitud"})
	}
}