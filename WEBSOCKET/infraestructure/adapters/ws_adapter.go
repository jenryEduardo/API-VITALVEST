package adapters

import (
	"API-VITALVEST/WEBSOCKET/domain"
	"log"

	"github.com/gorilla/websocket"
)

// WebSocketServer gestiona las conexiones WebSocket.
type WebSocketServer struct {
    Clients    map[*websocket.Conn]bool
    Register   chan *websocket.Conn
    Unregister chan *websocket.Conn
    Broadcast  chan domain.Sensors
}

// NewWebSocketServer crea una nueva instancia de WebSocketServer.
func NewWebSocketServer() *WebSocketServer {
    return &WebSocketServer{
        Clients:    make(map[*websocket.Conn]bool),
        Register:   make(chan *websocket.Conn),
        Unregister: make(chan *websocket.Conn),
        Broadcast:  make(chan domain.Sensors),
    }
}

// Run inicia el servidor WebSocket y maneja la conexión de los clientes.
func (s *WebSocketServer) Run() {
    for {
        select {
        case conn := <-s.Register:
            s.Clients[conn] = true
        case conn := <-s.Unregister:
            delete(s.Clients, conn)
        case data := <-s.Broadcast:
            for client := range s.Clients {
                err := client.WriteJSON(data)
                if err != nil {
                    log.Println("Error al enviar mensaje al cliente:", err)
                    client.Close()
                    delete(s.Clients, client)
                }
            }
        }
    }
}
