package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	data := make(map[string]string)
	router := gin.Default()
	router.GET("/key/:input", func(c *gin.Context) {
		input := c.Param("input")
		value, exists := data[input]
		if exists {
			c.JSON(http.StatusOK, gin.H{
				"key":   input,
				"value": value,
			})
		} else {
			c.JSON(http.StatusNotFound, gin.H{
				"error message": "entered Key is not found!!!",
			})
		}
	})
	router.DELETE("/delete/:input", func(c *gin.Context) {
		input := c.Param("input")
		_, exists := data[input]
		if exists {
			var vertex struct {
				Value string `json:"value"`
			}
			delete(data, vertex.Value)
			c.JSON(http.StatusOK, gin.H{
				"message": "deleted successfully",
				"data":    data,
			})
		} else {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "entered key not found",
			})
		}
	})
	router.PUT("/update/:input", func(c *gin.Context) {
		input := c.Param("input")
		_, exists := data[input]
		if exists {
			var vertex struct {
				Value string `json:"value"`
			}

			data[input] = vertex.Value
			c.JSON(http.StatusOK, gin.H{
				"output": "successfully update the entered value",
				"data":   data,
			})
		} else {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "entered key is not found",
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
				"error message": "Invalid input",
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
