package controllers

import (
	"API-VITALVEST/users/application"
	"API-VITALVEST/users/infraestructure"
	"log"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context){
	id_str := c.Param("user_id")
	id,err:= strconv.Atoi(id_str)
		if err!=nil{
			log.Fatal("no se pudo obtener el id del usuario")
		}

	repo:=infraestructure.NewMysqlRepo()
	UseCase:=application.NewDelete(repo)

		if err = UseCase.Execute(id);err!=nil{
			log.Fatal("no se pudo realizar la conexion a  la bd verifique la sintaxis")
		}
	c.JSON(http.StatusOK,gin.H{"ok":"se elimino el usuario con exito"})
}