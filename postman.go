package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	data := make(map[string]string)
	data["a"] = "12"
	router := gin.Default()
	router.GET("/key/:data1", func(c *gin.Context) {
		data1 := c.Param("data1")
		value, exists := data[data1]
		if exists {
			c.String(http.StatusOK, "Key %s, Value: %s", data1, value)
		} else {
			c.String(http.StatusNotFound, "Key %s not exist", data1)
		}
	})

	router.POST("/store", func(c *gin.Context) {
		var vertex struct {
			Key   string `json:"key"`
			Value string `json:"value"`
		}

		if err := c.BindJSON(&vertex); err != nil {
			c.String(http.StatusBadRequest, "Donot added")
			return
		}

		data[vertex.Key] = vertex.Value
		c.String(http.StatusOK, "key-value : %s", data)
		
		c.JSON(http.StatusOK, gin.H{"message": "key and value successfully added"})
	})

	router.Run(":8080")
}
