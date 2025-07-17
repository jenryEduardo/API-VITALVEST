package app

import (
	"API-VITALVEST/WEBSOCKET/domain"
	"API-VITALVEST/WEBSOCKET/infraestructure/controllers"
)

type SensorService struct {
	WebSocketController *controllers.WebSocketController
}

func NewSensorService(wsController *controllers.WebSocketController) *SensorService {
	return &SensorService{WebSocketController: wsController}
}

func (s *SensorService) EnviarDatos(data domain.Sensors) error {
	return s.WebSocketController.SendData(data) 
}
