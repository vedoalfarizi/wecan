package rest

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/vedoalfarizi/wecan/src/Infrastructures/database/postgresql"
	"github.com/vedoalfarizi/wecan/src/Infrastructures/google"
	"github.com/vedoalfarizi/wecan/src/models"
	"net/http"
)

func GetFundraiserSheetHandler(c *gin.Context) {
	id := c.Param("id")

	var fundraiser models.Fundraiser
	if err := postgresql.DB.Where("id = ?", id).First(&fundraiser).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if fundraiser.SheetID == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sheet not found"})
		return
	}

	ctx := context.Background()
	gSheet := google.NewGSheet(ctx)
	spreadsheet, err := gSheet.GetSpreadsheet(fundraiser.SheetID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": spreadsheet.SpreadsheetUrl})
}
