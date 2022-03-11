package main

import (
	"net/http"
	"webapi/calclib"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/min", calculateMin)
	router.POST("/max", calculateMax)
	router.POST("/avg", calculateAvg)
	router.POST("/median", calculateMedian)
	router.POST("/percentile", calculatePercentile)

	router.Run("localhost:8080")
}

// getAlbums responds with the list of all albums as JSON.
func calculateMin(c *gin.Context) {

	var data calclib.InputWrapper

	if err := c.BindJSON(&data); err != nil {
		return
	}
	res := calclib.CalculateMin(data)
	c.IndentedJSON(http.StatusOK, res)
}

func calculateMax(c *gin.Context) {

	var data calclib.InputWrapper

	if err := c.BindJSON(&data); err != nil {
		return
	}

	res := calclib.CalculateMax(data)
	c.IndentedJSON(http.StatusOK, res)
}

func calculateAvg(c *gin.Context) {
	var list []float64
	if err := c.BindJSON(&list); err != nil {
		return
	}
	res := calclib.CalculateAvg(list)
	c.IndentedJSON(http.StatusOK, res)
}

func calculateMedian(c *gin.Context) {
	var list []float64
	if err := c.BindJSON(&list); err != nil {
		return
	}

	res := calclib.CalculateMedian(list)
	c.IndentedJSON(http.StatusOK, res)
}

func calculatePercentile(c *gin.Context) {

	var data calclib.PercentileInputWrapper

	if err := c.BindJSON(&data); err != nil {
		return
	}

	res := calclib.CalculatePercentile(data)

	c.IndentedJSON(http.StatusOK, res)
}
