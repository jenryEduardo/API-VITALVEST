package middleware

import (
	"API-VITALVEST/core/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Obtener el header Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token no proporcionado"})
			c.Abort()
			return
		}

		// Validar formato Bearer <token>
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Formato inválido del token"})
			c.Abort()
			return
		}

		token := parts[1]

		// Validar token
		claims, err := jwtservice.ValidateJWT(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido o expirado"})
			c.Abort()
			return
		}

		// Guardar claims en contexto
		c.Set("usuario_id", claims.UsuarioID)
		c.Set("email", claims.Email)

		c.Next()
	}
}
