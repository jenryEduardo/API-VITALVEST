package controllers

import (
	"API-VITALVEST/WEBSOCKET/domain"
	"API-VITALVEST/WEBSOCKET/infraestructure/adapters"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// WebSocketController maneja las conexiones WebSocket.
type WebSocketController struct {
    WebSocketServer *adapters.WebSocketServer
}

// NewWebSocketController crea un nuevo controlador WebSocket.
func NewWebSocketController(wsServer *adapters.WebSocketServer) *WebSocketController {
    return &WebSocketController{WebSocketServer: wsServer}
}

// HandleWebSocket maneja la conexión WebSocket con el cliente.
func (c *WebSocketController) HandleWebSocket(w http.ResponseWriter, r *http.Request) {
    upgrader := websocket.Upgrader{
        CheckOrigin: func(r *http.Request) bool { return true },
    }
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println("Error al establecer WebSocket:", err)
        return
    }
    defer conn.Close()

    c.WebSocketServer.Register <- conn
    for {
        _, _, err := conn.ReadMessage()
        if err != nil {
            c.WebSocketServer.Unregister <- conn
            break
        }
    }
}

// SendPedido envía un pedido a través del WebSocket.
func (c *WebSocketController) SendData(data domain.Sensors) error {
    c.WebSocketServer.Broadcast <- data
    return nil
}
