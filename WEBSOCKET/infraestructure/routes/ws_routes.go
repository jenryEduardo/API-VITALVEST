package routes

import (
	"API-VITALVEST/WEBSOCKET/domain"
	"API-VITALVEST/WEBSOCKET/infraestructure/controllers"
	"encoding/json"
	"net/http"
)

// Definir la ruta que recibirá el iddata
func InitializeRoutes(wsController *controllers.WebSocketController) {
    http.HandleFunc("/ws", wsController.HandleWebSocket)

    // Ruta para recibir el iddata desde la API y enviarlo a WebSocket
    http.HandleFunc("/sendData", func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
            return
        }

        var data domain.Sensors
        err := json.NewDecoder(r.Body).Decode(&data)
        if err != nil {
            http.Error(w, "Error al leer el cuerpo de la solicitud", http.StatusBadRequest)
            return
        }

        // Llamar al servicio WebSocket para enviar el data
        if err := wsController.SendData(data); err != nil {
            http.Error(w, "Error al enviar el data", http.StatusInternalServerError)
            return
        }

        w.WriteHeader(http.StatusOK)
        w.Write([]byte("data enviado exitosamente"))
    })
}
