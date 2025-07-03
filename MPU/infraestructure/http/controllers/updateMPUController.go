package controllers

import (
	"API-VITALVEST/MPU/application"
	"API-VITALVEST/MPU/domain"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateMPUController struct {
	uc *application.UpdateMPU
}

func NewUpdateMPUController(uc *application.UpdateMPU) *UpdateMPUController {
	return &UpdateMPUController{uc: uc}
}

func (ctrl *UpdateMPUController) Run(c *gin.Context) {
	var MPU domain.Mpu
	if err := c.ShouldBindJSON(&MPU); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Todos los campos son requeridos"})
		return
	}

	ID := c.Param("id")
	id, err := strconv.Atoi(ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "El id debe ser un número"})
		return
	}
	if err := ctrl.uc.Run(id, MPU); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} 
		c.JSON(http.StatusOK, gin.H{"message": "MPU actualizado correctamente"})
}