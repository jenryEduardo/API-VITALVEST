package controllers

import (
	"API-VITALVEST/GSR/application"
	"API-VITALVEST/GSR/domain"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateGSRController struct {
	uc *application.UpdateGSR
}

func NewUpdateGSRController(uc *application.UpdateGSR) *UpdateGSRController {
	return &UpdateGSRController{uc: uc}
}

func (ctrl *UpdateGSRController) Run(c *gin.Context) {
	var gsr domain.Gsr
	if err := c.ShouldBindJSON(&gsr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Todos los campos son requeridos"})
		return
	}

	ID := c.Param("id")
	id, err := strconv.Atoi(ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "El id debe ser un número"})
		return
	}
	if err := ctrl.uc.Run(id, gsr); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} 
		c.JSON(http.StatusOK, gin.H{"message": "GSR actualizado correctamente"})
}