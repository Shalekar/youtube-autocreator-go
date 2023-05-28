package content

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/Shalekar/youtube-autoup/pkg/structures"
)

func GetContent(subReddit string) *structures.Content {
	client := &http.Client{}
	URLBase := "https://www.reddit.com/r/%s/top.json?limit=10&after=%s"
	URL := fmt.Sprintf("https://www.reddit.com/r/%s/top.json?limit=10", subReddit)
	var content structures.Content
	var acceptedPost *structures.RedditTopListChild
	i := 0
	for ; i < 5; i++ {
		req, err := http.NewRequest("GET", URL, nil)
		if err != nil {
			log.Println(err)
			return nil
		}
		req.Header.Set("User-Agent", "RedditTopPostGetter/1.0")

		resp, err := client.Do(req)
		if err != nil {
			log.Println("ooopsss an error occurred, please try again")
			return nil
		}
		defer resp.Body.Close()
		var redditList structures.RedditTopList
		if err := json.NewDecoder(resp.Body).Decode(&redditList); err != nil {
			log.Println(err.Error())
			return nil
		}
		if len(redditList.Data.Children) < 1 {
			log.Println("ooopsss an error occurred, please try again")
			return nil
		}
		for j := 0; j < len(redditList.Data.Children); j++ {
			if "" != redditList.Data.Children[j].Data.Selftext {
				acceptedPost = &redditList.Data.Children[j]
			}
		}
		URL = fmt.Sprintf(URLBase, subReddit, redditList.Data.After)
		if nil != acceptedPost {
			break
		}
	}
	if 5 == i {
		log.Println("No text post available")
		return nil
	}
	content.Content = strings.Join(strings.Fields(acceptedPost.Data.Selftext), " ")
	content.Source = "reddit"
	content.Title = acceptedPost.Data.Title
	return &content
}
