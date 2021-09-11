package main

import "testing"

func TestEndPoint(t *testing.T) {
	endpoint := generateEndPoint(10)
	t.Log(endpoint)
}

func TestFetch(t *testing.T) {
	fetch(0, false)
}

func TestContentLength(t *testing.T) {
	url := "https://cdn.lizhi.fm/audio/2015/09/14/22877854718340102_hd.mp3"
	length := remoteContentLength(url)
	t.Log(length)
}
