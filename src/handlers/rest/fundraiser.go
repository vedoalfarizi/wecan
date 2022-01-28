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
