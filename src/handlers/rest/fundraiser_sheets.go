package rest

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/vedoalfarizi/wecan/src/Infrastructures/database/postgresql"
	"github.com/vedoalfarizi/wecan/src/Infrastructures/google"
	"github.com/vedoalfarizi/wecan/src/models"
	"google.golang.org/api/drive/v3"
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
	s, err := gSheet.GetSpreadsheet(fundraiser.SheetID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	gDrive := google.NewGDrive(ctx)
	sheetPermission := &drive.Permission{
		Type: "anyone",
		Role: "reader",
	}
	err = gDrive.AddPermission(s.SpreadsheetId, sheetPermission)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": s.SpreadsheetUrl})
}
