package main

import (
	"context"
	"fmt"
	"log"
	"os"

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
	// function should pass in spotify playlist id
	results, err := Client.GetPlaylistItems(c, "59eBtjSiUluqHX1BweDZgu")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	playlist := Playlist{
		origin: "spotify",
		song: []Songs{},
	}

	for _,item:= range results.Items{
		//probably more efficent if we created this outside and just updated the values
		song :=Songs{
			name: item.Track.Track.Name,
			spotifyId: string(item.Track.Track.ID),
		}
		playlist.song=append(playlist.song,song)
		fmt.Println(item.Track.Track.Artists)// this returns a slice with a bunch of crap but the artist name is the first thing that shows up in the slice
	}
	//fmt.Printf("here  %v",playlist.song)


	c.String(200, results.Items[2].Track.Track.Name)
}