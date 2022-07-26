package main

import (
	"context"
	"fmt"
	"log"
	//"os"

	"github.com/gin-gonic/gin"
	"github.com/zmb3/spotify/v2"
	"github.com/zmb3/spotify/v2/auth"
	"golang.org/x/oauth2/clientcredentials"
)

func authClient() *spotify.Client {
	ctx := context.Background()
	config := &clientcredentials.Config{
		ClientID:     "5731fe965737499186663c59665746d5",
		ClientSecret: "ce32ecddb8d84b13adc81413fe1d9242",
		TokenURL:     spotifyauth.TokenURL,
	}
	token, err := config.Token(ctx)
	if err != nil {
		log.Fatalf("couldn't get token: %v", err)
	}

	httpClient := spotifyauth.New().Client(ctx, token)
	return spotify.New(httpClient)
}

func  getSpotifyPlaylistSongs(c *gin.Context) {
	results, err := Client.GetPlaylistItems(c, "59eBtjSiUluqHX1BweDZgu")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Println(results)

	c.String(200, results.Items[2].Track.Track.Name)
}