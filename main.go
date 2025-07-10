package main

import (
	"log"
	"net/http"

	dependenciesBME "API-VITALVEST/BME/infraestructure/dependencies"
	bme "API-VITALVEST/BME/infraestructure/http/routes"
	dependenciesGSR "API-VITALVEST/GSR/infraestructure/dependencies"
	gsr "API-VITALVEST/GSR/infraestructure/http/routes"
	dependenciesmlx "API-VITALVEST/MLX/infraestructure/dependencies"
	mlx "API-VITALVEST/MLX/infraestructure/http/routes"
	dependenciesMPU "API-VITALVEST/MPU/infraestructure/dependencies"
	mpu "API-VITALVEST/MPU/infraestructure/http/routes"
	// Rutas de sensores
	users "API-VITALVEST/users/infraestructure/routes"

	// WebSocket
	wsAdapters "API-VITALVEST/WEBSOCKET/infraestructure/adapters"
	wsControllers "API-VITALVEST/WEBSOCKET/infraestructure/controllers"

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

	// Registrar rutas de sensores
	gsr.RegisterGSREndpoints(router)
	mlx.RegisterMLXEndpoints(router)
	bme.RegisterBMEEndpoints(router)
	mpu.RegisterMPUEndpoints(router)
	users.UserRoutes(router)

	// Configurar WebSocket
	wsServer := wsAdapters.NewWebSocketServer()
	go wsServer.Run()
	wsController := wsControllers.NewWebSocketController(wsServer)

	// Rutas WebSocket
	router.GET("/ws", wsController.HandleWebSocket)
	router.POST("/sendData", wsController.HandleSendData)
	router.GET("/ws-status", wsController.HandleStatus)

	// Ruta de login
	router.POST("/login", handleLogin)

	// Información del servidor
	port := ":8080"
	log.Println("🚀 Servidor VitalVest iniciado")
	log.Println("📡 WebSocket disponible en ws://localhost:8080/ws")
	log.Println("📤 Endpoint sendData: http://localhost:8080/sendData")
	log.Println("📊 Status WebSocket: http://localhost:8080/ws-status")
	log.Println("🔐 Login: http://localhost:8080/login")
	log.Printf("🌐 Servidor corriendo en http://localhost%s", port)

	log.Fatal(router.Run(port))
}

// handleLogin maneja la autenticación de usuarios
func handleLogin(c *gin.Context) {
	var loginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	// Usuarios válidos (en producción esto vendría de una base de datos)
	validUsers := map[string]string{
		"admin":  "admin123",
		"juan":   "juan123",
		"maria":  "maria123",
		"carlos": "carlos123",
	}

	if password, exists := validUsers[loginRequest.Username]; exists && password == loginRequest.Password {
		userData := []map[string]interface{}{
			{
				"id":       1,
				"username": loginRequest.Username,
				"name":     loginRequest.Username,
				"role":     "user",
			},
		}
		log.Printf("✅ Login exitoso para usuario: %s", loginRequest.Username)
		c.JSON(http.StatusOK, userData)
	} else {
		log.Printf("❌ Login fallido para usuario: %s", loginRequest.Username)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciales incorrectas"})
	}
}
