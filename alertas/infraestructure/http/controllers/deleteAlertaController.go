package controllers

import (
	"API-VITALVEST/alertas/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteAlertaController struct {
	uc *application.DeleteAlerta
}

func NewDeleteAlertaController(uc *application.DeleteAlerta) *DeleteAlertaController {
	return &DeleteAlertaController{uc: uc}
}

func (ctrl *DeleteAlertaController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "El ID debe ser un número entero válido"})
		return
	}

	if err := ctrl.uc.Run(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Alerta eliminada correctamente"})
}
