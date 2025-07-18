package controllers

import (
	"API-VITALVEST/alertas/application"
	"API-VITALVEST/alertas/domain"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type SaveAlertaController struct {
	uc *application.SaveAlerta_UC
}

func NewSaveAlertaController(uc *application.SaveAlerta_UC) *SaveAlertaController {
	return &SaveAlertaController{uc: uc}
}

func (ctrl *SaveAlertaController) Run(c *gin.Context) {
	var alerta domain.Alerta

	if err := c.ShouldBindJSON(&alerta); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Todos los campos son requeridos o inválidos"})
		return
	}	

	// Validar sensor permitido
	if err := alerta.Validar(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Asignar fecha actual si no se proporcionó
	if alerta.Fecha.IsZero() {
		alerta.Fecha = time.Now()
	}

	if err := ctrl.uc.Run(alerta); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}


	c.JSON(http.StatusCreated, gin.H{
		"status": true,
		"data": gin.H{
			"type": "alerta",
			"attributes": gin.H{
				"id":                        alerta.ID,
				"nombre_del_sensor":         alerta.NombreDelSensor,
				"fecha":                     alerta.Fecha.Format(time.RFC3339),
				"cantidad_de_veces_enviado": alerta.CantidadDeVecesEnviado,
			},
		},
	})

}
