package controllers

import (
	"API-VITALVEST/BME/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAllBMEController struct {
	uc *application.GetAllBME_UC
}

func NewGetAllBMEController(uc *application.GetAllBME_UC) *GetAllBMEController{
	return &GetAllBMEController{uc: uc}
}

func (ctrl *GetAllBMEController) Run(c *gin.Context) {
	BMEs, err := ctrl.uc.Run()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Todos los campos son requeridos"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"BME" : BMEs})
}