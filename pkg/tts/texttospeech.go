package tts

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

func GetTTSFile(src string, key string) string {
	client := &http.Client{}
	URL := "http://api.voicerss.org/"
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	_ = writer.WriteField("key", key)
	_ = writer.WriteField("hl", "en-us")
	_ = writer.WriteField("src", src)
	_ = writer.WriteField("c", "MP3")
	_ = writer.WriteField("r", "3")
	_ = writer.WriteField("f", "48khz_16bit_mono")
	_ = writer.WriteField("v", "John")
	err := writer.Close()
	req, err := http.NewRequest("POST", URL, body)
	if err != nil {
		log.Println(err)
		return ""
	}
	req.Header.Set("User-Agent", "RedditTopPostGetter/1.0")
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := client.Do(req)
	if err != nil {
		log.Println("ooopsss an error occurred, please try again")
		return ""
	}
	defer resp.Body.Close()

	hash := md5.Sum([]byte(src))
	name := fmt.Sprintf("%s.mp3", hex.EncodeToString(hash[:]))
	out, err := os.Create(name)
	if err != nil {
		log.Println("ooopsss an error occurred, please try again")
		return ""
	}
	defer out.Close()
	_, er := io.Copy(out, resp.Body)
	if er != nil {
		log.Println("ooopsss an error occurred, please try again")
		return ""
	}
	return name
}
