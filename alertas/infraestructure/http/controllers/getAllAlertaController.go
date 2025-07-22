package controllers

import (
	"API-VITALVEST/alertas/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAllAlertasController struct {
	uc *application.GetAllAlertas_UC
}

func NewGetAllAlertasController(uc *application.GetAllAlertas_UC) *GetAllAlertasController {
	return &GetAllAlertasController{uc: uc}
}

func (ctrl *GetAllAlertasController) Run(c *gin.Context) {
	alertas, err := ctrl.uc.Run()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudieron obtener las alertas"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   alertas,
	})
}
