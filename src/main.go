package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vedoalfarizi/wecan/src/Infrastructures/database/postgresql"
)

func main() {
	r := gin.Default()

	postgresql.ConnectDatabase()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	r.Run()
}
