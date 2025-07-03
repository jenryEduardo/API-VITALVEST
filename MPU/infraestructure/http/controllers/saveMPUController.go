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
	var MPU domain.Mpu

	if err := c.ShouldBindJSON(&MPU); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Todos los campos son requeridos"})
		return
	} 

	err := ctrl.uc.Run(MPU)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusCreated, gin.H{
			"status": true,
			"data": gin.H{
				"type": "MPU",
				"attributes": gin.H{
					"Id del sensor MPU6050": MPU.Id,
					"aceleración X": MPU.Aceleracion_x,
					"Aceleración Y": MPU.Aceleracion_y,
					"Aceleración Z": MPU.Aceleracion_z,
					"Pasos": MPU.Pasos,
					"Nivel de actividad": MPU.Nivel_actividad,
				},
			},
		})
	}
}