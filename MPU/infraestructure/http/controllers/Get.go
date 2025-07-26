package controllers

import (
	"API-VITALVEST/MPU/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAllTableMPU_Controller struct {
	uc *application.GetAllTableMPU_UC
}

func NewGetAllTableMpuController(uc *application.GetAllTableMPU_UC) *GetAllTableMPU_Controller{
	return &GetAllTableMPU_Controller{uc: uc}
}

func (ctrl *GetAllTableMPU_Controller) Run(c *gin.Context) {
	mpus, err := ctrl.uc.Run()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Todos los campos son requeridos"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"MPUs" : mpus})
}