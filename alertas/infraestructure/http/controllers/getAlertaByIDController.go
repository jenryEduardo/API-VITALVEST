package controllers

import (
	"API-VITALVEST/alertas/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetAlertaByIDController struct {
	uc *application.GetAlertaByID
}

func NewGetAlertaByIDController(uc *application.GetAlertaByID) *GetAlertaByIDController {
	return &GetAlertaByIDController{uc: uc}
}

func (ctrl *GetAlertaByIDController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "El ID debe ser un número entero"})
		return
	}

	Alerta, err := ctrl.uc.Run(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"Alerta by ID": Alerta})
	}
}
