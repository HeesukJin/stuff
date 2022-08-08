package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"github.com/gin-gonic/gin"
	"github.com/zmb3/spotify/v2"
	"github.com/zmb3/spotify/v2/auth"
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

func  getSpotifyPlaylistSongs(c *gin.Context) {
	
	url := c.Request.URL.Query()["spotifyPlaylistURL"][0]

	//fmt.Println(url);
	playlistID:= spotify.ID(strings.Split(strings.Split(url, "playlist/")[1], "?")[0])
	//playlistID := String(strings.Split(strings.Split(url, "playlist/")[1], "?")[0])

	// function should pass in spotify playlist id

	
	results, err := Client.GetPlaylistItems(c, playlistID)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	playlist := Playlist{
		Origin: "spotify",
		Songs: []Song{},
	}

	for _,item:= range results.Items{
		//probably more efficent if we created this outside and just updated the values
		song := Song{
			Name: item.Track.Track.Name,
			SpotifyId: string(item.Track.Track.ID),
			Artist: item.Track.Track.Artists[0].Name,
		}
		//fmt.Println(item.Track.Track.Artists[0].Name)
		playlist.Songs = append(playlist.Songs, song)

	}

	//fmt.Println(playlist)// this returns a slice with a bunch of crap but the artist name is the first thing that shows up in the slice
	c.JSON(200, playlist.Songs)
	
}