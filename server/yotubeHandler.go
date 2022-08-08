package main

import (
	"context"
	"fmt"

	"google.golang.org/api/youtube/v3"
)


func authYoutube() *youtube.YoutubeService{
	ctx := context.Background()

	yotubeService,err := youtube.NewService(ctx)
	if err!= nil{
		fmt.Println("Bad")
	}
	return youtube.NewYoutubeService(yotubeService)
}