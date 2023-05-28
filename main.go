package main

import (
	"log"

	"github.com/Shalekar/youtube-autoup/pkg/content"
	"github.com/Shalekar/youtube-autoup/pkg/tts"
	"github.com/Shalekar/youtube-autoup/pkg/upload"
	ffmpeg_go "github.com/u2takey/ffmpeg-go"
	"google.golang.org/api/youtube/v3"
)

func main() {
	content := content.GetContent("AMD")
	var contentLen int
	if 1000 >= len(content.Content) {
		contentLen = len(content.Content) - 1
	} else {
		contentLen = 1000
	}
	file := tts.GetTTSFile(content.Content[:contentLen], "change with key")
	client := upload.GetClient(youtube.YoutubeUploadScope)
	service, err := youtube.New(client)
	if err != nil {
		log.Fatalf("Error creating YouTube client: %v", err)
	}
	ip1 := ffmpeg_go.Input("videoplayback.webm")
	ip2 := ffmpeg_go.Input(file)
	err1 := ffmpeg_go.Output([]*ffmpeg_go.Stream{ip1, ip2}, "output.mp4", ffmpeg_go.KwArgs{"t": 40}).OverWriteOutput().ErrorToStdOut().Run()
	if nil != err1 {
		log.Fatalf("Error creating video: %v", err1)
	}
	upload.Upload(service, "output.mp4", "teste"+file, "testDesc #Shorts", "22", "test1,test2", "public")
}
