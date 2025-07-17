package controllers

import (
	"API-VITALVEST/alertas/application"
	"API-VITALVEST/alertas/domain"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type UpdateAlertaController struct {
	uc *application.UpdateAlerta
}

func NewUpdateAlertaController(uc *application.UpdateAlerta) *UpdateAlertaController {
	return &UpdateAlertaController{uc: uc}
}

func (ctrl *UpdateAlertaController) Run(c *gin.Context) {
	var alerta domain.Alerta

	// Validar JSON
	if err := c.ShouldBindJSON(&alerta); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido o campos requeridos faltantes"})
		return
	}

	// Validar el ID del path
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "El ID debe ser un número entero"})
		return
	}

	// Validar que el sensor sea uno permitido (enum lógico)
	if err := alerta.Validar(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if alerta.Fecha.IsZero() {
		alerta.Fecha = time.Now()
	}

	// Ejecutar caso de uso
	if err := ctrl.uc.Run(id, alerta); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Alerta actualizada correctamente",
	})
}
