package main

import (
	"log"

	"github.com/Shalekar/youtube-autoup/pkg/upload"
	"google.golang.org/api/youtube/v3"
)

func main() {
	client := upload.GetClient(youtube.YoutubeUploadScope)
	service, err := youtube.New(client)
	if err != nil {
		log.Fatalf("Error creating YouTube client: %v", err)
	}
	upload.Upload(service, "testfile.mp4", "teste", "testDesc #Shorts", "22", "test1,test2", "public")
}
