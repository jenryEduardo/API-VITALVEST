package controllers

import (
	"API-VITALVEST/login/application"
	"API-VITALVEST/login/domain"
	"API-VITALVEST/login/infraestructure"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login_apps(c *gin.Context) {
	var inicio domain.Login

	if err := c.ShouldBindJSON(&inicio); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Verifique el formato que está enviando"})
		return
	}

	repo := infraestructure.NewMysqlRepo()
	use_case := application.NewLogin(repo)

	data, err := use_case.Execute(inicio)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciales incorrectas"})
		return
	}

	if len(data) == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no encontrado"})
		return
	}

	c.JSON(http.StatusOK, data)
}
