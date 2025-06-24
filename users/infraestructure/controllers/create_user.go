package controllers

import (
	"API-VITALVEST/users/application"
	"API-VITALVEST/users/domain"
	"API-VITALVEST/users/infraestructure"
	"fmt"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
)

func Create_user(r *gin.Context) {

	var User domain.User

	if err := r.ShouldBindJSON(&User);err!=nil{
		log.Println("error",err)
		r.JSON(http.StatusBadRequest,gin.H{"error":"no se pudo crrear el usuario"})
	}

	repo:=infraestructure.NewMysqlRepo()
	UseCase := application.NEWUSER(repo)

	if err:=UseCase.Execute(User);err!=nil{
		fmt.Println("error al guardar los datos",err)
	}

	r.JSON(http.StatusCreated,gin.H{"ok":"ya se creo correctamente el usuario"})

}