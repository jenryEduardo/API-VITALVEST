package main

import (
	users "API-VITALVEST/users/infraestructure/routes"
	"log"
	gsr "API-VITALVEST/GSR/infraestructure/http/routes"
	dependenciesGSR "API-VITALVEST/GSR/infraestructure/dependencies"

	mlx "API-VITALVEST/MLX/infraestructure/http/routes"
	dependenciesmlx "API-VITALVEST/MLX/infraestructure/dependencies"
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

	users.UserRoutes(router)

	port :=":8080"
	log.Println("Servidor corriendo en el puerto", port)
	log.Fatal(router.Run(port))
}