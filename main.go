package main

import (
	users "API-VITALVEST/users/infraestructure/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main(){
	router := gin.Default()

	port :=":8080"
	
	users.UserRoutes(router)

	log.Fatal(router.Run(port))
}