package controllers

import (
	"API-VITALVEST/MLX/application"
	"API-VITALVEST/MLX/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SaveMLXController struct {
	uc *application.SaveMLX_uc
}

func NewSaveMLXController(uc *application.SaveMLX_uc) *SaveMLXController {
	return &SaveMLXController{uc: uc}
}

func (ctrl *SaveMLXController) Run(c *gin.Context) {
	var mlx domain.Mlx

	if err := c.ShouldBindJSON(&mlx); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Todos los campos son requeridos"})
		return
	} 

	err := ctrl.uc.Run(mlx)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusCreated, gin.H{
			"status": true,
			"data": gin.H{
				"type": "MLX",
				"attributes": gin.H{
					"ID del sensor de MLX": mlx.Id,
					"Temperatura corporal":mlx.Temperatura_corporal,
				},
			},
		})
	}
}