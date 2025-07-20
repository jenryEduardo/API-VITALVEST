package controllers

import (
	jwtservice "API-VITALVEST/core/jwt"
	"API-VITALVEST/users/application"
	"API-VITALVEST/users/domain"
	"API-VITALVEST/users/infraestructure"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginController(c *gin.Context) {
	var req domain.User
	if err := c.ShouldBindJSON(&req); err != nil || req.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nombre requerido"})
		return
	}

	repo := infraestructure.NewMysqlRepo()
	uc := application.NewLogin(repo)

	user, err := uc.Execute(req.Name)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no encontrado"})
		return
	}

	token, err := jwtservice.GenerateJWT(user.Id, user.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo generar el token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"user":  user,
	})
}
