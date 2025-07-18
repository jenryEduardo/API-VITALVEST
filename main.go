package main

import (
	"log"

	dependenciesBME "API-VITALVEST/BME/infraestructure/dependencies"
	bme "API-VITALVEST/BME/infraestructure/http/routes"
	dependenciesGSR "API-VITALVEST/GSR/infraestructure/dependencies"
	gsr "API-VITALVEST/GSR/infraestructure/http/routes"
	dependenciesmlx "API-VITALVEST/MLX/infraestructure/dependencies"
	mlx "API-VITALVEST/MLX/infraestructure/http/routes"
	dependenciesMPU "API-VITALVEST/MPU/infraestructure/dependencies"
	mpu "API-VITALVEST/MPU/infraestructure/http/routes"
	dependenciesAlertas "API-VITALVEST/alertas/infraestructure/dependencies"
	alertas "API-VITALVEST/alertas/infraestructure/http/routes"

	// Rutas de sensores
	users "API-VITALVEST/users/infraestructure/routes"

	// WebSocket

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Configurar CORS
	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))

	// Inicializar dependencias de sensores
	dependenciesGSR.InitGSR()
	dependenciesmlx.InitMLX()
	dependenciesBME.InitBME()
	dependenciesMPU.InitMPU()
	dependenciesAlertas.InitAlerta()

	// Registrar rutas de sensores
	gsr.RegisterGSREndpoints(router)
	mlx.RegisterMLXEndpoints(router)
	bme.RegisterBMEEndpoints(router)
	mpu.RegisterMPUEndpoints(router)
	users.UserRoutes(router)
	alertas.RegisterAlertasEndpoints(router)

	// Información del servidor
	port := ":8080"
	log.Fatal(router.Run(port))
}

// // handleLogin maneja la autenticación de usuarios
// func handleLogin(c *gin.Context) {
// 	var loginRequest struct {
// 		Username string `json:"username"`
// 		Password string `json:"password"`
// 	}

// 	if err := c.ShouldBindJSON(&loginRequest); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
// 		return
// 	}

// 	// Usuarios válidos (en producción esto vendría de una base de datos)
// 	validUsers := map[string]string{
// 		"admin":  "admin123",
// 		"juan":   "juan123",
// 		"maria":  "maria123",
// 		"carlos": "carlos123",
// 	}

// 	if password, exists := validUsers[loginRequest.Username]; exists && password == loginRequest.Password {
// 		userData := []map[string]interface{}{
// 			{
// 				"id":       1,
// 				"username": loginRequest.Username,
// 				"name":     loginRequest.Username,
// 				"role":     "user",
// 			},
// 		}
// 		log.Printf("✅ Login exitoso para usuario: %s", loginRequest.Username)
// 		c.JSON(http.StatusOK, userData)
// 	} else {
// 		log.Printf("❌ Login fallido para usuario: %s", loginRequest.Username)
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciales incorrectas"})
// 	}

