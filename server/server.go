package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Playlist struct {
	origin string
	song   []Songs
}
type Songs struct {
	name      string
	spotifyId string
	yotubeID  string
	artist    string
}

var Client = authClient()

func main() {
	router := gin.Default()

	router.GET("/get_spotify_playlist_songs", getSpotifyPlaylistSongs)
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	fmt.Printf("Starting server at port 8080\n")

	router.Run()
}
