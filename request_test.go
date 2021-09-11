package main

import (
	"encoding/json"
	"log"
	"net/http"
	"testing"
)

func TestEndPoint(t *testing.T) {
	endpoint := generateEndPoint(10)
	t.Log(endpoint)
}

func TestFetch(t *testing.T) {
	endPoint := generateEndPoint(10)
	resp, err := http.Get(endPoint)
	if nil != err {
		log.Println("get radios failed", err)
		return
	}
	defer func() { _ = resp.Body.Close() }()

	var response Response
	_ = json.NewDecoder(resp.Body).Decode(&response)
	radios := response.toRadio()

	rss := generateRss(radios)
	xml, err := rss.Xml()
	t.Log(string(xml))
}

func TestContentLength(t *testing.T) {
	url := "https://cdn.lizhi.fm/audio/2015/09/14/22877854718340102_hd.mp3"
	length := remoteContentLength(url)
	t.Log(length)
}
