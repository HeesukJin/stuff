package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
	"golang.org/x/oauth2/clientcredentials"
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

	http.HandleFunc("/get_playlist_songs", func(w http.ResponseWriter, r *http.Request) {
		results, err := client.GetPlaylistItems(ctx, "59eBtjSiUluqHX1BweDZgu")
		if err != nil {
			log.Fatalf("error: %v", err)
		}

		fmt.Fprintf(w, results.Items[2].Track.Track.Name)

	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello!")
	})

	fmt.Printf("Starting server at port 8080\n")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func authClient() *spotify.Client {
	ctx := context.Background()
	config := &clientcredentials.Config{
		ClientID:     os.Getenv("SPOTIFY_ID"),
		ClientSecret: os.Getenv("SPOTIFY_SECRET"),
		TokenURL:     spotifyauth.TokenURL,
	}
	token, err := config.Token(ctx)
	if err != nil {
		log.Fatalf("couldn't get token: %v", err)
	}

	httpClient := spotifyauth.New().Client(ctx, token)
	return spotify.New(httpClient)
}
