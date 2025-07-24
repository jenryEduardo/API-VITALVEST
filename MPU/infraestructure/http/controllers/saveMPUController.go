package controllers

import (
	"API-VITALVEST/MPU/application"
	"API-VITALVEST/MPU/domain"
	"net/http"
	"github.com/gin-gonic/gin"
)

type SaveMPUController struct {
	uc *application.SaveMPU_UC
}

func NewSaveMPUController(uc *application.SaveMPU_UC) *SaveMPUController {
	return &SaveMPUController{uc: uc}
}

func (ctrl *SaveMPUController) Run(c *gin.Context) {

		var data  domain.Mpu 
	

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos o incompletos"})
		return
	}

	// // Calcular pasos y nivel con aceleracion recibida
	// pasos, nivel := control.ConvertirDatosEnPasos(input.Mpu6050.Aceleracion.X, input.Mpu6050.Aceleracion.Y, input.Mpu6050.Aceleracion.Z)

	// // Crear objeto para guardar
	// mpu := domain.Mpu{
	// 	Mpu6050:       input.Mpu6050,
	// 	Pasos:         pasos,
	// 	NivelActividad: nivel,
	// }

	err := ctrl.uc.Run(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": true,
		"data": gin.H{
			"type": "MPU",
			"attributes": gin.H{
				"pasos":          data.Pasos,
			},
		},
	})
}
