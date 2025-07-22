package controllers

import (
	"API-VITALVEST/GSR/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAllGSRController struct {
	uc *application.GetAllGsr_UC
}

func NewGetAllGSRController(uc *application.GetAllGsr_UC) *GetAllGSRController{
	return &GetAllGSRController{uc: uc}
}

func (ctrl *GetAllGSRController) Run(c *gin.Context) {
	gsrs, err := ctrl.uc.Run()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Todos los campos son requeridos"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"GSR" : gsrs})
}