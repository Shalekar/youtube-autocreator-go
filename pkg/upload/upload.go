package upload

import (
	"fmt"
	"log"
	"os"
	"strings"

	"google.golang.org/api/youtube/v3"
)

func Upload(service *youtube.Service, filename string, title string, description string, category string, keywords string, privacy string) {
	upload := &youtube.Video{
		Snippet: &youtube.VideoSnippet{
			Title:       title,
			Description: description,
			CategoryId:  category,
		},
		Status: &youtube.VideoStatus{PrivacyStatus: privacy},
	}

	// The API returns a 400 Bad Request response if tags is an empty string.
	if strings.Trim(keywords, "") != "" {
		upload.Snippet.Tags = strings.Split(keywords, ",")
	}

	call := service.Videos.Insert([]string{"snippet", "status"}, upload)

	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		log.Fatalf("Error opening %v: %v", filename, err)
	}

	response, err := call.Media(file).Do()
	if nil != err {
		log.Fatalf("Error making API call: %v", err.Error())
	}
	fmt.Printf("Upload successful! Video ID: %v\n", response.Id)
}
