package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type response struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
	var data = make(map[string]string)

func getKey(c *gin.Context) {
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
}

func deleteKey(c *gin.Context) {
	input := c.Param("input")
	_, exists := data[input]
	if exists {
		delete(data, input)
		c.JSON(http.StatusOK, gin.H{
			"message": "deleted successfully",
			"data":    data,
		})
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "entered key not found",
		})
	}
}

func updateKey(c *gin.Context) {
	var req response
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid input",
		})
		return
	}
	_, exists := data[req.Key]
	if exists {
		data[req.Key] = req.Value
		c.JSON(http.StatusOK, gin.H{
			"output": "successfully updated the entered value",
			"data":   data,
		})
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "entered key is not found",
		})
	}
}

func storeKey(c *gin.Context) {
	var requestBody response
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid input",
		})
		return
	}
	data[requestBody.Key] = requestBody.Value
	c.JSON(http.StatusOK, gin.H{
		"message": "Key-value pair added successfully",
		"data":    data,
	})
}

func displayKeys(c *gin.Context) {
	var responseData []response
	for key, value := range data {
		responseData = append(responseData, response{key, value})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "map key-values",
		"data":    responseData,
	})
}
func main() {
	router := gin.Default()
	router.GET("/key/:input", getKey)
	router.DELETE("/delete/:input", deleteKey)
	router.PUT("/update/:req.key", updateKey)
	router.POST("/store", storeKey)
	router.GET("/display", displayKeys)
	
	router.Run(":8080")
}
