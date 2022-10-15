package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
		v1.POST("/img", imgUploadHandler)
		v1.GET("/img", imgGetHandler)
	}

	r.Run() // listen and serve on
}
