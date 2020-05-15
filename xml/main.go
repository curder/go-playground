package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"time"
)

// Videos 视频数组
type Videos struct {
	XMLName xml.Name `xml:"videos"`
	Video   []Video  `xml:"video"`
}

// Video 视频
type Video struct {
	Version      string    `xml:"version"`
	ID           int       `xml:"id"`
	Title        string    `xml:"title"`
	Description  string    `xml:"description"`
	CreatedAt    time.Time `xml:"created_at"`
	URL          string    `xml:"url"`
	ThumbnailURL string    `xml:"thumbnail_url"`
	ImageURL     string    `xml:"image_url"`
	PlayCount    int       `xml:"play_count"`
	Duration     float64   `xml:"duration"`
	Height       int       `xml:"height"`
	Width        int       `xml:"width"`
}

func main() {
	var (
		data   []byte
		err    error
		videos Videos
		video  Video
	)
	if data, err = ioutil.ReadFile("./23125.xml"); err != nil {
		fmt.Printf("open file failed, err: %s", err.Error())
		return
	}

	if err = xml.Unmarshal(data, &videos); err != nil {
		fmt.Printf("XML file unmasrshaler fialed, err: %s", err.Error())
		return
	}

	for _, video = range videos.Video {
		fmt.Printf("%#v\n\n", video)
	}
}
