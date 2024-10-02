package main

import (
	"github.com/gin-gonic/gin"
)

// shape of an album
type album struct {
	Id     string  `json:id`
	Title  string  `json:title`
	Artist string  `json:artist`
	Price  float32 `json:price`
}

// seed data
var albums = []album{
	{Id: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.66},
	{Id: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{Id: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	var r = gin.Default()

	// controllers
	r.GET("/albums", getAlbums)
	r.GET("/albums/:id", getAlbumById)
	r.POST("/albums", createAlbum)

	r.Run("127.0.0.1:3000")
}

// controllers
func getAlbums(c *gin.Context) {
	c.IndentedJSON(200, albums)
}

func getAlbumById(c *gin.Context) {
	var id = c.Param("id")

	for _, a := range albums {
		if a.Id == id {
			c.IndentedJSON(200, a)
			return
		}
	}

	c.IndentedJSON(404, gin.H{"message": "album not found"})
}

func createAlbum(c *gin.Context) {
	var newAlbum album

	// Call "BindJSON" to bind/attach the 'received' JSON in the request body
	// to newAlbum.
	var err = c.BindJSON(&newAlbum)
	if err != nil {
		return
	}

	// add the new album to the db
	albums = append(albums, newAlbum)
	c.IndentedJSON(201, newAlbum)
}

// inspired by: https://go.dev/doc/tutorial/web-service-gin
