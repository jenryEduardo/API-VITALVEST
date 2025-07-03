package controllers

import (
	"API-VITALVEST/BME/application"
	"API-VITALVEST/BME/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SaveBMEController struct {
	uc *application.SaveBME_UC
}

func NewSaveBMEController(uc *application.SaveBME_UC) *SaveBMEController {
	return &SaveBMEController{uc: uc}
}

func (ctrl *SaveBMEController) Run(c *gin.Context) {
	var BME domain.Bme

	if err := c.ShouldBindJSON(&BME); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Todos los campos son requeridos"})
		return
	} 

	err := ctrl.uc.Run(BME)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusCreated, gin.H{
			"status": true,
			"data": gin.H{
				"type": "BME",
				"attributes": gin.H{
					"Id del sensor BME280": BME.Id,
					"Temperatura ambiente": BME.Temperatura_ambiente,
					"Humedad relativa (BME280)":BME.Humedad_relativa,
				},
			},
		})
	}
}