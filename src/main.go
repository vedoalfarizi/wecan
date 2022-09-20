package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vedoalfarizi/wecan/src/Infrastructures/database/postgresql"
	"github.com/vedoalfarizi/wecan/src/handlers/rest"
)

func main() {
	r := gin.Default()

	postgresql.ConnectDatabase()

	r.GET("/fundraisers", rest.GetFundraisersHandler)
	r.POST("/fundraisers", rest.AddFundraiserHandler)
	r.GET("/fundraisers/:id", rest.FindOneFundraiserHandler)
	r.PATCH("/fundraisers/:id", rest.UpdateFundraisersHandler)
	r.DELETE("/fundraisers/:id", rest.DeleteFundraisersHandler)

	r.POST("/fundraisers/:id/disburse", rest.AddDisbursement)

	r.GET("/fundraisers/:id/sheets", rest.GetFundraiserSheetHandler)

	r.Run()
}
