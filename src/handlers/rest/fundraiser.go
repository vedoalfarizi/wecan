package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vedoalfarizi/wecan/src/Infrastructures/database/postgresql"
	"github.com/vedoalfarizi/wecan/src/models"
)

func GetFundraisersHandler(c *gin.Context) {
	var fundraisers []models.Fundraiser
	postgresql.DB.Find(&fundraisers)

	c.JSON(http.StatusOK, gin.H{"data": fundraisers})
}

func AddFundraiserHandler(c *gin.Context) {
	var payload models.CreateFundraiserPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fundraiser := models.Fundraiser{
		Name:        payload.Name,
		Description: payload.Description,
	}

	postgresql.DB.Create(&fundraiser)

	c.JSON(http.StatusCreated, gin.H{"data": fundraiser})
}
