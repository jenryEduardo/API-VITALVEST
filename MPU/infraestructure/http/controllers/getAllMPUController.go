package controllers

import (
	"API-VITALVEST/MPU/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAllMPUController struct {
	uc *application.GetAllMPU_UC
}

func NewGetAllMPUController(uc *application.GetAllMPU_UC) *GetAllMPUController{
	return &GetAllMPUController{uc: uc}
}

func (ctrl *GetAllMPUController) Run(c *gin.Context) {
	total, err := ctrl.uc.Run()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Todos los campos son requeridos"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"pasos":total})
}