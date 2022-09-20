package rest

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/vedoalfarizi/wecan/src/Infrastructures/database/postgresql"
	"github.com/vedoalfarizi/wecan/src/Infrastructures/google"
	"github.com/vedoalfarizi/wecan/src/models"
	"google.golang.org/api/sheets/v4"
	"net/http"
	"time"
)

func createSpreadSheet(ctx context.Context, title string, fundraiserID uint) (sheetID string, err error) {
	gSheet := google.NewGSheet(ctx)

	sheetProps := &sheets.Spreadsheet{
		Properties: &sheets.SpreadsheetProperties{
			Title: fmt.Sprintf("%d - %s", fundraiserID, title),
		},
	}

	sheetID, err = gSheet.CreateSpreadsheet(sheetProps)
	if err != nil {
		return
	}

	headerProps := sheets.ValueRange{
		MajorDimension: "ROWS",
		Range:          "Sheet1!A1:F1",
		Values: [][]interface{}{
			{
				"Code",
				"Tanggal",
				"Tujuan",
				"Jumlah",
				"Rekening Tujuan",
				"Rekening Penerima",
			},
		},
	}

	// Set header of sheet
	err = gSheet.UpdateSpreadsheet(sheetID, &headerProps)
	if err != nil {
		return
	}

	// set access of spreadsheet

	return
}

func AddDisbursement(c *gin.Context) {
	var payload models.Disbursement
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fundraiserID := c.Param("id")

	var fundraiser models.Fundraiser
	if err := postgresql.DB.Where("id = ?", fundraiserID).First(&fundraiser).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := context.Background()
	gSheet := google.NewGSheet(ctx)

	if fundraiser.SheetID == "" {
		sheetID, err := createSpreadSheet(ctx, fundraiser.Name, fundraiser.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		reqProtect := &sheets.BatchUpdateSpreadsheetRequest{
			Requests: []*sheets.Request{
				{
					AddProtectedRange: &sheets.AddProtectedRangeRequest{
						ProtectedRange: &sheets.ProtectedRange{
							Range: &sheets.GridRange{
								SheetId: 0, // default value for first sheet
							},
							Editors: &sheets.Editors{
								Users: []string{
									"vedoalfarizi@gmail.com",
								},
							},
						},
					},
				},
			},
		}

		err = gSheet.BatchUpdateSpreadsheet(sheetID, reqProtect)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		var updatePayload models.UpdateFundraiserSheetPayload
		updatePayload.SheetID = sheetID
		postgresql.DB.Model(&fundraiser).Updates(updatePayload)
	}

	loc, _ := time.LoadLocation("Asia/Jakarta")
	disburse := models.Disbursement{
		ID:           payload.ID,
		FundraiserID: fundraiser.ID,
		Purpose:      payload.Purpose,
		Amount:       payload.Amount,
		Bank:         payload.Bank,
		AccHolder:    payload.AccHolder,
		DisburseAt:   time.Now().In(loc).Format("02-01-2006"),
	}

	rowValues := sheets.ValueRange{
		MajorDimension: "ROWS",
		Range:          "Sheet1!A1:F1",
		Values: [][]interface{}{
			{
				disburse.ID,
				disburse.DisburseAt,
				disburse.Purpose,
				disburse.Amount,
				disburse.Bank,
				disburse.AccHolder,
			},
		},
	}

	err := gSheet.AppendValue(fundraiser.SheetID, &rowValues)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": fundraiser.SheetID})
}
