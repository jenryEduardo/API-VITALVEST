package controllers

import (
	"API-VITALVEST/BME/application"
	"API-VITALVEST/BME/domain"
	"API-VITALVEST/core/workerpool"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type SaveBMEController struct {
	uc   *application.SaveBME_UC
	pool *workerpool.WorkerPool
}

func NewSaveBMEController(uc *application.SaveBME_UC, pool *workerpool.WorkerPool) *SaveBMEController {
	return &SaveBMEController{uc: uc, pool: pool}
}

func (ctrl *SaveBMEController) Run(c *gin.Context) {
	var BME domain.Bme

	if err := c.ShouldBindJSON(&BME); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Todos los campos son requeridos"})
		return
	}

	// Enviar al pool para evitar saturar la DB
	resultChan := ctrl.pool.Submit(func() (interface{}, error) {
		err := ctrl.uc.Run(BME)
		if err != nil {
			return nil, err
		}
		return BME, nil
	})

	// Esperar resultado con timeout
	select {
	case result := <-resultChan:
		if result.Err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Err.Error()})
			return
		}

		savedBME := result.Data.(domain.Bme)
		c.JSON(http.StatusCreated, gin.H{
			"status": true,
			"data": gin.H{
				"type": "BME",
				"attributes": gin.H{
					"Id del sensor BME280":          savedBME.Id,
					"Temperatura ambiente":          savedBME.Temperatura,
					"Humedad relativa (BME280)":     savedBME.Humedad,
					"Presión":                       savedBME.Presion,
				},
			},
		})

	case <-time.After(5 * time.Second):
		c.JSON(http.StatusGatewayTimeout, gin.H{
			"error": "timeout al guardar datos del sensor",
		})
	}
}