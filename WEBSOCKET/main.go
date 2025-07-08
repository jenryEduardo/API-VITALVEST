package main

import (
	"API-VITALVEST/WEBSOCKET/infraestructure/adapters"
	"API-VITALVEST/WEBSOCKET/infraestructure/controllers"
	"API-VITALVEST/WEBSOCKET/infraestructure/routes"
	"log"
	"net/http"
)

func main() {
    // Configurar WebSocket Server
    wsServer := adapters.NewWebSocketServer()
    go wsServer.Run()

    // Crear el controlador WebSocket
    wsController := controllers.NewWebSocketController(wsServer)

    // Inicializar las rutas
    routes.InitializeRoutes(wsController)

    log.Println("Servidor WebSocket iniciado en :3010")
    log.Fatal(http.ListenAndServe(":3010", nil))
}
