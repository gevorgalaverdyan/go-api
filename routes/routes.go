package routes

import (
	"net/http"

	"github.com/gevorgalaverdyan/go-api/models"
	"github.com/gin-gonic/gin"
)

func getAllFacilities(ctx *gin.Context){
	healthFacilities, err := models.GetAll()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, healthFacilities)
}

func getFacilityById(ctx *gin.Context) {
	id := ctx.Param("id")

	if id=="" {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "error",
		})
		return
	}

	facility, err := models.GetById(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, facility)
}