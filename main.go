package main

import (
	"flag"
	"log"
	"os"
	"strings"

	"github.com/Shalekar/youtube-autoup/pkg/content"
	"github.com/Shalekar/youtube-autoup/pkg/tts"
	"github.com/Shalekar/youtube-autoup/pkg/upload"
	ffmpeg_go "github.com/u2takey/ffmpeg-go"
	"google.golang.org/api/youtube/v3"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func deleteFiles(toDelete bool, list []string) {
	if toDelete {
		for i := 0; i < len(list); i++ {
			e := os.Remove(list[i])
			if e != nil {
				log.Fatal("Error deleting files: %v", e)
			}
		}
	}
}

var (
	contentId = flag.String("c", "AMD", "a string")
	key       = flag.String("k", "c88590bb66c749b99e2926295338088a", "a string")
	video     = flag.String("v", "videoplayback.webm", "a string")
	op        = flag.String("o", "output.mp4", "a string")
	oplength  = flag.Int("l", 40, "a string")
	deleteOp  = flag.Bool("d", true, "a bool")
)

func main() {
	flag.Parse()
	content := content.GetContent(*contentId, nil)
	var contentLen int
	if 1000 >= len(content.Content) {
		contentLen = len(content.Content) - 1
	} else {
		contentLen = 1000
	}
	file := tts.GetTTSFile(content.Title+content.Content[:contentLen], *key)
	fileList := []string{*op, file}

	ip1 := ffmpeg_go.Input(*video)
	ip2 := ffmpeg_go.Input(file)
	err1 := ffmpeg_go.Output([]*ffmpeg_go.Stream{ip1, ip2}, *op, ffmpeg_go.KwArgs{"t": *oplength}).OverWriteOutput().ErrorToStdOut().Run()
	if nil != err1 {
		deleteFiles(*deleteOp, fileList)
		log.Fatalf("Error creating video: %v", err1)
	}
	videoTitle := content.Title[:min(100, len(content.Title))]
	videoTitle = videoTitle[:strings.LastIndex(videoTitle, " ")]
	client := upload.GetClient(youtube.YoutubeUploadScope)
	service, err := youtube.New(client)
	if err != nil {
		deleteFiles(*deleteOp, fileList)
		log.Fatalf("Error creating YouTube client: %v", err)
	}
	upload.Upload(service, *op, videoTitle, "testDesc #Shorts", "22", "test1,test2", "public")
	deleteFiles(*deleteOp, fileList)
}
