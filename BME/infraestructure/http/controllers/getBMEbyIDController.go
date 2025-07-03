package controllers

import (
	"API-VITALVEST/BME/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetBMEbyIDController struct {
	uc *application.GetBMEbyID_UC
}

func NewGetBMEbyIDController(uc *application.GetBMEbyID_UC) *GetBMEbyIDController{
	return &GetBMEbyIDController{uc:uc}
}

func (ctrl *GetBMEbyIDController) Run(c *gin.Context) {
	id := c.Param("id")
	ID, err := strconv.Atoi(id)

	if err != nil {
		 c.JSON(http.StatusBadRequest, gin.H{"error": "El id debe ser numérico"})
		 return
	}

	BME, err := ctrl.uc.Run(ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"BME by ID": BME})
	}
}