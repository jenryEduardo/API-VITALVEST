package main

import (
	"log"
	"github.com/gin-gonic/gin"
)

func main(){
	router := gin.Default()

	port :=":8080"

	log.Fatal(router.Run(port))
}