package controllers

import (
	jwtservice "API-VITALVEST/core/jwt"
	"API-VITALVEST/users/application"
	"API-VITALVEST/users/domain"
	"API-VITALVEST/users/infraestructure"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func LoginController(c *gin.Context) {
	var req domain.User

	// Validar JSON y campos obligatorios
	if err := c.ShouldBindJSON(&req); err != nil || req.UserName == "" || req.Passwords == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Usuario y contraseña requeridos"})
		return
	}

	repo := infraestructure.NewMysqlRepo()
	uc := application.NewLogin(repo)

	// Verifica username y password
	user, err := uc.Execute(req.UserName, req.Passwords)
	if err != nil {
		// Diferenciar entre error de validación y credenciales incorrectas
		if strings.Contains(err.Error(), "caracteres") {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciales inválidas"})
		return
	}

	// Generar JWT
	token, err := jwtservice.GenerateJWT(user.Id, user.UserName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo generar el token"})
		return
	}

	// Respuesta exitosa (no enviar la contraseña)
	user.Passwords = ""
	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"user":  user,
	})
}