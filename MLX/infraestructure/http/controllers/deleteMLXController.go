package controllers

import (
	"API-VITALVEST/MLX/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteMLXController struct {
	uc *application.DeleteMLX
}

func NewDeleteMLXController(uc *application.DeleteMLX) *DeleteMLXController{
	return &DeleteMLXController{uc: uc}
}
func (ctrl *DeleteMLXController) Run(c *gin.Context) {
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
		c.JSON(http.StatusOK, gin.H{"message": "MLX eliminado correctamente"})
	}
}