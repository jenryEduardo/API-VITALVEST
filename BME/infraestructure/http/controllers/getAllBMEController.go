package controllers

import (
	"API-VITALVEST/BME/application"
	"API-VITALVEST/core/workerpool"
	"time"

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
	// Submit retorna directamente el canal de resultados
	resultChan := ctrl.pool.Submit(func() (interface{}, error) {
		return ctrl.uc.Run()
	})
	
	// Esperar resultado con timeout
	select {
	case result := <-resultChan:
		if result.Err != nil {
			c.JSON(400, gin.H{"error": result.Err.Error()})
			return
		}
		c.JSON(200, gin.H{"BME": result.Data})
		
	case <-time.After(5 * time.Second):
		c.JSON(504, gin.H{"error": "request timeout"})
	}
}