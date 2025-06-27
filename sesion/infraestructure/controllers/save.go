package controllers

import (
	"API-VITALVEST/sesion/application"
	"API-VITALVEST/sesion/domain"
	"API-VITALVEST/sesion/infraestructure"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Save_session(c *gin.Context){

	var session domain.Sesion

	if err:=c.ShouldBindJSON(session);err!=nil{
		fmt.Println("no se pudo deserializar el json")
	}

	repo:=infraestructure.NewMysqlRepo()
	useCase:= application.NewSaveUser(repo)

	if err:=useCase.Execute(session);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":"verifique sus datos"})
	}

	c.JSON(http.StatusCreated,gin.H{"ok":"se creo la session"})
}