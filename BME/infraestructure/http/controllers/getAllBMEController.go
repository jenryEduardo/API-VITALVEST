package controllers

import (
	"API-VITALVEST/BME/application"
	"API-VITALVEST/core/workerpool"

	"github.com/gin-gonic/gin"
)

type GetAllBMEController struct {
	uc   *application.GetAllBME_UC
	pool *workerpool.WorkerPool
}

func NewGetAllBMEController(uc *application.GetAllBME_UC, pool *workerpool.WorkerPool) *GetAllBMEController {
	return &GetAllBMEController{uc: uc, pool: pool}
}

func (ctrl *GetAllBMEController) Run(c *gin.Context) {

	resultChan := make(chan any)

	// 🚀 Enviar el GET a un worker de manera concurrente
	ctrl.pool.Submit(func() any {
		data, err := ctrl.uc.Run()
		if err != nil {
			resultChan <- err
			return nil
		}
		resultChan <- data
		return nil
	})

	// Esperar el resultado
	result := <-resultChan

	// Validar errores
	if err, ok := result.(error); ok {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"BME": result})
}
