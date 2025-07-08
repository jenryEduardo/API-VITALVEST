package main

import (
	users "API-VITALVEST/users/infraestructure/routes"
	"log"
	gsr "API-VITALVEST/GSR/infraestructure/http/routes"
	dependenciesGSR "API-VITALVEST/GSR/infraestructure/dependencies"

	mlx "API-VITALVEST/MLX/infraestructure/http/routes"
	dependenciesmlx "API-VITALVEST/MLX/infraestructure/dependencies"

	bme "API-VITALVEST/BME/infraestructure/http/routes"
	dependenciesBME "API-VITALVEST/BME/infraestructure/dependencies"

	mpu "API-VITALVEST/MPU/infraestructure/http/routes"
	dependenciesMPU "API-VITALVEST/MPU/infraestructure/dependencies"

	wsAdapters "API-VITALVEST/WEBSOCKET/infraestructure/adapters"
	wsControllers "API-VITALVEST/WEBSOCKET/infraestructure/controllers"
	wsRoutes "API-VITALVEST/WEBSOCKET/infraestructure/routes"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func main(){
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:    []string{"*"},
		AllowCredentials: true,
	}))

	// GSR
	dependenciesGSR.InitGSR()
	gsr.RegisterGSREndpoints(router)

	//MLX
	dependenciesmlx.InitMLX()
	mlx.RegisterMLXEndpoints(router)

	//BME
	dependenciesBME.InitBME()
	bme.RegisterBMEEndpoints(router)

	//MPU
	dependenciesMPU.InitMPU()
	mpu.RegisterMPUEndpoints(router)

	//USER
	users.UserRoutes(router)

	//WEBSOCKET
	wsSender := wsAdapters.NewWSSender()
	go wsSender.Run()
	wsController := wsControllers.NewWebSocketController(wsSender)
	wsRoutes.RegisterWSEndpoints(router, wsController)

	port :=":8080"
	log.Println("Servidor corriendo en el puerto", port)
	log.Fatal(router.Run(port))
}