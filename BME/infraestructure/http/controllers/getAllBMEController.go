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
    BMEs, err := ctrl.uc.Run(c.Request.Context())
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{"BME": BMEs})
}
