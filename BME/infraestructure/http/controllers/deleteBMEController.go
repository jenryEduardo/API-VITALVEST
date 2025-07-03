package controllers

import (
	"API-VITALVEST/BME/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteBMEController struct {
	uc *application.DeleteBME
}

func NewDeleteBMEController(uc *application.DeleteBME) *DeleteBMEController{
	return &DeleteBMEController{uc: uc}
}
func (ctrl *DeleteBMEController) Run(c *gin.Context) {
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
		c.JSON(http.StatusOK, gin.H{"message": "BME eliminado correctamente"})
	}
}