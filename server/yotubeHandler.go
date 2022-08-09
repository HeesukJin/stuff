package main

import (
	"context"
	"fmt"
	"os"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

func authYoutube() *youtube.Service {
	ctx := context.Background()

	youtubeService, err := youtube.NewService(ctx, option.WithAPIKey(os.Getenv("YOUTUBE_API")))
	if err != nil {
		fmt.Println(err)
	}


	return youtubeService
}

func getSongs(playlist Playlist){
	
	test := make([]string, 2)

	test[0] = "id"
	test[1] = "snippet"
	for _, song := range playlist.Songs{
		ret, err  := YoutubeClient.Search.List(test).Q(song.Name+" "+song.Artist).Do()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(ret.Items[0].Snippet.Title)
	}

}
