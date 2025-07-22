package controllers

import (
	"API-VITALVEST/MLX/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetMLXbyIDController struct {
	uc *application.GetMlxbyID_UC
}

func NewGetMLXbyIDController(uc *application.GetMlxbyID_UC) *GetMLXbyIDController{
	return &GetMLXbyIDController{uc:uc}
}

func (ctrl *GetMLXbyIDController) Run(c *gin.Context) {
	id := c.Param("id")
	ID, err := strconv.Atoi(id)

	if err != nil {
		 c.JSON(http.StatusBadRequest, gin.H{"error": "El id debe ser numérico"})
		 return
	}

	MLX, err := ctrl.uc.Run(ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"MLX by ID": MLX})
	}
}