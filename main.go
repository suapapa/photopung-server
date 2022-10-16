package main

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
		v1.POST("/img", imgUploadHandler)
		v1.GET("/img", imgGetHandler)
	}

	ctx, cancelF := context.WithCancel(context.Background())
	defer cancelF()
	go cleanGarbageImageCache(ctx)

	r.Run(":8080") // listen and serve on
}
