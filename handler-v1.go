package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func imgUploadHandler(c *gin.Context) {
	// Multipart form
	form, _ := c.MultipartForm()
	if form == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "no form data",
		})
		return
	}
	files := form.File["upload[]"]
	var cNames []string

	for _, file := range files {
		log.Println(file.Filename)

		// Upload the file to specific dst.
		cName, err := cacheImageFromFile(file)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "failed to upload image",
			})
			return
		}
		cNames = append(cNames, cName)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"names":   cNames,
	})
}

func imgGetHandler(c *gin.Context) {
	name := c.Query("name")
	imgData, ok := imgCache[name]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "image not found",
		})
		return
	}

	c.Data(http.StatusOK, "image/jpeg", imgData)
}
