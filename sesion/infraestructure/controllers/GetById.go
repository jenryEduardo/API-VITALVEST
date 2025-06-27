package controllers

import (
	"API-VITALVEST/sesion/application"
	"API-VITALVEST/sesion/infraestructure"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetByIdSession(c *gin.Context) {


	id_params := c.Param("id_session")
	id,err := strconv.Atoi(id_params)

		if err!=nil{
			c.JSON(http.StatusNotFound,gin.H{"error":"no se pudo encontrar nada en la solicitud"})
		}

	repo := infraestructure.NewMysqlRepo()
	use_case := application.NewGetUser(repo)

	session,err:= use_case.Execute(id)
		if err!=nil{
			c.JSON(http.StatusBadGateway,gin.H{"error":"verifique los datos que envio o su conexion a la BD"})
		}

		c.JSON(http.StatusOK,session)
}	