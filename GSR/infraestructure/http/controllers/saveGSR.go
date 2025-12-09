package controllers

import (
	"API-VITALVEST/GSR/application"
	"API-VITALVEST/GSR/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SaveGSRController struct {
	uc *application.SaveGSR_UC
}

func NewSaveGSRController(uc *application.SaveGSR_UC) *SaveGSRController {
	return &SaveGSRController{uc: uc}
}

func (ctrl *SaveGSRController) Run(c *gin.Context) {
	var gsr domain.Gsr

	if err := c.ShouldBindJSON(&gsr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Todos los campos son requeridos"})
		return
	} 

	err := ctrl.uc.Run(gsr)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusCreated, gin.H{
			"status": true,
			"data": gin.H{
				"type": "gsr",
				"attributes": gin.H{
					"porcentaje": gsr.Porcentaje,
				},
			},
		})
	}
}