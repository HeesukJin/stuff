package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Playlist struct {
	Origin string
	Songs   []Song
}
type Song struct {
	Name      string
	SpotifyId string
	YotubeID  string
	Artist    string
}

var SpotifyClient = authClient()
var YoutubeClient = authYoutube()

func main() {
	router := gin.Default()

	router.GET("/get_spotify_playlist_songs", getSpotifyPlaylistSongs)
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "f",
		})
	})

	fmt.Printf("Starting server at port 8080\n")

	router.Run()
}
