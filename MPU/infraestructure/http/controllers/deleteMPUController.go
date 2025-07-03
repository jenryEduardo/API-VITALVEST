package controllers

import (
	"API-VITALVEST/MPU/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteMPUController struct {
	uc *application.DeleteMPU
}

func NewDeleteMPUController(uc *application.DeleteMPU) *DeleteMPUController{
	return &DeleteMPUController{uc: uc}
}
func (ctrl *DeleteMPUController) Run(c *gin.Context) {
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
		c.JSON(http.StatusOK, gin.H{"message": "MPU eliminado correctamente"})
	}
}