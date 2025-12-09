package controllers

import (
	"API-VITALVEST/GSR/application"
	"API-VITALVEST/GSR/domain"
	"net/http"
	"API-VITALVEST/core/workerpool"
	"github.com/gin-gonic/gin"
)

type SaveGSRController struct {
	uc *application.SaveGSR_UC
	pool *workerpool.WorkerPool
}

func NewSaveGSRController(uc *application.SaveGSR_UC, pool *workerpool.WorkerPool) *SaveGSRController {
	return &SaveGSRController{uc: uc, pool: pool}
}

func (ctrl *SaveGSRController) Run(c *gin.Context) {
	var gsr domain.Gsr
		if err := c.ShouldBindJSON(&gsr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Todos los campos son requeridos"})
		return
	}

	// Enviar al pool para evitar saturar la DB
	resultChan := ctrl.pool.Submit(func() (interface{}, error) {
		err := ctrl.uc.Run(gsr)
		if err != nil {
			return nil, err
		}
		return gsr, nil
	})

	// Esperar resultado con timeout
	select {
	case result := <-resultChan:
		if result.Err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Err.Error()})
			return
		}

		savedgsr := result.Data.(domain.Gsr)
		c.JSON(http.StatusCreated, gin.H{
			"status": true,
			"data": gin.H{
				"type": "gsr",
				"attributes": gin.H{
					"porcentaje": savedgsr.Porcentaje,
				},
			},
		})
	}
}