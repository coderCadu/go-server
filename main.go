package main

import (
	"github.com/codercadu/go-server/api"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/albums", api.GetAlbums)
	router.GET("/albums/:id", api.GetAlbumByID)
	router.POST("/albums", api.PostAlbums)

	router.Run("localhost:8080")
}
