package main

import (
	// login "API-VITALVEST/login/infraestructure/routes"
	//sesiones "API-VITALVEST/sesion/infraestructure"
	"log"
	"API-VITALVEST/core/workerpool"
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

	session "API-VITALVEST/session/infraestructure/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	// workerpool
	pool := workerpool.New(10)

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
	dependenciesBME.InitBME(pool)
	dependenciesMPU.InitMPU()
	dependenciesAlertas.InitAlerta()

	// Registrar rutas de sensores
	gsr.RegisterGSREndpoints(router)
	mlx.RegisterMLXEndpoints(router)
	bme.RegisterBMEEndpoints(router)
	mpu.RegisterMPUEndpoints(router)
	users.UserRoutes(router)

	// login.SetUpRoutes(router)
	session.SetUproutesSession(router)
	alertas.RegisterAlertasEndpoints(router)

	// Información del servidor
	port := ":8080"
	log.Println("🚀 Servidor VitalVest iniciado")
	log.Printf("🌐 Servidor corriendo en http://localhost%s", port)
	log.Fatal(router.Run(port))
}