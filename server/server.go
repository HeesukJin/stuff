package main

import (
	"context"

	"fmt"
	"log"
	"net/http"
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

func main() {
	ctx := context.Background()
	client := authClient()

	http.HandleFunc("/get_playlist_songs", getSpotifyPlaylistSongs(w http.ResponseWriter, r *http.Request, client))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello!")
	})

	fmt.Printf("Starting server at port 8080\n")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}


