package controllers

import (
	"API-VITALVEST/MPU/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetMPUbyIDController struct {
	uc *application.GetMPUbyID_UC
}

func NewGetMPUbyIDController(uc *application.GetMPUbyID_UC) *GetMPUbyIDController{
	return &GetMPUbyIDController{uc:uc}
}

func (ctrl *GetMPUbyIDController) Run(c *gin.Context) {
	id := c.Param("id")
	ID, err := strconv.Atoi(id)

	if err != nil {
		 c.JSON(http.StatusBadRequest, gin.H{"error": "El id debe ser numérico"})
		 return
	}

	MPU, err := ctrl.uc.Run(ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"MPU by ID": MPU})
	}
}