package main

import (
	users "API-VITALVEST/users/infraestructure/routes"
	login "API-VITALVEST/login/infraestructure/routes"
	//sesiones "API-VITALVEST/sesion/infraestructure"
	"log"

	"github.com/gin-gonic/gin"
)

func main(){
	router := gin.Default()

	port :=":8080"
	
	users.UserRoutes(router)
	login.SetUpRoutes(router)

	log.Fatal(router.Run(port))
}