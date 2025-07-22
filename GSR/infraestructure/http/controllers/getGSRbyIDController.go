package controllers

import (
	"API-VITALVEST/GSR/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetGSRbyIDController struct {
	uc *application.GetGSRbyID_UC
}

func NewGetGSRbyIDController(uc *application.GetGSRbyID_UC) *GetGSRbyIDController{
	return &GetGSRbyIDController{uc:uc}
}

func (ctrl *GetGSRbyIDController) Run(c *gin.Context) {
	id := c.Param("id")
	ID, err := strconv.Atoi(id)

	if err != nil {
		 c.JSON(http.StatusBadRequest, gin.H{"error": "El id debe ser numérico"})
		 return
	}

	gsr, err := ctrl.uc.Run(ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"gsr by ID": gsr})
	}
}