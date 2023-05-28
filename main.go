package main

import (
	"log"

	"github.com/Shalekar/youtube-autoup/pkg/content"
	"github.com/Shalekar/youtube-autoup/pkg/tts"
	"github.com/Shalekar/youtube-autoup/pkg/upload"
	"google.golang.org/api/youtube/v3"
)

func main() {
	content := content.GetContent("AMD")
	file := tts.GetTTSFile(content.Content[:300], "change with key")
	client := upload.GetClient(youtube.YoutubeUploadScope)
	service, err := youtube.New(client)
	if err != nil {
		log.Fatalf("Error creating YouTube client: %v", err)
	}
	upload.Upload(service, file, "teste"+file, "testDesc #Shorts", "22", "test1,test2", "public")
}
