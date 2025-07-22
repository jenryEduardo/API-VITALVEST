package controllers

import (
	"API-VITALVEST/users/application"
	"API-VITALVEST/users/domain"
	"API-VITALVEST/users/infraestructure"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdateUser(c *gin.Context) {

	var user domain.User

	id_string := c.Param("user_id")
	id,err:=strconv.Atoi(id_string)
		if err!=nil{
			c.JSON(http.StatusForbidden,gin.H{"error":"No se encontro ningun id"})
		}

	if err :=c.ShouldBindJSON(&user);err!=nil{
		c.JSON(http.StatusNotFound,gin.H{"error":"no se encontro o envio datos nulos"})
	}

	repo :=infraestructure.NewMysqlRepo()
	use_case :=application.NewUpdate(repo)

		if err:= use_case.Execute(user,id);err!=nil{
			c.JSON(http.StatusBadRequest,gin.H{"eeror":"no se pudo actualizar el user"})
		}
	
		c.JSON(http.StatusOK,gin.H{"ok":"se actualizo el usuario correctamente"})

}