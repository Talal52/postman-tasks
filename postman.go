package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	data := make(map[string]string)
	router := gin.Default()
	router.GET("/key/:data1", func(c *gin.Context) {
		data1 := c.Param("data1")
		value, exists := data[data1]
		if exists {
			c.JSON(http.StatusOK, gin.H{
				"key":   data1,
				"value": value,
			})
		} else {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Key not found",
			})
		}
	})
	router.POST("/store", func(c *gin.Context) {
		var vertex struct {
			Key   string `json:"key"`
			Value string `json:"value"`
		}

		if err := c.BindJSON(&vertex); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid request body",
			})
		}

		data[vertex.Key] = vertex.Value
		c.JSON(http.StatusOK, gin.H{
			"message": "Key-value pair added successfully",
			"data":    data,
		})
	})

	router.Run(":8080")
}
