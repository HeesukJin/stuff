package main

import (
	"github.com/zmb3/spotify/v2/auth"
	"log"
	"os"
	"github.com/zmb3/spotify/v2"
	"golang.org/x/oauth2/clientcredentials"

)

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

func getSpotifyPlaylistSongs(w http.ResponseWriter, r *http.Request, client *spotify.Client) {
	results, err := client.GetPlaylistItems(ctx, "59eBtjSiUluqHX1BweDZgu")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	fmt.Fprintf(w, results.Items[2].Track.Track.Name)
}