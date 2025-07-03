package controllers

import (
	"API-VITALVEST/MLX/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAllMLXController struct {
	uc *application.GetAllMlx_UC
}

func NewGetAllMLXController(uc *application.GetAllMlx_UC) *GetAllMLXController{
	return &GetAllMLXController{uc: uc}
}

func (ctrl *GetAllMLXController) Run(c *gin.Context) {
	MLXs, err := ctrl.uc.Run()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Todos los campos son requeridos"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"MLX" : MLXs})
}