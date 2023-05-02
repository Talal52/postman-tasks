// package main

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// type response struct {
// 	Key   string `json:"key"`
// 	Value string `json:"value"`
// }
// type vertex struct {
// 	Key   string `json:"key"`
// 	Value string `json:"value"`
// }

// func main() {
// 	data := make(map[string]string)
// 	router := gin.Default()
// 	router.GET("/key/:input", func(c *gin.Context) {
// 		input := c.Param("input")
// 		value, exists := data[input]
// 		if exists {
// 			c.JSON(http.StatusOK, gin.H{
// 				"key":   input,
// 				"value": value,
// 			})
// 		} else {
// 			c.JSON(http.StatusNotFound, gin.H{
// 				"error message": "entered Key is not found!!!",
// 			})
// 		}
// 	})
// 	router.DELETE("/delete/:input", func(c *gin.Context) {
// 		input := c.Param("input")
// 		_, exists := data[input]
// 		if exists {
// 			var vertex struct {
// 				Value string `json:"value"`
// 			}
// 			delete(data, vertex.Value)
// 			c.JSON(http.StatusOK, gin.H{
// 				"message": "deleted successfully",
// 				"data":    data,
// 			})
// 		} else {
// 			c.JSON(http.StatusNotFound, gin.H{
// 				"error": "entered key not found",
// 			})
// 		}
// 	})
// 	router.PUT("/update/:input", func(c *gin.Context) {
// 		input := c.Param("input")
// 		_, exists := data[input]
// 		if exists {
// 			var vertex struct {
// 				Value string `json:"value"`
// 			}

// 			data[input] = vertex.Value
// 			c.JSON(http.StatusOK, gin.H{
// 				"output": "successfully update the entered value",
// 				"data":   data,
// 			})
// 		} else {
// 			c.JSON(http.StatusNotFound, gin.H{
// 				"error": "entered key is not found",
// 			})
// 		}
// 	})
// 	router.POST("/store", func(c *gin.Context) {

// 	var vertx vertex
// 		if err := c.BindJSON(&vertx); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"message": "Invalid input",
// 			})
// 		}
// 		for key,value :=range data{
// 			vertx.Key=key
// 			vertx.Value=value
// 			c.JSON(http.StatusOK, gin.H{
// 				"message": "Key-value pair added successfully",
// 				"data": vertx,
// 			})

// 		}
// 	})
// 	router.GET("/display", func(c *gin.Context) {
// 		var res response
// 		// var totalValues []string
// 		for key, value := range data {
// 			res.Key = key
// 			res.Value = value
// 			c.JSON(http.StatusOK, gin.H{
// 				"message": "map key-values",
// 				"data":    res,
// 			})
// 		}

// 	})
// 	router.Run(":8080")
// }





package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type response struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

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
	})

	router.PUT("/update/:input", func(c *gin.Context) {
		input := c.Param("input")
		_, exists := data[input]
		if exists {
			var requestBody response
			if err := c.BindJSON(&requestBody); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": "Invalid input",
				})
				return
			}
			data[input] = requestBody.Value
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
	})

	router.GET("/display", func(c *gin.Context) {
		var responseData []response
		for key, value := range data {
			responseData = append(responseData, response{key, value})
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "map key-values",
			"data":    responseData,
		})
	})

	router.Run(":8080")
}
