package infraestructure

import (
	"API-VITALVEST/alertas/domain"
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
)

type WSNotifier struct {
	conn *websocket.Conn
}

func NewWSNotifier(wsURL string) *WSNotifier {
	conn, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	if err != nil {
		log.Fatalf("No se pudo conectar al WebSocket externo: %v", err)
	}
	return &WSNotifier{conn: conn}
}

func (w *WSNotifier) EnviarAlerta(alerta domain.Alerta) error {
	data, err := json.Marshal(alerta)
	if err != nil {
		return err
	}
	return w.conn.WriteMessage(websocket.TextMessage, data)
}
