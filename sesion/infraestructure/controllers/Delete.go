package controllers

import (
	"API-VITALVEST/sesion/application"
	"API-VITALVEST/sesion/infraestructure"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteSession(c *gin.Context){


	id_params := c.Param("id_session")

	id,err := strconv.Atoi(id_params)

		if err!=nil{
			c.JSON(http.StatusNotFound,gin.H{"error":"no se pudo obtener el id"})
		}

	repo := infraestructure.NewMysqlRepo()
	use_case := application.NewDelete(repo)

		if err:=use_case.Execute(id);err!=nil{
			c.JSON(http.StatusInternalServerError,gin.H{"error":"ocurrio un error al ejecutar la solicitud en la BD"})
		}

		c.JSON(http.StatusContinue,gin.H{"succesfull":"se elimino el usuario correctamente"})

}