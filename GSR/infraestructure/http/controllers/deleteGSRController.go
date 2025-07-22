package controllers

import (
	"API-VITALVEST/GSR/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteGSRController struct {
	uc *application.DeleteGSR
}

func NewDeleteGSRController(uc *application.DeleteGSR) *DeleteGSRController{
	return &DeleteGSRController{uc: uc}
}
func (ctrl *DeleteGSRController) Run(c *gin.Context) {
	ID := c.Param("id")
	id, err := strconv.Atoi(ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "El id debe ser un número"})
		return
	}

	err = ctrl.uc.Run(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "GSR eliminado correctamente"})
	}
}